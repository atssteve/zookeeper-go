package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// This example we are going to grab a list of all of the kafka brokers from zookeeper.
func main() {
	// Here we connect a set of zookeeper nodes. We set our timeout to 3 secounds.
	connection, _, err := zk.Connect([]string{"localhost"}, time.Second*3)
	if err != nil {
		panic(err)
	}

	// Grab a slice of broker ids
	children, _, err := connection.Children("/")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", children)

}
