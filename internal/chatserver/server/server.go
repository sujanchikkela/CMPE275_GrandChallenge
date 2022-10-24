package server

import (
	"os"
	"sync"

	proto "grandChallenge1/grpcchat123/internal/chatserver/proto"
	context "golang.org/x/net/context"
	glog "google.golang.org/grpc/grpclog"
)

const (
	config = "../../../config.yaml"
)

//Connection struct holds user,room and stream to the GRPC server
type Connection struct {
	stream  proto.Broadcast_CreateStreamServer
	id      string
	room_id string
	blocked string
	active  bool
	error   chan error
}

//Server struct holds all the active connections, file descriptor and a logger to be used by GRPC server
type Server struct {
	Connection []*Connection
	FileDesc   *os.File
	GrpcLog    glog.LoggerV2
}

//CreateStream method will be used by clients to create a stream with the GRPC Server via RPC
func (s *Server) CreateStream(pconn *proto.Connect, stream proto.Broadcast_CreateStreamServer) error {
	conn := &Connection{
		stream:  stream,
		id:      pconn.User.Id,
		active:  true,
		room_id: pconn.Room.Id,
		blocked: pconn.Blocked,
		error:   make(chan error),
	}

	s.Connection = append(s.Connection, conn)

	return <-conn.error
}

//BroadcastMessage method will be used by clients to send messages  via RPC
func (s *Server) BroadcastMessage(ctx context.Context, msg *proto.Message) (*proto.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)

	//Append messages to the chat log
	if _, err := s.FileDesc.Write([]byte(msg.Name + ":" + msg.Timestamp + ":" + msg.Content + "\n")); err != nil {
		s.GrpcLog.Fatal(err)
	}
	for _, conn := range s.Connection {

		//Skip clients if they are not part of the message's chat room
		if msg.Room != conn.room_id {
			continue
		}
		//Skip messages from blocked users
		if conn.blocked == msg.Name {
			continue
		}
		wait.Add(1)

		go func(msg *proto.Message, conn *Connection) {
			defer wait.Done()
			if conn.active {
				err := conn.stream.Send(msg)
				s.GrpcLog.Info("Sending message to: ", conn.stream)

				if err != nil {
					s.GrpcLog.Errorf("Error with Stream: %v - Error: %v", conn.stream, err)
					conn.active = false
					conn.error <- err
				}
			}
		}(msg, conn)

	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &proto.Close{}, nil
}
