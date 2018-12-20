package main

import (
	"fmt"
	"strconv"
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

	// Store an int value for the binary number 11111 (turn everything on)
	perms, _ := strconv.ParseInt("11111", 2, 64)
	// Create a World readable/writable ACL
	acl := zk.WorldACL(int32(perms))
	// A byte slice of data to store.
	data := []byte("Testing")
	// Create znode
	connection.Create("/tester", data, 0, acl)

	// Print out the znodes
	children, _, err := connection.Children("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(children)
	// Print out the ACL we just created.
	fmt.Println(connection.GetACL("/tester"))
}
