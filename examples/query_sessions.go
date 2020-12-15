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

func querySessions() {
	// Create a new instance of redfish client, ignoring self-signed certs
	config := redfish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c, err := redfish.Connect(ctx, config)
	if err != nil {
		panic(err)
	}
	defer c.Logout(ctx)

	// Retrieve the service root
	service := c.Service

	// Query the active sessions using the session token
	sessions, err := service.Sessions(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", sessions)

	for _, session := range sessions {
		fmt.Printf("Sessions: %#v\n\n", session)
	}
}
