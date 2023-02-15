package dto

type SysInfo struct {
	CpuInfo  CpuInfo             `json:"cpuInfo"`
	MemInfo  MemInfo             `json:"memInfo"`
	DiskNode map[string]DiskNode `json:"diskNode"`
	HostInfo HostInfo            `json:"hostInfo"`
	NetNode  map[string]NetNode  `json:"netNode"`
}

type CpuInfo struct {
	// cpu厂商
	CpuVendor string `json:"CpuVendor"`
	// cpu型号
	CpuName string
	// 核心数
	CpuCore int
	// 逻辑线程
	CpuHt int
	// cpu频率
	CpuHz float64
}

type MemInfo struct {
	//	总内存
	Total int
	// 使用内存
	Used int
	// 剩余内存
	Free int
	// 使用占比
	UsedPercent float64
}

type DiskNode struct {
	//	总磁盘大小
	Total int
	// 使用磁盘大小
	Used int
	// 剩余磁盘大小
	Free int
	// 使用占比
	UsedPercent float64
}

type HostInfo struct {
	// 主机用户名称
	HostName string
	// BootTime
	BootTime string
	// 运行时间
	UpTime string
	// 操作系统
	OS string
	// 操作平台
	PlatForm string
	// 系统版本号
	PlatFormVersion string
	// 系统架构
	KernelArch string
	// 系统标识
	HostID string
}

type NetNode struct {
	// 字节发送量
	BytesSent int
	// 字节接收量
	BytesRect int
	// 包发送量
	PacketsSent int
	// 包接收量
	PacketsRect int
}
