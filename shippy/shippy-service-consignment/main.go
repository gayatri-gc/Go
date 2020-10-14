package main

import(
	"context"
	"fmt"
	"net"
	"sync"

	pb "shippy-service-consignment/proto/consignment"
//	pb "github.com/GayatriChougale/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type repository interface{
	Create(*pb.Consignment) (*pb.Consignment, error)
} 

type Repository struct{
	mu sync.RWMutex
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment)(*pb.Consignment, error){
	repo.mu.Lock()
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

type service struct{
	repo repository
}

func (s *service) CreateConsignment(ctx context.Context, consignment *pb.Consignment)(*pb.Response, error){
	consignment, err := s.repo.Create(consignment)
	if err != nil{
		return nil, err
	}
	return &pb.Response{Created: true , Consignment: consignment}, nil
}

func main(){
	repo := &Repository{}
	lis, err := net.Listen("tcp", ":5001")
	if err != nil{
		fmt.Println("Failed to create connection", err)
	}

	s := grpc.NewServer()
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)

	fmt.Println("Running on port:", 5001)
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: %v", err)
	}





}