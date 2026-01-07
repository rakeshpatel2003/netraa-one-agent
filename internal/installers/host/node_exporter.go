package host

import (
	"fmt"
	"os"
	"os/exec"
)

func installNodeExporter(binPath string) error {
	service := `
[Unit]
Description=Netraa Node Exporter
After=network.target

[Service]
User=root
Group=root
ExecStart=` + binPath + `
Restart=always
RestartSec=5
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true

[Install]
WantedBy=multi-user.target
`

	path := "/etc/systemd/system/node-exporter.service"
	if err := os.WriteFile(path, []byte(service), 0644); err != nil {
		return err
	}

	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "--now", "node-exporter").Run()

	if err := exec.Command("systemctl", "is-active", "--quiet", "node-exporter").Run(); err != nil {
		return fmt.Errorf("node_exporter failed to start")
	}

	return nil
}
