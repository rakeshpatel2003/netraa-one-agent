package host

import (
	"os"
	"os/exec"
	"fmt"
)

func installNetflowAgent(binPath string) error {
	service := `
[Unit]
Description=Netraa Netflow Agent
After=network.target

[Service]
User=root
ExecStart=` + binPath + `
Restart=always
CapabilityBoundingSet=CAP_NET_RAW CAP_NET_ADMIN

[Install]
WantedBy=multi-user.target
`
	path := "/etc/systemd/system/netraa-netflow.service"
	if err := os.WriteFile(path, []byte(service), 0644); err != nil {
		return err
	}

	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", "--now", "netraa-netflow").Run()

	if err := exec.Command("systemctl", "is-active", "--quiet", "netraa-netflow").Run(); err != nil {
		return fmt.Errorf("netflow failed to start")
	}

	return nil
}

