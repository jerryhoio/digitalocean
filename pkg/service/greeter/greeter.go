package greeter

import (
	"context"

	proto "github.com/jerryhoio/digitalocean/pkg/proto/greeter"
)

// Greeter test service
type Greeter struct{}

// Hello function returns Hello with provided string - just for fun
func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}
