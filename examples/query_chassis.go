//
// SPDX-License-Identifier: BSD-3-Clause
//
package main

import (
	"context"
	"fmt"
	"time"

	redfish "github.com/jacobweinstock/gophish"
)

func queryChassis() {
	// Create a new instance of redfish client, ignoring self-signed certs
	config := redfish.ClientConfig{
		Endpoint: "https://localhost:5000",
		Username: "admin",
		Password: "admin",
		Insecure: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	c, err := redfish.Connect(ctx, config)
	if err != nil {
		panic(err)
	}
	defer c.Logout(ctx)

	// Retrieve the service root

	// Query the chassis data using the session token
	chassis, err := c.Service.Systems(ctx)
	if err != nil {
		panic(err)
	}

	for _, chass := range chassis {
		fmt.Println("Chassis:", chass.PowerState)
	}
}
