package qforge

type QForge struct {
	server DatabaseServer
	graph  Graph
}

func (q *QForge) BuildQuery(selectColumnIds []string) string {
	return buildQuery(q, selectColumnIds)
}

func MakeQForge(server DatabaseServer) *QForge {
	return &QForge{
		server: server,
		graph:  server.CreateGraph(),
	}
}
