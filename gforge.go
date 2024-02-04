package qforge

type QForge struct {
	Server DatabaseServer
}

func (q *QForge) BuildQuery(selectColumnIds []string) string {
	return buildQuery(q.Server, selectColumnIds)
}
