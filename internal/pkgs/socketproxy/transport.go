package socketproxy

import (
	"bytes"
	"fmt"
	socket_api "givc/api/socket"
	"io"
	"net"
	"time"

	givc_grpc "givc/internal/pkgs/grpc"
	givc_types "givc/internal/pkgs/types"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"
)

type SocketProxyServer struct {
	socketProxy     *SocketProxyController // Socket proxy controller
	conn            net.Conn               // Socket connection
	streamConnected bool                   // Stream connection status; used to sync grpc stream
	socketConnected bool                   // Socket connection status; used to indicate that socket is connected
	clientConnected bool                   // Client connection status; used to indicate that grpc client is connected
	socket_api.UnimplementedSocketStreamServer
}

func (s *SocketProxyServer) Name() string {
	return "Socket Proxy Server"
}

func (s *SocketProxyServer) RegisterGrpcService(srv *grpc.Server) {
	socket_api.RegisterSocketStreamServer(srv, s)
}

/*
	SocketProxyServer is a GRPC service that proxies data between a local socket and a remote GRPC server.

	It can run in two modes, server or client:

		1. As a server: The server waits for a remote client to connect.
		Once connected, it dials to a local socket and proxies data between the remote client and the local socket.

		2. As a client: The client creates a socket listener and waits for a connection to the local socket.
		Once a client (application) connects to the socket, it initiates the connection to the server.

	Both client and server run a GRPC server, and the main routine 'StreamToRemote' which must be run in a different go routine
	than the GRPC server itself.

	When the client is connected, the server sends a 'PROXY_CLIENT_CONNECTED' message to the remote server. The server awaits
	this message and then connects to the local socket via 'dial()'. The client repeats the request until the server signals
	that the socket is connected and the GRPC stream is initiated.

	As the socket read() function is blocking, the end of a socket connection must be forwarded to the remote, and locally
	closed (here, by closing the socket connection) on both ends. This is done by sending an EOF message to the remote counterpart.

	Both ends unwind and close both socket and GRPC connections when any socket read error occurs, EOF or otherwise.
*/

func NewSocketProxyServer(socket string, runAsServer bool) (*SocketProxyServer, error) {

	// Create a new socket proxy controller
	var err error
	socketProxy, err := NewSocketProxyController(socket, runAsServer)
	if err != nil {
		return nil, err
	}

	return &SocketProxyServer{
		socketProxy:     socketProxy,
		conn:            nil,
		streamConnected: false,
		socketConnected: false,
		clientConnected: false,
	}, nil
}

func (s *SocketProxyServer) Close() error {
	return s.socketProxy.Close()
}

func (s *SocketProxyServer) StreamToRemote(ctx context.Context, cfg *givc_types.EndpointConfig) error {

	defer s.socketProxy.Close()

	// Setup and dial GRPC client
	grpcClientConn, err := givc_grpc.NewClient(cfg, true)
	if err != nil {
		return err
	}
	defer grpcClientConn.Close()

	// Create streaming client
	socketStreamClient := socket_api.NewSocketStreamClient(grpcClientConn)
	if socketStreamClient == nil {
		return fmt.Errorf("failed to create 'NewSocketStreamClient'")
	}

	for {
		select {
		// Return if context is done
		case <-ctx.Done():
			return nil

		// Stream data to remote
		default:

			// Wait for local socket connection
			if !s.socketProxy.runAsServer {

				// Wait for new connection to socket; blocking
				s.conn, err = s.socketProxy.Accept()
				if err != nil {
					return err
				}
				s.socketConnected = true

				// Tell remote server that client is connected and wait for server connection
				status, err := socketStreamClient.Sync(ctx, &socket_api.SyncData{Status: givc_types.PROXY_CLIENT_CONNECTED})
				for err != nil || status.GetStatus() != givc_types.PROXY_SERVER_CONNECTED {
					time.Sleep(1 * time.Second)
					status, err = socketStreamClient.Sync(ctx, &socket_api.SyncData{Status: givc_types.PROXY_CLIENT_CONNECTED})
				}
			}

			// Wait for client connection
			if s.socketProxy.runAsServer {

				// Wait for remote client to make connection; blocking
				for !s.clientConnected {
					time.Sleep(500 * time.Millisecond)
				}

				// Connect to socket
				s.conn, err = s.socketProxy.Dial()
				if err != nil {
					return err
				}
				s.socketConnected = true
			}

			// Create new GRPC stream
			stream, err := socketStreamClient.ReceiveData(ctx)
			if err != nil {
				log.Warnf("GRPC stream connection failed: %v", err)
				return err
			}
			s.streamConnected = true

			// Proxy loop for active connection
			for {
				// Read data from socket
				data, err := s.socketProxy.Read(s.conn)
				if err != nil {
					// Forward any read error to terminate stream and socket connections on both ends
					log.Infof(">> Socket read error: %v", err)
					message := &socket_api.SocketDataStream{
						Data: []byte(io.EOF.Error()),
					}
					err = stream.Send(message)
					if err != nil {
						log.Errorf("failed to send EOF to remote: %v", err)
					}
					break
				}

				// Send data to grpc stream
				message := &socket_api.SocketDataStream{
					Data: data,
				}
				err = stream.Send(message)
				if err != nil {
					break
				}
				log.Infof("Sent data: %s", data)
			}

			// Close socket connection
			if s.conn != nil {
				s.conn.Close()
			}
			s.socketConnected = false

			// Close stream connection
			if stream != nil {
				stream.CloseSend()
			}
			s.streamConnected = false

			// Reset client connection status
			s.clientConnected = false
		}
	}

}

func (s *SocketProxyServer) ReceiveData(stream socket_api.SocketStream_ReceiveDataServer) error {

	for {
		// Read data from grpc stream
		data, err := stream.Recv()
		if err != nil {
			log.Warnf(">> GRPC failure: %v", err)
			return err
		}
		log.Infof("Recv data: %s", data.GetData())

		// Check for EOF; and close socket connection
		if bytes.Equal(data.GetData(), []byte(io.EOF.Error())) {
			if s.streamConnected {
				if s.conn != nil {
					s.conn.Close()
				}
			}
			return fmt.Errorf("EOF received")
		}

		// Write data to socket
		err = s.socketProxy.Write(s.conn, data.GetData())
		if err != nil {
			log.Warnf("Error writing to socket: %v", err)
			return err
		}
	}
}

func (s *SocketProxyServer) Sync(ctx context.Context, status *socket_api.SyncData) (*socket_api.SyncData, error) {

	resp := givc_types.PROXY_NULL

	if s.socketProxy.runAsServer {

		if status.GetStatus() == givc_types.PROXY_CLIENT_CONNECTED {
			s.clientConnected = true
		}
		if s.socketConnected {
			resp = givc_types.PROXY_SERVER_CONNECTED
		}

	} else {
		if s.socketConnected {
			resp = givc_types.PROXY_CLIENT_CONNECTED
		}
	}

	return &socket_api.SyncData{Status: resp}, nil
}
