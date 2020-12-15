package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jacobweinstock/gophish"
	"github.com/jacobweinstock/gophish/redfish"
)

func reboot() {
	// Create a new instance of redfish client, ignoring self-signed certs
	config := gophish.ClientConfig{
		Endpoint: "https://bmc-ip",
		Username: "my-username",
		Password: "my-password",
		Insecure: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	c, err := gophish.Connect(ctx, config)
	if err != nil {
		panic(err)
	}
	defer c.Logout(ctx)

	// Attached the client to service root
	service := c.Service

	// Query the computer systems
	ss, err := service.Systems(ctx)
	if err != nil {
		panic(err)
	}

	// Creates a boot override to pxe once
	bootOverride := redfish.Boot{
		BootSourceOverrideTarget:  redfish.PxeBootSourceOverrideTarget,
		BootSourceOverrideEnabled: redfish.OnceBootSourceOverrideEnabled,
	}

	for _, system := range ss {
		fmt.Printf("System: %#v\n\n", system)
		err := system.SetBoot(ctx, bootOverride)
		if err != nil {
			panic(err)
		}
		err = system.Reset(ctx, redfish.ForceRestartResetType)
		if err != nil {
			panic(err)
		}
	}
}
