package render

import (
	"encoding/base64"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// Render embeds all in one html file
func Render(href string, fgetter FileGetter) ([]byte, error) {
	return render(href, fgetter)
}

type tagInfo struct {
	name      string
	hasImage  bool
	imageName string
	imgAttr   string
	image     string
	content   string
	start     int
	end       int
}

type tagsInfo []tagInfo

func (t tagsInfo) String() string {
	result := ""
	for _, v := range t {
		result += v.name + " " + strconv.Itoa(v.start) + " " + strconv.Itoa(v.end) + " \n"
	}

	return result
}

func render(href string, fgetter FileGetter) ([]byte, error) {
	result := make([]byte, 0)
	file, err := fgetter.GetFile(href)
	if err != nil {
		return nil, err
	}

	tags := tagReader(file)
	tags = cutBody(tags)
	tags = cutImages(tags)

	result, err = renderImages(tags, file, fgetter)

	return result, err
}

func tagReader(file []byte) tagsInfo {
	tagStarted := false
	currTag := ""
	tagStart := 0
	tagText := make([]byte, 0)
	result := make([]tagInfo, 0)

	for i, v := range file {
		switch v {
		case '<':
			if currTag != "" {
				result = append(result, tagInfo{
					name:    currTag,
					start:   tagStart,
					end:     i - 1,
					content: string(tagText),
				})

				currTag = ""
				tagText = make([]byte, 0)
			}
			tagStarted = true
			tagStart = i
			break
		case '>':
			tagStarted = false
			break
		default:
			if tagStarted {
				currTag += string(v)
			} else {
				tagText = append(tagText, v)
			}
		}
	}

	return result
}

func cutBody(tags tagsInfo) tagsInfo {
	result := make(tagsInfo, 0)
	bodyStarted := false

	for _, v := range tags {
		if strings.HasPrefix(v.name, "body") {
			bodyStarted = true
			continue
		}

		if strings.HasPrefix(v.name, "/body") {
			break
		}

		if bodyStarted {
			result = append(result, v)
		}
	}

	return result
}

func cutImages(tags tagsInfo) tagsInfo {
	result := make(tagsInfo, 0)
	rex := regexp.MustCompile("(xlink:href=|src=)\".+\\.(jpg)\"")
	imgRex := regexp.MustCompile("\".+(jpg)\"")

	for _, v := range tags {
		for _, k := range strings.Split(v.name, " ") {
			if rex.MatchString(k) {
				imgName := strings.Replace(imgRex.FindString(k), "\"", "", 2)
				v.imgAttr = rex.FindString(k)
				v.imageName = imgName
				v.hasImage = true
				break
			}
		}

		result = append(result, v)
	}

	return result
}

func renderImages(tags tagsInfo, file []byte, fgetter FileGetter) ([]byte, error) {
	result := make([]byte, 0)
	imgRex := regexp.MustCompile("\".+(jpg)\"")

	for _, v := range tags {
		if v.hasImage {
			img, err := fgetter.GetFile(v.imageName)
			if err != nil {
				log.Println(err)
				return nil, err
			}

			img64 := base64.StdEncoding.EncodeToString(img)
			imgAttr := imgRex.ReplaceAllString(v.imgAttr, "\"data:image/jpeg;base64,"+img64+"\"")

			el := strings.Replace(v.name, v.imgAttr, imgAttr, 1)
			el = "<" + el + ">"
			result = append(result, []byte(el)...)
		} else {
			result = append(result, file[v.start:v.end+1]...)
		}

		if v.content != "" {
			// result = append(result, []byte("gangster"+v.content)...)
		}
	}

	return result, nil
}
