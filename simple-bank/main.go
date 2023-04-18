package main

import (
	_ "github.com/lib/pq"
	"github.com/rizkyfazri23/mnc/delivery"
)

func main() {
	// Run the server
	delivery.Server().Run()
}
