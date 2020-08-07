package main

import (
	"log"
	"net"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-gRPC-protobuff-CRUD/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type server struct{}

type mong struct {
	Operation *mgo.Collection
}

var DB *mong

func main() {
	mongo, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatalf("Could not connect to the MongoDB server: %v", err)
	}
	defer mongo.Close()

	DB = &mong{mongo.DB("inventory_db").C("inventory_product")}

	// Host grpc service
	listen, err := net.Listen("tcp", "127.0.0.1:50052")
	if err != nil {
		log.Fatalf("Could not listen on port: %v", err)
	}

	// gRPC server
	s := grpc.NewServer()
	pb.RegisterCRUDServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Printf("Hosting server on: %s", listen.Addr().String())
}

func (s *server) CreateProduct(ctx context.Context, product *pb.InventoryProduct) (*pb.ID, error) {
	if product.Barcode == "" {
		return nil, status.Error(codes.InvalidArgument, "barcode can't be empty")
	}

	return &pb.ID{Id: product.Barcode}, DB.Operation.Insert(product)
}


func (s *server) ReadProduct(ctx context.Context, req *pb.ID) (*pb.InventoryProduct, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "barcode can't be empty")
	}

	var result pb.InventoryProduct
	err := DB.Operation.Find(bson.M{"barcode": req.Id}).One(&result)
	if err != nil {
		log.Printf("Error retrieving product with id: %s, error: %v", req.Id, err)
		return nil, err
	}

	return &result, nil
}

func (s *server) UpdateProduct(ctx context.Context, product *pb.InventoryProduct) (*pb.ID, error) {
	if product.Barcode == "" {
		return nil, status.Error(codes.InvalidArgument, "barcode can't be empty")
	}

	find := bson.M{"barcode": product.Barcode}
	update := bson.M{"$set": bson.M{"store": product.Store, "name": product.Name, "category": product.Category, "description": product.Description}}

	return &pb.ID{Id: product.Barcode}, DB.Operation.Update(find, update)
}


func (s *server) DeleteProduct(ctx context.Context, req *pb.ID) (*pb.ID, error) {
	return &pb.ID{Id: req.Id}, DB.Operation.Remove(bson.M{"id": req.Id})
}
