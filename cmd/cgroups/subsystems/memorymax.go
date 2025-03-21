package subsystems

import (
	"fmt"
	"os"
	"path"
)

type MemoryMax struct{}

func (m *MemoryMax) Set(cgroupPath string, res *Resources) error {
	if res.MemoryMax == "" {
		return nil
	}

	err := os.WriteFile(path.Join(cgroupPath, "memory.max"), []byte(res.MemoryMax), 0644)
	return fmt.Errorf("set memory.max fail %v", err)
}
