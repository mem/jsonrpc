// Copyright 2015 Marcelo E. Magallon <marcelo.magallon@gmail.com>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jsonrpc

import (
	"io"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type serverCodec struct {
	rpc.ServerCodec
	mapping map[string]string
}

// NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser, mapping map[string]string) rpc.ServerCodec {
	return &serverCodec{
		ServerCodec: jsonrpc.NewServerCodec(conn),
		mapping:     mapping,
	}
}

func (c *serverCodec) ReadRequestHeader(r *rpc.Request) error {
	err := c.ServerCodec.ReadRequestHeader(r)
	if err != nil {
		return err
	}

	if _, exists := c.mapping[r.ServiceMethod]; exists {
		r.ServiceMethod = c.mapping[r.ServiceMethod]
	}

	return nil
}

// ServeConn runs the JSON-RPC server on a single connection.
// ServeConn blocks, serving the connection until the client hangs up.
// The caller typically invokes ServeConn in a go statement.
func ServeConn(conn io.ReadWriteCloser, mapping map[string]string) {
	rpc.ServeCodec(NewServerCodec(conn, mapping))
}
