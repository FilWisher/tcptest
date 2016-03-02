package tcptest

import (
  "testing"
  "net"
)

func TestStartServer(t *testing.T) {
  buffer := make([]byte, 1024)
  NewServer(":8080", func (conn net.Conn) {
    n, err := conn.Read(buffer)
    if err != nil {
      t.Fatalf("Test server failed: %s", err.Error())
    }
    conn.Write(buffer[:n])
  })
}

func TestConnectionAcrossTests(t *testing.T) {

  msg := "Attempting normal tcp connection"
  conn, err := net.Dial("tcp", ":8080")
  if err != nil {
    t.Fatalf("Test server failed: %s", err.Error())
  }

  received := make([]byte, 1024)
  conn.Write([]byte(msg))
  n, err := conn.Read(received)
  if err != nil {
    t.Fatalf("Test server failed: %s", err.Error())
  }
  if msg != string(received[:n]) {
    t.Errorf("Wanted: %s, got: %s", err.Error())
  }
}
