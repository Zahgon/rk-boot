// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	_ "embed"
	"net/http"

	rkboot "github.com/rookie-ninja/rk-boot/v2"
)

// @title RK Swagger for Mux
// @version 1.0
// @description This is a greeter service with rk-boot.
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Get MuxEntry
	muxEntry := rkmux.GetMuxEntry("greeter")
	// Use *mux.Router adding handler.
	muxEntry.Router.NewRoute().Path("/v1/greeter").HandlerFunc(Greeter)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

// Greeter handler
// @Summary Greeter service
// @Id 1
// @version 1.0
// @produce application/json
// @Param name query string true "Input name"
// @Success 200 {object} GreeterResponse
// @Router /v1/greeter [get]
func Greeter(writer http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

type GreeterResponse struct {
	Message string
}
