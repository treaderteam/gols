package render

// FileGetter provide file accesing mechanizm
type FileGetter interface {
	GetFile(name string) ([]byte, error)
}
