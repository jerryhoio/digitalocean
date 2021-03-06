package main

import (
	"context"
	"log"
	"os"

	greeterProto "github.com/jerryhoio/digitalocean/pkg/proto/greeter"
	greeterSvc "github.com/jerryhoio/digitalocean/pkg/service/greeter"

	digitaloceanProto "github.com/jerryhoio/digitalocean/pkg/proto/digitalocean"
	digitaloceanSvc "github.com/jerryhoio/digitalocean/pkg/service/digitalocean"

	do "github.com/jerryhoio/digitalocean/pkg/digitalocean"

	"github.com/micro/go-micro/server"

	"github.com/digitalocean/godo"
	"github.com/micro/go-micro"
	"golang.org/x/oauth2"
)

// implements the server.HandlerWrapper
func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("Server request: %s", req.Method())
		//TODO: Make it log more details (maybe client IP, provided opts etc.)

		return fn(ctx, req, rsp)
	}
}

func main() {
	log.Println("Checking environment...")
	pat := os.Getenv("DIGITALOCEAN_API_TOKEN")
	if pat == "" {
		log.Fatal("Missing environment variable: DIGITALOCEAN_API_TOKEN")
	}
	log.Println("Environment is OK")
	log.Println("Connecting to DigitalOcean...")
	token := &do.AccessToken{
		AccessToken: pat,
	}
	oauthClient := oauth2.NewClient(context.Background(), token)
	client := godo.NewClient(oauthClient)
	log.Println("Connected.")

	service := micro.NewService(
		micro.Name("digitalocean"),
		micro.WrapHandler(logWrapper),
	)

	service.Init()

	greeterProto.RegisterGreeterHandler(service.Server(), new(greeterSvc.Greeter))
	digitaloceanProto.RegisterDropletsHandler(service.Server(), &digitaloceanSvc.Droplets{Client: client})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
