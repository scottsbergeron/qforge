package qforge

func buildQuery(server DatabaseServer, selectColumnIds []string) string {
	return string(server.ServerType())
}
