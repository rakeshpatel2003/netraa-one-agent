package host

import (
	"fmt"
	"os"
	"path/filepath"

	"netraa-agent/internal/agents"
)

const HostBaseDir = "/opt/netraa/host"

func Install() error {
	fmt.Println("▶ Installing Host Monitoring")

	// 1. Prepare base directory
	if err := os.MkdirAll(HostBaseDir, 0750); err != nil {
		return err
	}

	// 2. Define host agents using GENERIC AgentSpec
	hostAgents := []struct {
		Spec       agents.AgentSpec
		TargetName string
	}{
		{
			Spec: agents.AgentSpec{
				Category: "INFRA",
				Artifact: "node_exporter",
			},
			TargetName: "node_exporter",
		},
		{
			Spec: agents.AgentSpec{
				Category: "INFRA",
				Artifact: "netflow-agent",
			},
			TargetName: "netflow-agent",
		},
		{
			Spec: agents.AgentSpec{
				Category: "INFRA",
				Artifact: "otelcol",
			},
			TargetName: "otelcol",
		},
	}

	// 3. Install agents
	for _, a := range hostAgents {
		if err := installAgent(a.Spec, a.TargetName); err != nil {
			return err
		}
	}

	// 4. Generate OTel config (IP-based)
	if err := writeOtelConfig(HostBaseDir); err != nil {
		return err
	}

	// 5. Secure everything
	// if err := securePath(HostBaseDir); err != nil {
	// 	return err
	// }

	// 6. Install & start services
	if err := installNodeExporter(filepath.Join(HostBaseDir, "node_exporter")); err != nil {
		return err
	}

	if err := installNetflowAgent(filepath.Join(HostBaseDir, "netflow-agent")); err != nil {
		return err
	}

	if err := installOtelCollector(
		filepath.Join(HostBaseDir, "otelcol"),
		filepath.Join(HostBaseDir, "otel-collector.yaml"),
	); err != nil {
		return err
	}

	fmt.Println("✔ Host monitoring installed successfully")
	return nil
}

