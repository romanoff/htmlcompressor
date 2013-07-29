package htmlcompressor

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestCompress(t *testing.T) {
	compressor := Init()
	html := []byte{}
	if string(compressor.Compress(html)) != string(html) {
		t.Error("Empty html expected after compression, but got '%v'", string(compressor.Compress(html)))
	}
}

func testFromFile(t *testing.T, name string, compressor *HtmlCompressor) {
	source, err := ioutil.ReadFile("test_resources/" + name + ".html")
	if err != nil {
		t.Errorf("File not found: %v", "test_resources/"+name+".html")
	}
	expected, err := ioutil.ReadFile("test_resources/" + name + "Result.html")
	if err != nil {
		t.Errorf("File not found: %v", "test_resources/"+name+"Result.html")
	}
	expectedString := strings.TrimSpace(string(expected))
	result := strings.TrimSpace(string(compressor.Compress(source)))
	if result != expectedString {
		t.Errorf("Expected:\n%v\n, but got:\n%v\n", expectedString, result)
	}
}

func TestRemoveComments(t *testing.T) {
	compressor := Init()
	compressor.RemoveComments = true
	testFromFile(t, "testRemoveComments", compressor)
}

func TestSimpleDoctype(t *testing.T) {
	compressor := Init()
	compressor.SimpleDoctype = true
	testFromFile(t, "testSimpleDoctype", compressor)
}

func TestRemoveScriptAttributes(t *testing.T) {
	compressor := Init()
	compressor.RemoveScriptAttributes = true
	// testFromFile(t, "testRemoveScriptAttributes", compressor)
}

func TestRemoveIntertagSpaces(t *testing.T) {
	compressor := Init()
	compressor.RemoveIntertagSpaces = true
	testFromFile(t, "testRemoveIntertagSpaces", compressor)
}

func TestRemoveMultiSpaces(t *testing.T) {
	compressor := Init()
	compressor.RemoveMultiSpaces = true
	testFromFile(t, "testRemoveMultiSpaces", compressor)
}
