package subsystems

import (
	"fmt"
	"os"
	"path"
)

type CpuWeight struct{}

func (c *CpuWeight) Set(cgroupPath string, res *Resources) error {
	if res.CpuWeight == "" {
		return nil
	}

	err := os.WriteFile(path.Join(cgroupPath, "cpu.weight"), []byte(res.CpuWeight), 0644)
	return fmt.Errorf("set cpu weight fail %v", err)
}
