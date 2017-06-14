package main

import (
	"log"

	restful "github.com/emicklei/go-restful"
)

func logger(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	// log request data
	log.Printf("Incoming %s request to %s from %s\n",
		req.Request.Method,
		req.Request.RequestURI,
		req.Request.RemoteAddr,
	)

	// proceed to next filter
	chain.ProcessFilter(req, resp)
}
