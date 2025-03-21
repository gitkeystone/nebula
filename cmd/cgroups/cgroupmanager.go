package cgroups

import (
	"os"

	"github.com/gitkeystone/nebula/cmd/cgroups/subsystems"
)

type CgroupManager struct {
	Path      string
	Resources *subsystems.Resources
}

func NewCgroupManager(path string) *CgroupManager {
	subSystemCgroupPath, err := GetCgroupPath("nsdelegate", path, true)
	if err != nil {
		return nil
	}

	return &CgroupManager{
		Path: subSystemCgroupPath,
	}
}

// 将进程PID加入到每个cgroup中
func (c *CgroupManager) Apply(pid int) error {
	return NewCgroupProcs().Apply(c.Path, pid)
}

// 设置cgroup资源限制
func (c *CgroupManager) Set(res *subsystems.Resources) error {
	for _, subSysIns := range subsystems.SubsystemIns {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

// 释放cgroup资源
func (c *CgroupManager) Destroy() error {
	return os.Remove(c.Path)
}
