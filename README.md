# Performs CRUD (Create, Read, Update and Delete) operations of inventory product:

## Dependencies
Make sure you have [GOPATH](https://github.com/golang/go/wiki/GOPATH)
environment variable set properly.  
1. Install [Golang](https://golang.org/doc/install) and run:  
      `go get -u github.com/golang/protobuf/{proto,protoc-gen-go}`  
      `go get -u google.golang.org/grpc`  
      `go get gopkg.in/mgo.v2`  
      `go get golang.org/x/net/context`  
      
2. Install [Docker](https://www.docker.com) and run the following:
`docker run -p 27017:27017 -d mongo`

## Cloning the repository
`cd $HOME/src/github.com/`  
`git clone https://github.com/devendradora-ai/go-gRPC-protobuff-CRUD.git`

## Running the app
Make sure MongoDB server is up and running on `127.0.0.1:27017`.  

Open two terminals and run: `cd $HOME/src/github.com/go-gRPC-protobuff-CRUD` on both of them.    
First terminal:  
      `go run server/server.go`  
      Hit allow on the popup.    
Second terminal:  
      `go run client/main.go`


## TODO
1. Have n't got enough time to explore more so have n't followed best practices or have done enough validations for various operations    


