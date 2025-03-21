package cgroups

import (
	"fmt"
	"os"
	"path"
	"strconv"
)

type CgroupProcs struct{}

func NewCgroupProcs() *CgroupProcs {
	return &CgroupProcs{}
}

func (c *CgroupProcs) Apply(cgroupPath string, pid int) error {
	err := os.WriteFile(path.Join(cgroupPath, "cgroup.procs"), []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		return fmt.Errorf("set cgroup.proc fail %v", err)
	}
	return nil
}
