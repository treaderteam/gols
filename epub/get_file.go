package epub

import (
	"io"
	"io/ioutil"
	"strings"
)

// GetFile get bytes of file
func (p *Book) GetFile(name string) (result []byte, err error) {
	var file io.ReadCloser
	name = normalizeFilename(name)
	for _, v := range p.fd.File {
		if v.Name == name {
			file, err = v.Open()
			if err != nil {
				return
			}
			break
		} else if "OEBPS/"+name == v.Name {
			file, err = v.Open()
			if err != nil {
				return
			}
			break
		} else if "OPS/"+name == v.Name {
			file, err = v.Open()
			if err != nil {
				return
			}
			break
		}
	}

	if file == nil {
		return
	}

	defer file.Close()

	result, err = ioutil.ReadAll(file)
	return
}

func normalizeFilename(name string) string {
	if strings.HasPrefix(name, "../") {
		name = strings.Replace(name, "../", "", 1)
	}

	return name
}
