package htmlcompressor

import (
	"regexp"
)

type HtmlCompressor struct {
	Enabled                bool
	RemoveComments         bool
	SimpleDoctype          bool
	RemoveScriptAttributes bool
	RemoveIntertagSpaces   bool
	RemoveMultiSpaces      bool
}

func Init() *HtmlCompressor {
	compressor := &HtmlCompressor{
		Enabled:           true,
		RemoveComments:    true,
		RemoveMultiSpaces: true,
	}
	return compressor
}

func (self *HtmlCompressor) Compress(html []byte) []byte {
	if !self.Enabled || html == nil || len(html) == 0 {
		return html
	}
	html = self.processHtml(html)
	html = self.simpleDoctype(html)
	html = self.removeScriptAttributes(html)
	html = self.removeIntertagSpaces(html)
	html = self.removeMultiSpaces(html)
	html = self.removeSpacesInsideTags(html)
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

var jsTypeAttrPattern *regexp.Regexp = regexp.MustCompile(`(?is)(<script[^>]*)type\s*=\s*([\"']*)(?:text|application)\/javascript([\"']*)([^>]*>)`)
var jsLangAttrPattern *regexp.Regexp = regexp.MustCompile(`(?is)(<script[^>]*)language\s*=\s*([\"']*)javascript([\"']*)([^>]*>)`)

//Replaced \2 with [\"'] (as closing tag). This way check of proper closing quote is abandoned
func (self *HtmlCompressor) removeScriptAttributes(html []byte) []byte {
	if self.RemoveScriptAttributes {
		html = jsTypeAttrPattern.ReplaceAll(html, []byte("$1$4"))
		html = jsLangAttrPattern.ReplaceAll(html, []byte("$1$4"))
	}
	return html
}

var intertagPatternTagTag *regexp.Regexp = regexp.MustCompile(`(?is)>\s+<`)
var intertagPatternTagCustom *regexp.Regexp = regexp.MustCompile(`(?is)>\s+%%%~`)
var intertagPatternCustomTag *regexp.Regexp = regexp.MustCompile(`(?is)~%%%\s+<`)
var intertagPatternCustomCustom *regexp.Regexp = regexp.MustCompile(`(?is)~%%%\s+%%%~`)

func (self *HtmlCompressor) removeIntertagSpaces(html []byte) []byte {
	if self.RemoveIntertagSpaces {
		html = intertagPatternTagTag.ReplaceAll(html, []byte("><"))
		html = intertagPatternTagCustom.ReplaceAll(html, []byte(">%%%~"))
		html = intertagPatternCustomTag.ReplaceAll(html, []byte("~%%%<"))
		html = intertagPatternCustomCustom.ReplaceAll(html, []byte("~%%%%%%~"))
	}
	return html
}

var multiSpacePattern *regexp.Regexp = regexp.MustCompile(`(?is)\s+`)

func (self *HtmlCompressor) removeMultiSpaces(html []byte) []byte {
	if self.RemoveMultiSpaces {
		html = multiSpacePattern.ReplaceAll(html, []byte(" "))
	}
	return html
}

//var tagPropertyPattern *regexp.Regexp = regexp.MustCompile(`(?i)(\s\w+)\s*=\s*(?=[^<]*?>)`)
//Cannot be used as go regexp doesn't have ?=
var tagPropertyPattern *regexp.Regexp = regexp.MustCompile(`(?i)(\s\w+)\s*=\s*`)
var tagEndSpacePattern *regexp.Regexp = regexp.MustCompile(`(?is)(<(?:[^>]+?))(?:\s+?)(/?>)`)

func (self *HtmlCompressor) removeSpacesInsideTags(html []byte) []byte {
	html = tagPropertyPattern.ReplaceAll(html, []byte("$1="))
	html = tagEndSpacePattern.ReplaceAll(html, []byte("$1$2"))
	return html
}
