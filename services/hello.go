package services

// HelloService contains methods for finding and manipulating resources
type HelloService struct {
}

// NewHelloService creates an instance of HelloService
func NewHelloService() *HelloService {
	return &HelloService{}
}

// SayHello say hello
func (hs HelloService) SayHello(name string) string {
	return "Hello " + name
}
