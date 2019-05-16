package server

import "net/http"

type ICodecRequest interface {
}

type ICodec interface {
	Wrap(*http.Request) ICodecRequest
}

type Server struct {
	codecs map[string]ICodec
	//services *serviceMap
}

func NewServer() *Server {
	return &Server{
		codecs: make(map[string]ICodec),
		//services: new(serviceMap)
	}
}
