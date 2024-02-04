package main

import (
	"github.com/riny/go-rpc/service"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// register rpc service
	err := rpc.RegisterName("UserInfoService", new(service.UserInfoService))
	if err != nil {
		panic("register rpc service error: " + err.Error())
	}

	http.HandleFunc("/rpc", func(writer http.ResponseWriter, request *http.Request) {
		var connection io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			Writer:     writer,
			ReadCloser: request.Body,
		}

		err := rpc.ServeRequest(jsonrpc.NewServerCodec(connection))
		if err != nil {
			panic("register rpc request server error: " + err.Error())
		}
	})

	err = http.ListenAndServe(":60001", nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
