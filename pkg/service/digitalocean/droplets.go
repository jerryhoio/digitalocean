package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"
	proto "github.com/jerryhoio/digitalocean/pkg/proto/digitalocean"
)

// Droplets service implements DigitalOcean Droplets APIs
type Droplets struct {
	Client *godo.Client
}

// List function returns List with provided string - just for fun
func (g *Droplets) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	rDroplets, _, _ := g.Client.Droplets.List(ctx, &godo.ListOptions{})

	var droplets []*proto.Droplet
	for _, d := range rDroplets {
		droplets = append(droplets, &proto.Droplet{
			Id:      int64(d.ID),
			Name:    d.Name,
			Memory:  int32(d.Memory),
			Vcpus:   int32(d.Vcpus),
			Disk:    int32(d.Disk),
			Created: d.Created,
		})
	}

	//TODO: Add pagination, handle pagination

	rsp.Droplets = droplets
	return nil
}
