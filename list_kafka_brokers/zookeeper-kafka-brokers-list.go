package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// struct to contain Unmarshalled JSON
type broker struct {
	Listener_Security_Protocol_Map map[string]string
	Endpoints                      []string
	Jmx_Port                       int
	Host                           string
	Timestamp                      string
	Port                           int
	Version                        int
}

// This example we are going to grab a list of all of the kafka brokers from zookeeper.
func main() {
	// Here we connect a set of zookeeper nodes. We set our timeout to 3 secounds.
	connection, _, err := zk.Connect([]string{"localhost"}, time.Second*3)
	if err != nil {
		panic(err)
	}

	// Grab a slice of broker ids
	children, _, err := connection.Children("/brokers/ids")
	if err != nil {
		panic(err)
	}

	var brokers []broker
	// range over the broker ids to grab their information.
	for _, b := range children {
		node, _, err := connection.Get("/brokers/ids/" + b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", string(node))

		// Create a new instance of the broker struct.
		var brokerstruct broker
		// Unmarshal the json into this struct
		err = json.Unmarshal(node, &brokerstruct)
		if err != nil {
			fmt.Printf("%+v\n", err)
		}
		// Append this to our slice of brokers
		brokers = append(brokers, brokerstruct)
	}

	// Range over our slice of brokers printing out all of the brokers.
	for _, i := range brokers {
		fmt.Printf("%+v\n", i)
	}
}
