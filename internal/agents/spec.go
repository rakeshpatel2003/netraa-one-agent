package agents

type AgentSpec struct {
	Category string // INFRA, DATABASE, CACHE, APPLICATION
	Platform string // linux, windows
	Product  string // MYSQL, MSSQL, REDIS, HOST, JAVA
	Version  string // optional: 1.0.0, 9.4.0, empty = latest
	Artifact string // node_exporter, mysqld_exporter, otelcol
}

