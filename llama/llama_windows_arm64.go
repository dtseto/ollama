//go:build windows && arm64

package llama

type TokenData struct {
	ID    int32
	Logit float32
}

type Grammar struct{}

func NewGrammar(_ string, _ []uint32, _ []string, _ []int32) *Grammar {
	return &Grammar{}
}

func (g *Grammar) Apply(_ []TokenData) {}
func (g *Grammar) Accept(_ int32) bool { return true }
func (g *Grammar) Free()               {}

type Model struct{}

type ModelParams struct {
	VocabOnly    bool
	NumGpuLayers int
	MainGpu      int
	UseMmap      bool
	TensorSplit  []float32
	Progress     func(float32)
}

func LoadModelFromFile(_ string, _ ModelParams) (*Model, error) {
	return &Model{}, nil
}

func FreeModel(_ *Model) {}

func SchemaToGrammar(_ []byte) []byte { return []byte{} }

func (m *Model) Tokenize(_ string, _ bool, _ bool) ([]int, error) {
	return []int{}, nil
}

func (m *Model) TokenToPiece(_ int) string {
	return ""
}

// Context is a placeholder for the llama.cpp context object
type Context struct{}

// SamplingContext is a placeholder for sampling state
type SamplingContext struct{}

// SamplingParams is a placeholder for sampling parameters
type SamplingParams struct {
	TopK           int
	TopP           float32
	MinP           float32
	TypicalP       float32
	Temp           float32
	RepeatLastN    int
	PenaltyRepeat  float32
	PenaltyFreq    float32
	PenaltyPresent float32
	Seed           uint32
	Grammar        interface{} // was *Grammar
}

// Batch is a placeholder for batch inference data
type Batch struct{}

// KV cache helpers
func (c *Context) KvCacheSeqRm(_ int, _ int, _ int) bool    { return false }
func (c *Context) KvCacheSeqCp(_ int, _ int, _ int, _ int)  {}
func (c *Context) KvCacheCanShift() bool                    { return false }
func (c *Context) KvCacheSeqAdd(_ int, _ int, _ int, _ int) {}

// Vision/CLIP detection from a model path
func GetModelArch(_ string) (string, error) {
	return "", nil
}

// Multimodal context
type MtmdContext struct{}

func NewMtmdContext(_ *Context, _ string) (*MtmdContext, error) {
	return &MtmdContext{}, nil
}

func (m *MtmdContext) Free() {}

func (m *MtmdContext) NewEmbed(_ *Context, _ []byte) ([][]float32, error) {
	return nil, nil
}

// Context -> Model accessor and embedding size
func (c *Context) Model() *Model { return &Model{} }

func (m *Model) NEmbd() int { return 0 }

// BOS token control on model
func (m *Model) AddBOSToken() bool { return false }

func NewBatch(_ int, _ int, _ int) (*Batch, error) {
	return &Batch{}, nil
}

func (b *Batch) Free() {}

func (b *Batch) Clear() {}

func (b *Batch) IsEmbedding() bool { return false }

func (b *Batch) Size() int { return 0 }

func (b *Batch) Add(_ ...interface{}) {}

func (b *Batch) NumTokens() int { return 0 }

func NewSamplingContext(_ *Model, _ SamplingParams) (*SamplingContext, error) {
	return &SamplingContext{}, nil
}

func (sc *SamplingContext) Accept(_ int, _ bool) {}

func (sc *SamplingContext) Sample(_ *Context, _ int) int {
	return 0
}

func (c *Context) Decode(_ *Batch) error {
	return nil
}

func (c *Context) GetEmbeddingsSeq(_ int) []float32 {
	return nil
}

func (c *Context) GetEmbeddingsIth(_ int) []float32 {
	return nil
}

func (m *Model) TokenIsEog(_ int) bool {
	return false
}

type ContextParams struct{}

func NewContextParams(_ int, _ int, _ int, _ int, _ bool, _ string) ContextParams {
	return ContextParams{}
}

func NewContextWithModel(_ *Model, _ ContextParams) (*Context, error) {
	return &Context{}, nil
}

func (m *Model) ApplyLoraFromFile(_ *Context, _ string, _ float32, _ int) error {
	return nil
}

func EnumerateGPUs() []string {
	return nil
}

func BackendInit() {}
