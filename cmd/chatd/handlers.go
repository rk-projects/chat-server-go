package main

import (
	"bufio"
	"io"
	"net"
)

//connHandler is required to process data
type connHandler struct{}

//bind is called to init a reader and writer
func (connHandler) Bind(conn net.Conn) (io.Reader, io.Writer) {
	return conn, bufio.NewWriter(conn)
}

//reqHandler is required to process client messages
type reqHandler struct{}

//read implements tcp handler interface. INPUT : request and io.Reader from bind
func (reqHandler) Read(ipAddress string, reader io.Reader) ([]byte, int, error) {

	//block on network for our message
	data, n, err := msg.Read(reader)
}
