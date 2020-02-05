/*
 *
 * Copyright 2019 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Binary server is an example server.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/olongfen/note/examples/features/keepalive/pb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	// pb "google.golang.org/grpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50052, "port number")

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

// server implements EchoServer.
type server struct {
	pb.UnimplementedHallServer
}

//func (s *server) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
//	return &pb.EchoResponse{Message: req.Message}, nil
//}
func (s *server) NewRoom(ctx context.Context, in *pb.RoomAddRequest) (ret *pb.RoomResponse, err error) {
	return
}

func (s *server) GetRoom(ctx context.Context, in *pb.RoomGetRequest) (ret *pb.RoomResponse, err error) {
	return
}

func (s *server) ExitRoom(ctx context.Context, in *pb.RoomExistRequest) (ret *pb.RoomResponse, err error) {
	return
}

func (s *server) UpdateRoom(ctx context.Context, in *pb.RoomUpdateRequest) (ret *pb.RoomResponse, err error) {
	return
}
func (s *server) GetUser(ctx context.Context, in *pb.UserGetRequest) (ret *pb.UserResponse, err error) {

	return
}

func (s *server) GetWalletByWid(ctx context.Context, in *pb.WalletGetByWidRequest) (ret *pb.WalletResponse, err error) {

	return
}

func (s *server) GetWalletSelfList(ctx context.Context, in *pb.WalletGetByUIDRequest) (ret *pb.WalletListResponse, err error) {

	return
}

func (s *server) CheckToken(ctx context.Context, request *pb.CheckTokenRequest) (ret *pb.CheckTokenResponse, err error) {
	ret = &pb.CheckTokenResponse{
		Error: "sdsadasd",
	}
	return
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	pb.RegisterHallServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
