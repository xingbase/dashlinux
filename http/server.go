package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/xingbase/dashlinux"
)

type Server struct {
	Host string `long:"host" description:"The IP to listen on" default:"0.0.0.0" env:"HOST"`
	Port int    `long:"port" description:"The port to listen on for insecure connections, defaults to a random value" default:"8888" env:"PORT"`

	PprofEnabled bool `long:"pprof-enabled" description:"Enable the /debug/pprof/* HTTP routes" env:"PPROF_ENABLED"`

	TokenSecret string `short:"t" long:"token-secret" description:"Secret to sign tokens" env:"TOKEN_SECRET"`

	BuildInfo dashlinux.BuildInfo
}

func (s *Server) Serve(ctx context.Context) {
	jwt := jwtauth.New("HS256", []byte(s.TokenSecret), nil)

	service := Service{
		Jwt: jwt,
	}
	handler := NewMux(MuxOpts{
		PprofEnabled: s.PprofEnabled,
	}, service)

	httpServer := &http.Server{
		Handler:     handler,
		IdleTimeout: 5 * time.Second,
	}

	httpServer.SetKeepAlivesEnabled(true)

	listener, err := s.NewListener()
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	if err := httpServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) NewListener() (net.Listener, error) {
	addr := net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
	return net.Listen("tcp", addr)
}
