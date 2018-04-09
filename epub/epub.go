package epub

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"path"
)

// Book epub book
type Book struct {
	Ncx       Ncx       `json:"ncx"`
	Opf       Opf       `json:"opf"`
	Container Container `json:"-"`
	Mimetype  string    `json:"-"`

	fd *zip.ReadCloser
}

//Open open resource file
func (p *Book) Open(n string) (io.ReadCloser, error) {
	return p.open(p.filename(n))
}

//Files list resource files
func (p *Book) Files() []string {
	var fns []string
	for _, f := range p.fd.File {
		fns = append(fns, f.Name)
	}
	return fns
}

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
		}
	}

	if file == nil {
		return
	}

	defer file.Close()

	result, err = ioutil.ReadAll(file)
	return
}

//-----------------------------------------------------------------------------
func (p *Book) filename(n string) string {
	return path.Join(path.Dir(p.Container.Rootfile.Path), n)
}

func (p *Book) readXML(n string, v interface{}) error {
	fd, err := p.open(n)
	if err != nil {
		return nil
	}
	defer fd.Close()
	dec := xml.NewDecoder(fd)
	return dec.Decode(v)
}

func (p *Book) readBytes(n string) ([]byte, error) {
	fd, err := p.open(n)
	if err != nil {
		return nil, nil
	}
	defer fd.Close()

	return ioutil.ReadAll(fd)

}

func (p *Book) open(n string) (io.ReadCloser, error) {
	for _, f := range p.fd.File {
		if f.Name == n {
			return f.Open()
		}
	}
	return nil, fmt.Errorf("file %s not exist", n)
}
