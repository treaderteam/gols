// Package xmlreader provide reading operations
// on xml based files
package xmlreader

// GetWordsFromXMLBody implementation
func GetWordsFromXMLBody(data []byte) (words []string, err error) {
	// fmt.Printf("fb2 reader starts read %d bytes of data...\n", len(data))
	lines := 0
	tagOpened := false
	tagName := ""
	tags := []string{}
	bodyStartIndex := 0
	bodyEndIndex := 0
	for i, v := range data {
		if !tagOpened {
			if v == byte(60) {
				tagOpened = true
			}
		} else {
			if v == byte(62) {
				switch tagName {
				case "body":
					bodyStartIndex = i + 1
					break
				case "/body":
					bodyEndIndex = i - 7
					break
				}
				tagOpened = false
				tags = append(tags, tagName)
				tagName = ""
			} else {
				tagName += string(v)
			}
		}
		if v == byte(10) {
			lines++
		}
	}
	if words, err = readBody(data[bodyStartIndex:bodyEndIndex]); err != nil {
		return
	}

	// fmt.Printf("file contains %d lines, %d words\n", lines, len(words))
	// fmt.Printf("body starts at index %d, ends at index %d\n", bodyStartIndex, bodyEndIndex)

	return
}

// read body of .fb2 file
func readBody(body []byte) (words []string, err error) {
	wordStart := false
	word := ""
	tagStart := false
	for _, v := range body {
		// if tag was started, check next symbol for '<'
		if tagStart {
			if v == byte(62) {
				tagStart = false
			}
		} else {
			// if word was started
			if wordStart {
				// if symbol is a character
				if v == byte(60) {
					tagStart = true
				} else if v > 64 && v < 91 || v > 96 && v < 123 {
					word += string(v)
				} else {
					words = append(words, word)
					word = ""
					wordStart = false
				}
			} else {
				if v == byte(60) {
					tagStart = true
				} else if v > 64 && v < 91 || v > 96 && v < 123 {
					wordStart = true
					word += string(v)
				}
			}
		}
	}
	return
}
