package host

import (
	"fmt"
	"os"
	"path/filepath"

	"netraa-agent/internal/utils"
)

func writeOtelConfig(baseDir string) error {
	ip := utils.GetPrimaryIP()

	config := fmt.Sprintf(`
receivers:
  prometheus:
    config:
      scrape_configs:
        - job_name: "vm-monitoring"
          scrape_interval: 30s
          static_configs:
            - targets: ['%s:9100']
            - targets: ['%s:9101']

processors:
  batch:

exporters:
  otlp:
    endpoint: "localhost:11805"
    tls:
      insecure: true

service:
  pipelines:
    metrics:
      receivers: [prometheus]
      processors: [batch]
      exporters: [otlp]
`, ip, ip)

	path := filepath.Join(baseDir, "otel-collector.yaml")
	return os.WriteFile(path, []byte(config), 0640)
}

