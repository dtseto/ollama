//go:build windows && arm64

package discover

var (
	rocmMinimumMemory uint64 = 0
	IGPUMemLimit      uint64 = 0
	unsupportedGPUs   []UnsupportedGPUInfo
)