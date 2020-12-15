# Redfish and Swordfish client library

fork of https://github.com/stmcginnis/gofish

[![Go Doc](https://godoc.org/github.com/jacobweinstock/gophish?status.svg)](http://godoc.org/github.com/jacobweinstock/gophish)
[![Go Report Card](https://goreportcard.com/badge/github.com/jacobweinstock/gophish?branch=master)](https://goreportcard.com/report/github.com/jacobweinstock/gophish)
[![Releases](https://img.shields.io/github/release/jacobweinstock/redfish/all.svg?style=flat-square)](https://github.com/jacobweinstock/gophish/releases)
[![LICENSE](https://img.shields.io/github/license/jacobweinstock/redfish.svg?style=flat-square)](./LICENSE)

## Introduction

gophish is a Golang library for interacting with [DMTF Redfish](https://www.dmtf.org/standards/redfish) and [SNIA Swordfish](https://www.snia.org/forums/smi/swordfish) enabled devices.
For the moment, the goal of this repo is to stay up to date with https://github.com/stmcginnis/gofish.

## Usage

Basic usage:

```go

package main

import (
    "fmt"

    "github.com/jacobweinstock/gophish"
)

func main() {
    config := gophish.ClientConfig{
        Endpoint: "https://localhost:5000",
        Username: "admin",
        Password: "admin",
        Insecure: true,
    }
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    c, err := gophish.Connect(ctx, config)
    if err != nil {
        panic(err)
    }
    defer c.Logout(ctx)

    // Query the chassis data using the session token
    chassis, err := c.Service.Systems(ctx)
    if err != nil {
        panic(err)
    }

    for _, chass := range chassis {
        fmt.Println("Chassis:", chass.PowerState)
    }
}
```
