package qforge

import "fmt"

type Graph struct {
	Vertices  []Vertex
	ColumnIds Set[string]
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
		Edges:    []Edge{},
	}
	g.Vertices = append(g.Vertices, vertex)
	return &vertex
}

type Vertex struct {
	Database *Database
	Schema   *DatabaseSchema
	Table    *Table
	Edges    []Edge
}

func (v *Vertex) AddEdge() {

}

func (v *Vertex) Key() string {
	return fmt.Sprintf("\"%s\".\"%s\".\"%s\"", v.Database.Name, v.Schema.Name, v.Table.Name)
}

type Edge struct {
	Vertex       *Vertex
	SameDatabase bool
	SameSchema   bool
	ColumnIds    Set[string]
}

func createGraph(databases []*Database) Graph {
	graph := Graph{
		Vertices:  []Vertex{},
		ColumnIds: Set[string]{},
	}

	// Create initial vertex list
	vertexMap := createVertexMap(&graph, databases)

	// Create a column id => vertex key map
	vertexKeyMap := createVertexKeyMap(vertexMap)
	fmt.Println(vertexKeyMap)

	// TODO: for each vertex's table's columns, add an edge to other vertices if not exist
	// This should move the algo from O(n^2) to O(n)

	//for _, table2 := range schema.Tables[i+1:] {
	//	sharedColumns := table1.GetSharedColumns(&table2)
	//	if len(sharedColumns) > 0 {
	//		v1 := fmt.Sprintf("%s.%s", schema.Name, table1.Name)
	//		v2 := fmt.Sprintf("%s.%s", schema.Name, table2.Name)
	//		graph.AddEdge(v1, v2)
	//	}
	//}

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

func createVertexKeyMap(vertexMap map[string]*Vertex) map[string]Set[string] {
	vertexKeyMap := make(map[string]Set[string])
	for key, value := range vertexMap {
		for _, column := range value.Table.Columns {
			if _, exists := vertexKeyMap[column.Id]; !exists {
				vertexKeyMap[column.Id] = Set[string]{}
			}
			vertexKeyMap[column.Id].Add(key)
		}
	}
	return vertexKeyMap
}
