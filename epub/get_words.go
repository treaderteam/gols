package epub

import (
	"io/ioutil"
	"log"
	"strings"

	"gitlab.com/alexnikita/gols/xmlreader"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

// GetWords from all .html and .xhtml files
func (b Book) GetWords() (map[string]string, error) {
	result := make(map[string]string)

	for _, v := range b.Files() {
		if strings.HasSuffix(v, ".html") || strings.HasSuffix(v, ".xhtml") {
			fname := strings.Replace(v, "OEBPS/", "", 1)
			file, err := b.Open(fname)
			if err != nil {
				log.Println(err)
				continue
			}

			data, err := ioutil.ReadAll(file)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			file.Close()

			words, err := xmlreader.GetWordsFromXMLBody(data)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			for _, k := range words {
				result[k] = k
			}
		}
	}

	return result, nil
}

// func getWordsFromHTML()
