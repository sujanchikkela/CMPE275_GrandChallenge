package main

import( "fmt" "net" "os" "grandChallenge1/grpcchat123/conf" proto "grandChallenge1/grpcchat123/internal/chatserver/proto"
	server "grandChallenge1/grpcchat123/internal/chatserver/server" "google.golang.org/grpc" glog "google.golang.org/grpc/grpclog"
       )
const (	config = "../../../config.yaml" )
var grpcLog glog.LoggerV2
var fileDesc *os.File
var err error
var cnf *conf.Conf
func init() 
{
	cnf = &conf.Conf{} cnf = cnf.GetConf(config)
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
	fileDesc, err = os.OpenFile(cnf.ChatLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil { grpcLog.Fatal(err) }
	if err := os.Truncate(cnf.ChatLog, 0); err != nil { grpcLog.Errorf("Failed to truncate: %v", err) }
}
func main() 
{
	var connection []*server.Connection server := &server.Server{ Connection: connection, FileDesc:   fileDesc, GrpcLog:    grpcLog, }
	grpcServer := grpc.NewServer()
	tcpPort := fmt.Sprintf(":%d", cnf.Serverport)
	listener, err := net.Listen("tcp", tcpPort)
	if err != nil { grpcLog.Fatalf("error creating the server %v", err) }
	grpcLog.Infof("Starting server at port :%d", cnf.Serverport)
	proto.RegisterBroadcastServer(grpcServer, server)
	grpcServer.Serve(listener)
	if err := fileDesc.Close(); err != nil { grpcLog.Fatal(err) }
	err = os.Remove(cnf.ChatLog)
	if err != nil {	grpcLog.Fatal(err) }
}
