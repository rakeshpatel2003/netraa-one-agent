package agents

import (
	"os"
	"sort"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func latestVersion(base string) (string, error) {
	entries, err := os.ReadDir(base)
	if err != nil {
		return "", err
	}

	var versions []string
	for _, e := range entries {
		if e.IsDir() {
			versions = append(versions, e.Name())
		}
	}

	sort.Strings(versions)
	if len(versions) == 0 {
		return "", os.ErrNotExist
	}

	return versions[len(versions)-1], nil
}

