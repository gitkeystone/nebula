package subsystems

import (
	"fmt"
	"os"
	"path"
)

type CpusetCpus struct{}

func (c *CpusetCpus) Set(cgroupPath string, res *Resources) error {
	if res.CpusetCpus == "" {
		return nil
	}

	err := os.WriteFile(path.Join(cgroupPath, "cpuset.cpus"), []byte(res.CpusetCpus), 0644)
	return fmt.Errorf("set cpuset.cpus fail %v", err)
}
