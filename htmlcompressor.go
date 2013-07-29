package htmlcompressor

type HtmlCompressor struct {
	Options map[string]bool
}

type Blocks struct {
	preBlocks         [][]byte
	taBlocks          [][]byte
	scriptBlocks      [][]byte
	styleBlocks       [][]byte
	eventBlocks       [][]byte
	condCommentBlocks [][]byte
	skipBlocks        [][]byte
	lineBreakBlocks   [][]byte
	userBlocks        [][]byte
}

func InitBlocks() *Blocks {
	return &Blocks{
		preBlocks:         make([][]byte, 0),
		taBlocks:          make([][]byte, 0),
		scriptBlocks:      make([][]byte, 0),
		styleBlocks:       make([][]byte, 0),
		eventBlocks:       make([][]byte, 0),
		condCommentBlocks: make([][]byte, 0),
		skipBlocks:        make([][]byte, 0),
		lineBreakBlocks:   make([][]byte, 0),
		userBlocks:        make([][]byte, 0),
	}
}

func Init() *HtmlCompressor {
	DefaultOptions := make(map[string]bool)
	DefaultOptions["enabled"] = true
	return &HtmlCompressor{Options: DefaultOptions}
}

func (self *HtmlCompressor) Compress(html []byte) []byte {
	if !self.Options["enabled"] || html == nil || len(html) == 0 {
		return html
	}
	blocks := InitBlocks()
	html = self.preserveBlocks(html, blocks)
	return []byte{}
}

func (self *HtmlCompressor) preserveBlocks(html []byte, blocks *Blocks) []byte {
	return html
}
