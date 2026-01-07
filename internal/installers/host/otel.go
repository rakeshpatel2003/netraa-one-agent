package host

import (
	"fmt"
	"os"
	"os/exec"
)

func installOtelCollector(binPath, configPath string) error {
	service := `
[Unit]
Description=Netraa OTel Collector
After=network.target

[Service]
User=root
Group=root
ExecStart=` + binPath + ` --config ` + configPath + `
Restart=always
RestartSec=5
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/netraa

[Install]
WantedBy=multi-user.target
`

	path := "/etc/systemd/system/otelcol.service"
	if err := os.WriteFile(path, []byte(service), 0644); err != nil {
		return err
	}

	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "--now", "otelcol").Run()

	// ✅ Verify service is running
	if err := exec.Command("systemctl", "is-active", "--quiet", "otelcol").Run(); err != nil {
		return fmt.Errorf("otelcol failed to start")
	}

	return nil
}
