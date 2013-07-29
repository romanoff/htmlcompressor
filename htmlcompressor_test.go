package htmlcompressor

import "testing"

func TestCompress(t *testing.T) {
	compressor := Init()
	html := []byte{}
	if string(compressor.Compress(html)) != string(html) {
		t.Error("Empty html expected after compression, but got '%v'", string(compressor.Compress(html)))
	}
}

func TestRemoveCommnets(t *testing.T) {
	compressor := Init()
	html := []byte(`
<!---->
<!-- asdfasdf asdfsda -->
<div>
<!-- Some comment -->
asdf
</div>
`)
	expected := `


<div>

asdf
</div>
`
	if expected != string(compressor.removeComments(html)) {
		t.Errorf("Expected html with stripped comments:\n%v\n,but got:\n%v\n", expected, string(compressor.removeComments(html)))
	}
}
