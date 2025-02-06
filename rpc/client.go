package main

import (
	"fmt"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer client.Close()

	args := Args{A: 5, B: 3}
	var reply int

	err = client.Call("Calculator.Add", &args, &reply)
	if err != nil {
		fmt.Println("RPC call error:", err)
		return
	}

	fmt.Println("Result of 5 + 3 =", reply)
}
