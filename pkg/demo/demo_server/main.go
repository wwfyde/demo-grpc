package main

import (
	"context"
	"fmt"
	pb "github.com/wwfyde/demo-grpc/examples/demo/demo"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	pb.UnimplementedDemonstratorServer
}

func (s *server) DoSomething(ctx context.Context, in *pb.DoRequest) (*pb.DoReply, error) {
	return &pb.DoReply{Name: in.GetName()}, nil
}
func main() {
	// Listen port with tcp on 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterDemonstratorServer(s, &server{})
	_ = s.Serve(lis)

}
