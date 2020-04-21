package main

import (
	"context"
	"log"
	"time"

	api "github.com/ze0nni/grpc-validator/api"
	"google.golang.org/grpc"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewValidatorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	strt := &api.Struct{
		Foo: "foo",
		Bar: "bar",
		Baz: "baz",
	}
	filter := make(map[string]string)
	filter["foo"] = "foo"
	filter["bar"] = "bar"
	filter["baz"] = "baz"

	r, err := c.A(ctx, &api.ValidateRequest{Struct: strt, Filter: filter})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Result: %t", r.GetSuccess())
}
