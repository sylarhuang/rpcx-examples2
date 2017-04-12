package main

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/smallnest/rpcx"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972", DialTimeout: 10 * time.Second}
	client := rpcx.NewClient(s)

	args := &Args{7, 8}
	var reply Reply

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Call(ctx, "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	}

	client.Close()
}