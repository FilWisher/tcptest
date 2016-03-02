package tcptest

import (
  "net"
)

type Server struct {
  Port string
  Handler func(net.Conn)
  Listener net.Listener
}

func NewServer(port string, h func(net.Conn)) *Server {
  s := &Server{Port: port, Handler:h}
  s.Start(port)
  return s
}

func (s *Server) goStart() {
  go func() {
    for {
      conn, err := s.Listener.Accept()
      if err != nil {
        panic("Test server failed")
      }
      s.Handler(conn)
    }
  }()
}

func (s *Server) Start(port string) {
  s.Port = port
  l, err := net.Listen("tcp", port)
  if err != nil {
    panic("Test server failed")
  }
  s.Listener = l
  s.goStart()
}
