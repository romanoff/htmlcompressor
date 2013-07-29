package htmlcompressor

import "testing"

func TestCompress(t *testing.T) {
	compressor := Init()
	html := []byte{}
	if string(compressor.Compress(html)) != string(html) {
		t.Error("Empty html expected after compression, but got '%v'", string(compressor.Compress(html)))
	}
}
