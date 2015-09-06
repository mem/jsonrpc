// Copyright 2015 Marcelo E. Magallon <marcelo.magallon@gmail.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package jsonrpc implements a JSON-RPC ClientCodec and ServerCodec
// for the rpc package.
package jsonrpc

import (
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.
func NewClientCodec(conn io.ReadWriteCloser) rpc.ClientCodec {
	return jsonrpc.NewClientCodec(conn)
}

// NewClient returns a new rpc.Client to handle requests to the
// set of services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *rpc.Client {
	return jsonrpc.NewClient(conn)
}

// Dial connects to a JSON-RPC server at the specified network address.
func Dial(network, address string) (*rpc.Client, error) {
	return jsonrpc.Dial(network, address)
}
