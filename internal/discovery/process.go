package discovery

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func DetectProcesses() ([]*process.Process, error) {
	return process.Processes()
}

func IsJavaProcess(p *process.Process) bool {
	name, _ := p.Name()
	cmd, _ := p.Cmdline()

	return strings.Contains(name, "java") ||
		strings.Contains(cmd, ".jar")
}

func IsMySQLProcess(p *process.Process) bool {
	name, _ := p.Name()
	return strings.Contains(strings.ToLower(name), "mysqld")
}

func IsRedisProcess(p *process.Process) bool {
	name, _ := p.Name()
	return strings.Contains(strings.ToLower(name), "redis")
}

