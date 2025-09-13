//go:build windows && arm64

package discover

// Do NOT redeclare SystemInfo or GpuInfoList here; they exist in types.go

func GetSystemInfo() SystemInfo {
    return SystemInfo{}
}

func GetCPUInfo() GpuInfoList {
    return GpuInfoList{}
}

func GetGPUInfo() GpuInfoList {
    return GpuInfoList{}
}

func (g GpuInfoList) GetVisibleDevicesEnv() []string {
    return []string{}
}
