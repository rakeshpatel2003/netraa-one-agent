package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"netraa-agent/internal/discovery"
)

func AskUser(components []discovery.DiscoveredComponent) []discovery.ComponentType {
	items := make([]PromptItem, 0)

	for _, c := range components {
		items = append(items, PromptItem{
			Component: c,
			Selected:  c.Type == discovery.ComponentHost, // host default ON
		})
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		printMenu(items)

		fmt.Print("\nToggle component number (or press Enter to continue): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			break
		}

		idx := int(input[0]-'1')
		if idx >= 0 && idx < len(items) {
			items[idx].Selected = !items[idx].Selected
		}
	}

	return selectedTypes(items)
}

func printMenu(items []PromptItem) {
	fmt.Println("\nDetected components:\n")

	for i, item := range items {
		check := " "
		if item.Selected {
			check = "✔"
		}

		fmt.Printf("[%s] %d) %s %s\n",
			check,
			i+1,
			displayName(item.Component),
			componentDetails(item.Component),
		)
	}
}

func displayName(c discovery.DiscoveredComponent) string {
	switch c.Type {
	case discovery.ComponentJava:
		return "Java Application"
	case discovery.ComponentMySQL:
		return "MySQL Database"
	case discovery.ComponentRedis:
		return "Redis Cache"
	case discovery.ComponentHost:
		return "Linux Host Metrics"
	default:
		return string(c.Type)
	}
}

func componentDetails(c discovery.DiscoveredComponent) string {
	if c.ProcessID == 0 {
		return ""
	}
	return fmt.Sprintf("(PID %d, Ports %v)", c.ProcessID, c.Ports)
}

func selectedTypes(items []PromptItem) []discovery.ComponentType {
	var selected []discovery.ComponentType
	for _, i := range items {
		if i.Selected {
			selected = append(selected, i.Component.Type)
		}
	}
	return selected
}

