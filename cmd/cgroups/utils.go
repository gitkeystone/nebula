package cgroups

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func FindCgroupMountPoint(mountOpt string) (string, error) {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], ",") {
			if opt == mountOpt {
				return fields[4], nil
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", err
}

// 得到cgroup在文件系统中的绝对路径
func GetCgroupPath(mountOpt string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot, err := FindCgroupMountPoint(mountOpt)
	if err != nil {
		return "", err
	}

	if _, err := os.Stat(path.Join(cgroupRoot, "system.slice", cgroupPath)); err == nil || (os.IsNotExist(err) && autoCreate) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, "system.slice", cgroupPath), 0755); err != nil {
				return "", fmt.Errorf("create cgroup fail %v", err)
			}
		}
		return path.Join(cgroupRoot, "system.slice", cgroupPath), nil
	}
	return "", fmt.Errorf("cgroup path fail %v", err)
}
