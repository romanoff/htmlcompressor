package htmlcompressor

type HtmlCompressor struct {
	Options map[string]bool
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
	return []byte{}
}
