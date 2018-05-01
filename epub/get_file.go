package epub

import (
	"io"
	"io/ioutil"
)

// GetFile get bytes of file
func (p *Book) GetFile(name string) (result []byte, err error) {
	var file io.ReadCloser
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
