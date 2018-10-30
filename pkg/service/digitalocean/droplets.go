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
	var list []*proto.Droplet
	rsp.Droplets = list

	opt := &godo.ListOptions{}
	for {
		rDroplets, resp, err := g.Client.Droplets.List(ctx, opt)
		if err != nil {
			return err
		}

		for _, d := range rDroplets {
			list = append(list, &proto.Droplet{
				Id:      int64(d.ID),
				Name:    d.Name,
				Memory:  int32(d.Memory),
				Vcpus:   int32(d.Vcpus),
				Disk:    int32(d.Disk),
				Created: d.Created,
			})
		}

		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return err
		}

		// set the page we want for the next request
		opt.Page = page + 1
	}

	rsp.Droplets = list

	return nil
}
