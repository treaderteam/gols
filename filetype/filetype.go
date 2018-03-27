// Package filetype is for detecting filetype
package filetype

import (
	"io"
	"log"

	"gitlab.com/alexnikita/gols/filetype/types"
)

// Detect get type of file
func Detect(file []byte) types.Type {

	switch true {
	case isEpub(file):
		return types.EPUB
	case isPDF(file):
		return types.PDF
	case isFB2(file):
		return types.FB2
	}

	return types.Unknown
}

// DetectFromReader get type from reader
func DetectFromReader(rdr io.Reader) (types.Type, []byte, error) {
	buf := make([]byte, 58)
	readed, err := rdr.Read(buf)
	if err != nil {
		log.Println(err, readed)
		return types.Unknown, nil, err
	}

	return Detect(buf), buf, nil
}

// check if file is epub
func isEpub(buf []byte) bool {
	return len(buf) > 57 &&
		buf[0] == 0x50 && buf[1] == 0x4B && buf[2] == 0x3 && buf[3] == 0x4 &&
		buf[30] == 0x6D && buf[31] == 0x69 && buf[32] == 0x6D && buf[33] == 0x65 &&
		buf[34] == 0x74 && buf[35] == 0x79 && buf[36] == 0x70 && buf[37] == 0x65 &&
		buf[38] == 0x61 && buf[39] == 0x70 && buf[40] == 0x70 && buf[41] == 0x6C &&
		buf[42] == 0x69 && buf[43] == 0x63 && buf[44] == 0x61 && buf[45] == 0x74 &&
		buf[46] == 0x69 && buf[47] == 0x6F && buf[48] == 0x6E && buf[49] == 0x2F &&
		buf[50] == 0x65 && buf[51] == 0x70 && buf[52] == 0x75 && buf[53] == 0x62 &&
		buf[54] == 0x2B && buf[55] == 0x7A && buf[56] == 0x69 && buf[57] == 0x70
}

// check if file is pdf
func isPDF(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x25 && buf[1] == 0x50 &&
		buf[2] == 0x44 && buf[3] == 0x46
}

// check if file is fb2
func isFB2(buf []byte) bool {
	if len(buf) < 51 {
		return false
	}
	return string(buf[40:51]) == "FictionBook"
}
