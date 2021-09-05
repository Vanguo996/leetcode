package main

type HelloService struct {}

func (p *HelloService) hello(req string, reply *string) error {
	*reply = "hello" + req
	return nil
}
