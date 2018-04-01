// Package fb2 represent .fb2 format parser
package fb2

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"

	"gitlab.com/alexnikita/gols/xmlreader"
)

// Parser struct
type Parser struct {
	book   []byte
	reader io.Reader
}

// GetWords from fb2
func (p *Parser) GetWords() (map[string]string, error) {
	var (
		data []byte
		err  error
	)

	data = p.book

	if p.reader != nil {
		data, err = ioutil.ReadAll(p.reader)
		if err != nil {
			return nil, err
		}
	}

	words, err := xmlreader.GetWordsFromXMLBody(data)

	if err != nil {
		return nil, err
	}

	result := make(map[string]string)

	for _, v := range words {
		result[v] = v
	}

	return result, nil
}

// New creates new Parser
func New(data []byte) *Parser {
	return &Parser{
		book: data,
	}
}

// NewReader creates new Parser from reader
func NewReader(data io.Reader) *Parser {
	return &Parser{
		reader: data,
	}
}

// CharsetReader required for change encodings
func (p *Parser) CharsetReader(c string, i io.Reader) (r io.Reader, e error) {
	switch c {
	case "windows-1251":
		r = decodeWin1251(i)
	}
	return
}

// Unmarshal parse data to FB2 type
func (p *Parser) Unmarshal() (result FB2, err error) {
	if p.reader != nil {
		decoder := xml.NewDecoder(p.reader)
		decoder.CharsetReader = p.CharsetReader
		if err = decoder.Decode(&result); err != nil {
			return
		}

		return
	}
	reader := bytes.NewReader(p.book)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = p.CharsetReader
	if err = decoder.Decode(&result); err != nil {
		return
	}

	return
}
