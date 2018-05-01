package epub

import "gitlab.com/alexnikita/gols/epub/render"

// Render embeds all assets into html page
func (b *Book) Render(href string) (result []byte, err error) {
	return render.Render(href, b)
}

// RenderMany render many files in one html
func (b *Book) RenderMany(hrefs []string) (result []byte, err error) {
	return render.RenderMany(hrefs, b)
}
