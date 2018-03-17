package filetype_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/alexnikita/gols/filetype"

	"gitlab.com/alexnikita/gols/filetype/types"
)

func TestDetectFiletype(t *testing.T) {
	files := []struct {
		name string
		typ  types.Type
	}{
		{
			name: "./testfiles/test_epub.epub",
			typ:  types.EPUB,
		},
		{
			name: "./testfiles/test_pdf.pdf",
			typ:  types.PDF,
		},
		{
			name: "./testfiles/test_fb2.fb2",
			typ:  types.FB2,
		},
	}

	for _, v := range files {
		file, err := os.Open(v.name)
		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		data, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatal(err)
		}

		typ := filetype.Detect(data)
		assert.Equal(t, v.typ, typ)
	}
}
