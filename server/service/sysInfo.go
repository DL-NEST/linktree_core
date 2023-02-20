package service

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"linktree_core/server/model/dto"
	"strings"
	"time"
)

// sysInfoService
type sysInfoService struct {
}

func (sysInfoService) GetCpuInfo() dto.CpuInfo {
	core, _ := cpu.Info()
	coreInfo := core[0]
	cores, _ := cpu.Counts(false)
	Ht, _ := cpu.Counts(true)
	cpuInfo := dto.CpuInfo{
		CpuVendor: coreInfo.VendorID,
		CpuName:   strings.Replace(coreInfo.ModelName, " ", "", -1),
		CpuCore:   cores,
		CpuHt:     Ht,
		CpuHz:     coreInfo.Mhz / 1000,
	}
	return cpuInfo
}

func (sysInfoService) GetMemInfo() dto.MemInfo {
	ramInfo, _ := mem.VirtualMemory()
	memInfo := dto.MemInfo{
		Total:       int(ramInfo.Total),
		Used:        int(ramInfo.Used),
		Free:        int(ramInfo.Free),
		UsedPercent: ramInfo.UsedPercent,
	}
	return memInfo
}

func (sysInfoService) GetDiskInfo() map[string]dto.DiskNode {
	diskInfo := make(map[string]dto.DiskNode)
	diskList, _ := disk.Partitions(true)
	for i := range diskList {
		diskIO, _ := disk.Usage(diskList[i].Device)
		diskInfo[diskList[i].Device] = dto.DiskNode{
			Total:       int(diskIO.Total / 1024 / 1024 / 1024),
			Used:        int(diskIO.Used / 1024 / 1024 / 1024),
			Free:        int(diskIO.Free / 1024 / 1024 / 1024),
			UsedPercent: diskIO.UsedPercent,
		}
	}
	return diskInfo
}

func (sysInfoService) GetHostInfo() dto.HostInfo {
	hostInfo, _ := host.Info()
	hosts := dto.HostInfo{
		HostName:        hostInfo.Hostname,
		BootTime:        time.Unix(int64(hostInfo.BootTime), 0).Format("2006-01-02 15:04:05"),
		UpTime:          time.Unix(int64(hostInfo.Uptime), 0).Format("15:04:05"),
		OS:              hostInfo.OS,
		PlatForm:        hostInfo.Platform,
		PlatFormVersion: hostInfo.PlatformVersion,
		KernelArch:      hostInfo.KernelArch,
		HostID:          hostInfo.HostID,
	}
	return hosts
}

func (sysInfoService) GetNetInfo() map[string]dto.NetNode {
	netInfo, _ := net.IOCounters(true)
	NetInfo := make(map[string]dto.NetNode)
	for i := range netInfo {
		netIo := netInfo[i]
		NetInfo[netIo.Name] = dto.NetNode{
			BytesSent:   int(netIo.BytesSent),
			BytesRect:   int(netIo.BytesRecv),
			PacketsSent: int(netIo.PacketsSent),
			PacketsRect: int(netIo.PacketsRecv),
		}
	}
	return NetInfo
}

func (sysInfoService) GetSysInfo() dto.SysInfo {
	var SysInfoService sysInfoService
	return dto.SysInfo{
		CpuInfo:  SysInfoService.GetCpuInfo(),
		MemInfo:  SysInfoService.GetMemInfo(),
		DiskNode: SysInfoService.GetDiskInfo(),
		HostInfo: SysInfoService.GetHostInfo(),
		NetNode:  SysInfoService.GetNetInfo(),
	}
}
