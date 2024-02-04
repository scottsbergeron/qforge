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

type Edge struct {
	Vertex    *Vertex
	ColumnIds []string
}

func CreateGraph(databases []Database) Graph {
	graph := Graph{
		Vertices: []Vertex{},
	}

	for _, db := range databases {
		for _, schema := range db.Schemas {
			for i, table1 := range schema.Tables {
				vertex := graph.AddVertex(&db, &schema, &table1)

				// TODO: create a "db.schema.table" => vertex map
				// TODO: create a column.Id => []"db.schema.table"
				// TODO: for each vertex's table's columns, add an edge to other vertices if not exist
				// This should move the algo from O(n^2) to O(n)
				fmt.Println(i)
				fmt.Println(vertex)

				//for _, table2 := range schema.Tables[i+1:] {
				//	sharedColumns := table1.GetSharedColumns(&table2)
				//	if len(sharedColumns) > 0 {
				//		v1 := fmt.Sprintf("%s.%s", schema.Name, table1.Name)
				//		v2 := fmt.Sprintf("%s.%s", schema.Name, table2.Name)
				//		graph.AddEdge(v1, v2)
				//	}
				//}
			}
		}
	}

	return graph
}
