package htmlcompressor

import (
	"regexp"
)

type HtmlCompressor struct {
	Enabled        bool
	RemoveComments bool
	SimpleDoctype  bool
}

func Init() *HtmlCompressor {
	compressor := &HtmlCompressor{
		Enabled:        true,
		RemoveComments: true,
	}
	return compressor
}

func (self *HtmlCompressor) Compress(html []byte) []byte {
	if !self.Enabled || html == nil || len(html) == 0 {
		return html
	}
	html = self.processHtml(html)
	html = self.simpleDoctype(html)
	return html
}

func (self *HtmlCompressor) processHtml(html []byte) []byte {
	html = self.removeComments(html)
	return html
}

var commentPattern *regexp.Regexp = regexp.MustCompile(`(?is)<!---->|<!--[^\\[].*?-->`)

func (self *HtmlCompressor) removeComments(html []byte) []byte {
	if self.RemoveComments {
		html = commentPattern.ReplaceAll(html, []byte{})
	}
	return html
}

var doctypePattern *regexp.Regexp = regexp.MustCompile(`(?is)<!DOCTYPE[^>]*>`)

func (self *HtmlCompressor) simpleDoctype(html []byte) []byte {
	if self.SimpleDoctype {
		html = doctypePattern.ReplaceAll(html, []byte("<!DOCTYPE html>"))
	}
	return html
}
