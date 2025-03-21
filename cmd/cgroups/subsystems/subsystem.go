package subsystems

// 用于传递资源限制配置的结构体，包含内存限制， CPU 时间片权重 ， CPU 核心数
type Resources struct {
	MemoryMax  string
	CpuWeight  string
	CpusetCpus string
}

type Subsystem interface {
	Set(path string, res *Resources) error
}

var (
	SubsystemIns = []Subsystem{
		&MemoryMax{},
		&CpuWeight{},
		&CpusetCpus{},
	}
)
