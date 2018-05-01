package epub_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"gitlab.com/alexnikita/gols/epub"
)

func TestRender(t *testing.T) {
	filename := "testfiles/1.epub"

	book, err := epub.Open(filename)
	if err != nil {
		t.Fatal(err)
	}

	rendered := make([]byte, 0)

	for _, v := range book.Opf.Spine.Items {
		href := ""
		for _, k := range book.Opf.Manifest {
			if k.ID == v.IDref {
				href = k.Href
				break
			}
		}
		part, err := book.Render(href)
		if err != nil {
			t.Fatal(err)
		}

		rendered = append(rendered, part...)
	}

	head := []byte(`
		<html>
			<head>
				<meta charset='UTF-8'>
			</head>
			<body>

	`)
	foot := []byte(`
		</body>
		</html>	
	`)

	rendered = append(head, rendered...)
	rendered = append(rendered, foot...)

	// create html file, which contains full book
	os.Remove("index.html")

	file, err := os.OpenFile("index.html", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		t.Fatal(err)
	}

	io.Copy(file, bytes.NewReader([]byte(string(rendered))))
}
