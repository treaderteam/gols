package epub

import "gitlab.com/alexnikita/gols/epub/render"

// Render embeds all assets into html page
func (b *Book) Render(href string) (result []byte, err error) {
	return render.Render(href, b)
}
