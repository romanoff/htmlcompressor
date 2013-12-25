package htmlcompressor

import (
	"bytes"
	"regexp"
	"strconv"
)

type HtmlCompressor struct {
	Enabled                bool
	RemoveComments         bool
	SimpleDoctype          bool
	RemoveScriptAttributes bool
	RemoveIntertagSpaces   bool
	RemoveMultiSpaces      bool
	tempBlocks             map[string][][]byte
}

func Init() *HtmlCompressor {
	compressor := &HtmlCompressor{
		Enabled:           true,
		RemoveComments:    true,
		RemoveMultiSpaces: true,
		tempBlocks:        make(map[string][][]byte),
	}
	return compressor
}

func InitAll() *HtmlCompressor {
	compressor := &HtmlCompressor{
		Enabled:                true,
		RemoveComments:         true,
		SimpleDoctype:          true,
		RemoveScriptAttributes: true,
		RemoveIntertagSpaces:   true,
		RemoveMultiSpaces:      true,
		tempBlocks:             make(map[string][][]byte),
	}
	return compressor
}

func (self *HtmlCompressor) Compress(html []byte) []byte {
	if !self.Enabled || html == nil || len(html) == 0 {
		return html
	}
	html = self.preserveBlocks(html)
	html = self.processHtml(html)
	html = self.simpleDoctype(html)
	html = self.removeScriptAttributes(html)
	html = self.removeIntertagSpaces(html)
	html = self.removeMultiSpaces(html)
	html = self.removeSpacesInsideTags(html)
	html = self.restorePreservedBlocks(html)
	return html
}

var prePattern *regexp.Regexp = regexp.MustCompile(`(?is)(<pre[^>]*?>)(.*?)(</pre>)`)

func (self *HtmlCompressor) preserveBlocks(html []byte) []byte {
	tempPreBlock := "%%%~COMPRESS~PRE~{0,number,#}~%%%"
	matches := prePattern.FindAllSubmatch(html, -1)
	i := 0
	for _, match := range matches {
		content := match[2]
		if len(content) > 0 {
			tempBlock := messageFormat(tempPreBlock, i)
			i++
			replaceWith := []byte{}
			replaceWith = append(replaceWith, match[1]...)
			replaceWith = append(replaceWith, tempBlock...)
			replaceWith = append(replaceWith, match[3]...)
			self.tempBlocks[tempPreBlock] = append(self.tempBlocks[tempPreBlock], match[2])
			html = bytes.Replace(html, match[0], replaceWith, -1)
		} else {
			bytes.Replace(html, match[0], []byte{}, -1)
		}
	}
	return html
}

func (self *HtmlCompressor) restorePreservedBlocks(html []byte) []byte {
	for blockName, blocks := range self.tempBlocks {
		for i, block := range blocks {
			tempBlock := messageFormat(blockName, i)
			html = bytes.Replace(html, tempBlock, block, -1)
		}
	}
	return html
}

var messagePattern *regexp.Regexp = regexp.MustCompile(`(.+){0,number,#}(.+)`)

func messageFormat(message string, i int) []byte {
	return []byte(messagePattern.ReplaceAllString(message, "$1{"+strconv.Itoa(i)+"}$2"))
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
