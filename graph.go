package qforge

import "fmt"

type Graph struct {
	Vertices []Vertex
}

func (g *Graph) AddVertex(
	database *Database,
	schema *DatabaseSchema,
	table *Table,
) *Vertex {
	vertex := Vertex{
		Database: database,
		Schema:   schema,
		Table:    table,
		Edges:    make(map[string]Edge),
	}
	g.Vertices = append(g.Vertices, vertex)
	return &vertex
}

func (g *Graph) GetVertex(tableId string) *Vertex {
	// TODO: A map would be more efficient
	for _, vertex := range g.Vertices {
		if vertex.Key() == tableId {
			return &vertex
		}
	}
	return nil
}

type Vertex struct {
	Database *Database
	Schema   *DatabaseSchema
	Table    *Table
	Edges    map[string]Edge
}

func (v *Vertex) AddEdge(
	vertex *Vertex,
) *Edge {
	sharedColumns := v.Table.GetSharedColumns(vertex.Table)
	sharedColumnIds := Set[string]{}
	for _, sharedColumn := range sharedColumns {
		sharedColumnIds.Add(sharedColumn.Id)
	}
	edge := Edge{
		Vertex:       vertex,
		SameDatabase: v.Database.Name == vertex.Database.Name,
		ColumnIds:    sharedColumnIds,
	}
	v.Edges[vertex.Key()] = edge
	return &edge
}

func (v *Vertex) Key() string {
	return fmt.Sprintf("\"%s\".\"%s\".\"%s\"", v.Database.Name, v.Schema.Name, v.Table.Name)
}

func (v *Vertex) Traverse(columnIds []string) Vertex {
	queue := []*Vertex{v}
	visited := Set[*Vertex]{}
	visited.Add(v)
	for len(queue) > 0 {
		currentVertex := queue[0]
		queue = queue[1:]

		for _, edge := range currentVertex.Edges {
			if _, exists := visited[edge.Vertex]; exists {
				continue
			}
			// TODO: Keep track of found columns, return the smallest graph that contains all columns
			// TODO: Have a function to determine if an edge is "allowed". i.e. Cross-database.

			queue = append(queue, edge.Vertex)
			visited.Add(edge.Vertex)
		}
	}
	return Vertex{}
}

type Edge struct {
	Vertex       *Vertex
	SameDatabase bool
	ColumnIds    Set[string]
}

func createGraph(databases []*Database) Graph {
	graph := Graph{
		Vertices: []Vertex{},
	}
	// TODO: Duplicate Vertices are being created for Teacher in the current setup, investigate
	vertexMap := createVertexMap(&graph, databases)
	vertexKeyMap := createVertexKeyMap(vertexMap)
	addEdgesToVertices(vertexMap, vertexKeyMap)
	return graph
}

func createVertexMap(graph *Graph, databases []*Database) map[string]*Vertex {
	vertexMap := make(map[string]*Vertex)
	for _, db := range databases {
		for _, schema := range db.Schemas {
			for _, table := range schema.Tables {
				vertex := graph.AddVertex(db, schema, table)
				vertexMap[vertex.Key()] = vertex
			}
		}
	}
	return vertexMap
}

func createVertexKeyMap(vertexMap map[string]*Vertex) map[string]*Set[string] {
	vertexKeyMap := make(map[string]*Set[string])
	for key, value := range vertexMap {
		for _, column := range value.Table.Columns {
			if _, exists := vertexKeyMap[column.Id]; !exists {
				s := Set[string]{}
				vertexKeyMap[column.Id] = &s
			}
			vertexKeyMap[column.Id].Add(key)
		}
	}
	return vertexKeyMap
}

func addEdgesToVertices(vertexMap map[string]*Vertex, vertexKeyMap map[string]*Set[string]) {
	for _, vertex := range vertexMap {
		for _, column := range vertex.Table.Columns {
			if vKeys, exists := vertexKeyMap[column.Id]; exists {
				for vKey := range *vKeys {
					if vertex.Key() != vKey {
						vertex.AddEdge(vertexMap[vKey])
					}
				}
			}
		}
	}
}
