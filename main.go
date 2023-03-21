package main

import (
	"net/http"

	"github.com/andangrd/testiing/internal"
	"github.com/gocraft/web"
)

type Context struct {
	InternalServices internal.Handler
}

func NewContext(svc *internal.HandlerImpl) Context {
	return Context{
		InternalServices: svc,
	}
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	next(rw, req)
}

func main() {
	internalSvc := internal.ProvideHandler(3)
	NewCtx := NewContext(internalSvc)
	router := web.New(NewCtx). // Create your router
					Middleware(web.LoggerMiddleware).                       // Use some included middleware
					Middleware(web.ShowErrorsMiddleware).                   // ...
					Middleware((*Context).SetHelloCount).                   // Your own middleware!
					Get("/test/:counter", NewCtx.InternalServices.SayHello) // Add a route
	http.ListenAndServe("localhost:3000", router) // Start the server!
}
