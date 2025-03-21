package container

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"go.uber.org/zap"
)

func RunContainerInitProcess() error {
	cmdArray := readUserCommand()
	if cmdArray == nil || len(cmdArray) == 0 {
		return fmt.Errorf("Run container get user command error, cmdArray is nil")
	}

	setUpMount()

	path, err := exec.LookPath(cmdArray[0])
	if err != nil {
		log.Error("Exec loop path", zap.Error(err))
	}

	log.Info("Find path", zap.String("path", path))

	if err := syscall.Exec(path, cmdArray, os.Environ()); err != nil {
		log.Error("proc fs mount failed", zap.Error(err))
	}
	return nil
}

func readUserCommand() []string {
	pipe := os.NewFile(uintptr(3), "pipe")
	defer pipe.Close()
	msg, err := io.ReadAll(pipe)
	if err != nil {
		log.Error("init read pipe error %v", zap.Error(err))
		return nil
	}
	return strings.Split(string(msg), " ")
}

func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Error("Get current location error %v", zap.Error(err))
		return
	}

	log.Info("Current location", zap.String("location", pwd))

	pivotRoot(pwd)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}

func pivotRoot(newroot string) error {
	if err := syscall.Mount(newroot, newroot, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to itself error: %v", err)
	}

	putold := filepath.Join(newroot, ".pivot_root")
	if err := os.MkdirAll(putold, 0777); err != nil {
		return err
	}

	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}

	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	putold = "/.pivot_root"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir %v", err)
	}

	return os.Remove(putold)
}
