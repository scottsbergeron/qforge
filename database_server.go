package qforge

type DatabaseServerType int

const (
	SnowflakeDBType DatabaseServerType = iota
	PostgresDBType
)

type DatabaseServer interface {
	CreateGraph() Graph
	ServerType() DatabaseServerType
}

type SnowflakeDatabaseServer struct {
	Databases []*Database
}

func (s SnowflakeDatabaseServer) CreateGraph() Graph {
	return createGraph(s.Databases)
}

func (s SnowflakeDatabaseServer) ServerType() DatabaseServerType {
	return SnowflakeDBType
}

type PostgresDatabaseServer struct {
	Databases []*Database
}

func (s PostgresDatabaseServer) CreateGraph() Graph {
	return createGraph(s.Databases)
}

func (s PostgresDatabaseServer) ServerType() DatabaseServerType {
	return PostgresDBType
}
