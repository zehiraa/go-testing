package greeter

type reader struct{}

type Reader interface {
	GetMessage(lang string) (string, error)
	GetDefaultMessage() (string, error)
}

func NewReader() Reader {
	reader := reader{}
	return &reader
}

func (r *reader) GetMessage(lang string) (string, error) {
	// in real life, this code will call an external db
	//but for this sample we will just return the hardcoded example value
	if lang == "en" {
		return "hello", nil
	}
	if lang == "tr" {
		return "merhaba", nil
	}

	return "...-_-...", nil
}

func (r *reader) GetDefaultMessage() (string, error) {
	return "default message", nil
}
