package discovery

type ComponentType string

const (
	ComponentJava  ComponentType = "java"
	ComponentMySQL ComponentType = "mysql"
	ComponentRedis ComponentType = "redis"
	ComponentHost  ComponentType = "host"
)

type DiscoveredComponent struct {
	Type        ComponentType
	ProcessID   int32
	ProcessName string
	Ports       []int
	Metadata    map[string]string
}

