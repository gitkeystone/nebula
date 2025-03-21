/*
Copyright © 2025 Chen Xiaohui <ysucxh@163.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"os"
	"strings"

	"github.com/gitkeystone/nebula/cmd/cgroups"
	"github.com/gitkeystone/nebula/cmd/cgroups/subsystems"
	"github.com/gitkeystone/nebula/cmd/container"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Create a container with namespace and cgroups limit",
	// Long: `A longer description that spans multiple lines and likely contains examples
	//and usage of using your command. For example:

	//Cobra is a CLI library for Go that empowers applications.
	//This application is a tool to generate the needed files
	//to quickly create a Cobra application.`,

	Example: "nebula run [flags] command",
	Args:    cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		tty, err := cmd.Flags().GetBool("tty")
		cobra.CheckErr(err)

		interactive, err := cmd.Flags().GetBool("interactive")
		cobra.CheckErr(err)

		memory, err := cmd.Flags().GetString("memory")
		cobra.CheckErr(err)

		cpuWeight, err := cmd.Flags().GetString("cpu-weight")
		cobra.CheckErr(err)

		cpuSet, err := cmd.Flags().GetString("cpuset-cpus")
		cobra.CheckErr(err)

		res := &subsystems.Resources{
			MemoryMax:  memory,
			CpuWeight:  cpuWeight,
			CpusetCpus: cpuSet,
		}

		Run(tty && interactive, args, res)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().StringP("memory", "m", "", "Memory limit")
	runCmd.Flags().StringP("cpu-weight", "c", "", "CPU shares (relative weight)")
	runCmd.Flags().StringP("cpuset-cpus", "", "", "CPUs in which to allow execution (0-3, 0,1)")
}

func Run(tty bool, cmdArray []string, res *subsystems.Resources) {
	parent, writePipe := container.NewParentProcess(tty)
	if parent == nil {
		log.Error("New parent process fail")
		return
	}

	if err := parent.Start(); err != nil {
		log.Error("parent process start failed", zap.Error(err))
	}

	// 创建 cgroup manager, 并设置资源限制，使得限制在容器上生效
	cgroupManager := cgroups.NewCgroupManager("nebula-" + hexString())
	defer cgroupManager.Destroy()

	// 设置资源限制
	cgroupManager.Set(res)

	// 将容器进程加入到cgroup中
	cgroupManager.Apply(parent.Process.Pid)

	sendInitCommand(cmdArray, writePipe)

	if tty {
		_ = parent.Wait()
	}
}

func sendInitCommand(comArray []string, writePipe *os.File) {
	command := strings.Join(comArray, " ")
	log.Info("full command", zap.String("command", command))
	writePipe.WriteString(command)
	writePipe.Close()
}

func hexString() string {
	b := make([]byte, 32)
	rand.Read(b)

	return hex.EncodeToString(b)
}
