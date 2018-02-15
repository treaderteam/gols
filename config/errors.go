package config

type InvalidFormatError struct{}

func (i InvalidFormatError) Error() string {
	return "invalid format error"
}
