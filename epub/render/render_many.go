package render

// RenderMany render multilpe files as one html
func RenderMany(hrefs []string, fgetter FileGetter) ([]byte, error) {
	return renderMany(hrefs, fgetter)
}

func renderMany(hrefs []string, fgetter FileGetter) ([]byte, error) {
	result := make([]byte, 0)
	for _, v := range hrefs {
		file, err := fgetter.GetFile(v)
		if err != nil {
			return nil, err
		}

		tags := tagReader(file)
		tags = cutBody(tags)
		tags = cutImages(tags)

		res, err := renderImages(tags, file, fgetter)
		if err != nil {
			return nil, err
		}

		result = append(result, res...)
	}

	result = appendHeadAndFoot(result)

	return result, nil
}
