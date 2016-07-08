package main

import (
	"log"

	"github.com/achiku/testsvr"
	"github.com/achiku/testsvr/example"
)

func main() {
	s := testsvr.NewServer(example.DefaultHandlerMap, "8181")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
