package socketproxy

import (
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

type SocketProxyController struct {
	runAsServer bool
	socket      string
	listener    net.Listener
}

func NewSocketProxyController(socket string, runAsServer bool) (*SocketProxyController, error) {

	var listener net.Listener
	if !runAsServer {
		// Remove socket file if it exists
		_, err := os.Stat(socket)
		if err == nil {
			os.Remove(socket)
		}

		// Listen on unix socket
		listener, err = net.Listen("unix", socket)
		if err != nil {
			log.Errorf("Unable to listen on unix socket: %v", err)
			return nil, err
		}
	}

	// Change socket owner and permissions to allow any users in group 'users' (gid: 100)
	err := os.Chown(socket, os.Getuid(), 100)
	if err != nil {
		log.Errorf("Unable to change socket file ownership: %v", err)
	}
	err = os.Chmod(socket, 0770)
	if err != nil {
		log.Errorf("Unable to change socket file permissions: %v", err)
	}

	return &SocketProxyController{socket: socket, runAsServer: runAsServer, listener: listener}, nil
}

func (s *SocketProxyController) Dial() (net.Conn, error) {

	if !s.runAsServer {
		return nil, fmt.Errorf("socket proxy runs as client")
	}

	// Dial to the unix socket
	conn, err := net.Dial("unix", s.socket)
	if err != nil {
		log.Errorf("unable to dial unix socket: %v", err)
		return nil, err
	}
	return conn, nil
}

func (s *SocketProxyController) Accept() (net.Conn, error) {

	if s.runAsServer {
		return nil, fmt.Errorf("socket proxy runs as server")
	}
	// Accept new connection
	conn, err := s.listener.Accept()
	if err != nil {
		log.Errorf("unable to accept socket connection: %v", err)
		return nil, err
	}
	return conn, nil
}

func (s *SocketProxyController) Close() error {
	if s.listener != nil {
		err := s.listener.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *SocketProxyController) Write(conn net.Conn, data []byte) error {
	_, err := conn.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func (s *SocketProxyController) Read(conn net.Conn) ([]byte, error) {

	var data []byte
	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	data = append(data, buf[:n]...)

	return data, nil
}
