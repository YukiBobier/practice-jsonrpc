package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &struct{ A, B int }{7, 8}

	var reply string
	err2 := client.Call("Arith.Add", args, &reply)
	if err2 != nil {
		log.Fatal("arith error:", err2)
	}

	fmt.Printf("Arith: %d+%d=%v\n", args.A, args.B, reply)
}
