package main

import "io"

import "log"

import "net"

import "os"

func testNetCat() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	MustCopy(conn, os.Stdin)
	// if conn, ok := conn.(*net.TCPConn); ok {
	// 	conn.CloseWrite()
	// }
	conn.Close()
	<-done
}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
