package infrastructure

type DBType int

const (
	POSTGRES DBType = iota
	MYSQL
	SQLITE
	SQLSERVER
)
