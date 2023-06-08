package main

import (
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	log.Fatal(http.ListenAndServe(":1234", nil))
}

type Arith struct{}

func (*Arith) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

type Args struct {
	A, B int
}
