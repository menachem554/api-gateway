package bookstore

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	pb "github.com/menachem554/Bookstore/proto"
	"google.golang.org/grpc"
)

var (
	router = gin.Default()

	// C service Client
	C  pb.BookstoreClient
	cc *grpc.ClientConn
)

// StartClient : To start the client service
func StartClient() {
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
		MaxAge:           0,
	}))
	Router()
	connectServer()
	go func() {
		err := router.Run(":9091")
		if err != nil {
			log.Fatalf("error with the the server: %v\n", err)

		} else {
			fmt.Println("Server successfully started on port :9091")
		}

	}()

	// Create a channel to receive OS signals and detect CTRL+C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Closing the Connection with server")
	cc.Close()

}

// connect to grpc server
func connectServer() {
	opts := grpc.WithInsecure()
	var err error
	cc, err = grpc.Dial("bookstore:9090", opts)
	if err != nil {
		fmt.Println("Error while connection to the server", err.Error())
		panic(err)
	} else {
		fmt.Println("connect to grpc server localhost:9090")
	}

	C = pb.NewBookstoreClient(cc)
	fmt.Println("Connection to Server is successfull")
}
