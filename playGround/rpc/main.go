package main

import "net/rpc"

func main() {
	rpc.RegisterName("helloService", new(HelloService))
}
