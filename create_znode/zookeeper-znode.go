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

	//
	perms, _ := strconv.ParseInt("11111", 2, 64)
	acl := zk.WorldACL(int32(perms))
	data := []byte("Testing")
	connection.Create("/tester", data, 0, acl)

	children, _, err := connection.Children("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(children)
	fmt.Println(connection.GetACL("/tester"))
}
