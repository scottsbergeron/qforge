package qforge

import "fmt"

type QForge struct {
	server DatabaseServer
	graph  Graph
}

func (q *QForge) BuildQuery(startTable string, selectColumnIds []string) (string, error) {
	if vertex := q.graph.GetVertex(startTable); vertex == nil {
		return "", fmt.Errorf("no table located: %s", startTable)
	} else {
		fmt.Println(vertex.Traverse(selectColumnIds))
		return vertex.Key(), nil
	}
}

func MakeQForge(server DatabaseServer) *QForge {
	return &QForge{
		server: server,
		graph:  server.CreateGraph(),
	}
}
