package greeter

type greeter struct {
	reader Reader
}

type GreeterService interface {
	Greet(lang string) string
	GreetWithDefaultMessage() string
}

func NewGreeterService(reader Reader) GreeterService {
	service := greeter{reader}
	return &service
}

func (g *greeter) Greet(lang string) string {

	msg, _ := g.reader.GetMessage(lang)
	return msg
}

func (g *greeter) GreetWithDefaultMessage() string {

	msg, _ := g.reader.GetDefaultMessage()
	return msg
}
