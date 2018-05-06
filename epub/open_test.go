package epub_test

import (
	"testing"

	"gitlab.com/alexnikita/gols/epub"
)

func TestOpenEpub(t *testing.T) {
	var (
		folder   = "testfiles/"
		filename = "2.epub"
	)

	_, err := epub.Open(folder + filename)
	if err != nil {
		t.Fatal(err)
	}

	// log.Println(res)
}
