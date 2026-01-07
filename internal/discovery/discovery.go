package discovery

func Discover() ([]DiscoveredComponent, error) {
	var components []DiscoveredComponent

	// Host
	components = append(components, DetectHost())

	processes, err := DetectProcesses()
	if err != nil {
		return components, err
	}

	for _, p := range processes {
		pid := p.Pid
		name, _ := p.Name()

		switch {
		case IsJavaProcess(p):
			components = append(components, DiscoveredComponent{
				Type:        ComponentJava,
				ProcessID:   pid,
				ProcessName: name,
				Ports:       PortsByPID(pid),
			})

		case IsMySQLProcess(p):
			components = append(components, DiscoveredComponent{
				Type:        ComponentMySQL,
				ProcessID:   pid,
				ProcessName: name,
				Ports:       PortsByPID(pid),
			})

		case IsRedisProcess(p):
			components = append(components, DiscoveredComponent{
				Type:        ComponentRedis,
				ProcessID:   pid,
				ProcessName: name,
				Ports:       PortsByPID(pid),
			})
		}
	}

	return components, nil
}

