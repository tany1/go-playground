package main

import "fmt"

type OptsFunc func(*Options)

type Options struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Options {
	return Options{
		maxConn: 1,
		id:      "default",
		tls:     false,
	}
}

type Server struct {
	Opts Options
}

func withTls(o *Options) {
	o.tls = true
}

func withMaxConn(maxConn int) OptsFunc {
	return func(o *Options) {
		o.maxConn = maxConn
	}
}

func withId(id string) OptsFunc {
	return func(o *Options) {
		o.id = id
	}
}

func NewServer(opts ...OptsFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts: o,
	}
}

func main() {
	server := NewServer(withTls)
	// server := NewServer(withTls, withMaxConn(10), withId("server-1"))
	// server := NewServer(withTls, withMaxConn(10), withId("server-1"))

	fmt.Printf("%+v\n", server)
}
