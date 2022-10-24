package client

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"sync"

	"grandChallenge1/grpcchat123/conf"
	proto "grandChallenge1/grpcchat123/internal/chatserver/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	config = "../../../config.yaml"
)

//Client struct is the struct for the connected chat client
type Client struct {
	Wg          *sync.WaitGroup
	Client      proto.BroadcastClient
	Name        string
	RoomName    string
	BlockedName string
	Timestamp   string
	conf        *conf.Conf
}

func (c *Client) Connect(user *proto.User, room *proto.Room) error {
	var streamerror error

	stream, err := c.Client.CreateStream(context.Background(), &proto.Connect{
		User:    user,
		Room:    room,
		Active:  true,
		Blocked: c.BlockedName,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	c.Wg.Add(1)
	go func(str proto.Broadcast_CreateStreamClient) {
		defer c.Wg.Done()

		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("Error reading message: %v", err)
				break
			}

			fmt.Printf("%s : %s\n", msg.Name, msg.Content)

		}
	}(stream)

	return streamerror
}
func (c *Client) initialize() {
	c.Timestamp = time.Now().String()
	c.Wg = &sync.WaitGroup{}
	c.conf = &conf.Conf{}
	c.conf = c.conf.GetConf(config)
}

//Exec method is used to be called by the main function to initiate chat client
func (c *Client) Exec() {

	c.initialize()
	done := make(chan int)
	user_id_byte := sha256.Sum256([]byte(c.Timestamp + c.Name))
	room_id_byte := sha256.Sum256([]byte(c.RoomName))

	address := fmt.Sprintf("%s:%d", c.conf.ServerIP, c.conf.Serverport)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect to service: %v", err)
	}
	c.Client = proto.NewBroadcastClient(conn)
	user := &proto.User{
		Id:   hex.EncodeToString(user_id_byte[:]),
		Name: c.Name,
	}
	room := &proto.Room{
		Id:   hex.EncodeToString(room_id_byte[:]),
		Name: c.RoomName,
	}
	c.Connect(user, room)
	c.Wg.Add(1)
	go c.ReadMessage(user, room)
	go func() {
		c.Wg.Wait()
		close(done)
	}()
	<-done
}

//ReadMessage method reads message from the clients but not redpipe 
func (c *Client) ReadMessage(user *proto.User, room *proto.Room) {
	defer c.Wg.Done()
	if user.Name != "redpipe"{

	
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := &proto.Message{
			Id:        user.Id,
			Name:      user.Name,
			Content:   scanner.Text(),
			Timestamp: time.Now().String(),
			Room:      room.Id,
		}

		_, err := c.Client.BroadcastMessage(context.Background(), msg)
		if err != nil {
			fmt.Printf("Error Sending Message: %v", err)
			break
		}
	}
}
}

//ExecApi method is used to be called by the API to send messages
func (c *Client) ExecApi(msgTxt string) (err error) {

	c.initialize()
	user_id_byte := sha256.Sum256([]byte(c.Timestamp + c.Name))
	room_id_byte := sha256.Sum256([]byte(c.RoomName))

	address := fmt.Sprintf("%s:%d", c.conf.ServerIP, c.conf.Serverport)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c.Client = proto.NewBroadcastClient(conn)
	user := &proto.User{
		Id:   hex.EncodeToString(user_id_byte[:]),
		Name: c.Name,
	}
	room := &proto.Room{
		Id:   hex.EncodeToString(room_id_byte[:]),
		Name: c.RoomName,
	}
	msg := &proto.Message{
		Id:        user.Id,
		Name:      user.Name,
		Content:   msgTxt,
		Timestamp: time.Now().String(),
		Room:      room.Id,
	}

	_, err = c.Client.BroadcastMessage(context.Background(), msg)
	return
}
