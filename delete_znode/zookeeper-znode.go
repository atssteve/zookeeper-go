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

	// List all of the znodes.
	children, _, err := connection.Children("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(children)

	// Delete the tester znode.
	connection.Delete("/tester", 0)

	// Print them again.
	children, _, err = connection.Children("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(children)
}
