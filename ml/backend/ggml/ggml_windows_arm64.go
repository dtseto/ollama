//go:build windows && arm64

package ggml

import fsggml "github.com/ollama/ollama/fs/ggml"

func ConvertToF32(_ []byte, _ uint32, _ uint64) []float32 {
	return []float32{}
}

func Quantize(_ fsggml.TensorType, _ []float32, _ []uint64) []byte {
	return []byte{}
}
