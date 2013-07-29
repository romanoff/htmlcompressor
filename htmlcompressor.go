package htmlcompressor

import (
	"regexp"
)

type HtmlCompressor struct {
	Options map[string]bool
}

func Init() *HtmlCompressor {
	DefaultOptions := make(map[string]bool)
	DefaultOptions["enabled"] = true
	DefaultOptions["remove_comments"] = true
	return &HtmlCompressor{Options: DefaultOptions}
}

func (self *HtmlCompressor) Compress(html []byte) []byte {
	if !self.Options["enabled"] || html == nil || len(html) == 0 {
		return html
	}
	html = self.processHtml(html)
	return []byte{}
}

func (self *HtmlCompressor) processHtml(html []byte) []byte {
	html = self.removeComments(html)
	return html
}

var commentPattern *regexp.Regexp = regexp.MustCompile(`(?is)<!---->|<!--[^\\[].*?-->`)

func (self *HtmlCompressor) removeComments(html []byte) []byte {
	if self.Options["remove_comments"] {
		html = commentPattern.ReplaceAll(html, []byte{})
	}
	return html
}
