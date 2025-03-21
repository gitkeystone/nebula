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
package main

import (
	"github.com/gitkeystone/nebula/cmd"
	"github.com/gitkeystone/nebula/logger"
)

func main() {
	_ = logger.GetLogger()
	defer logger.Sync() // 确保刷新日志缓冲区

	cmd.Execute()
}
