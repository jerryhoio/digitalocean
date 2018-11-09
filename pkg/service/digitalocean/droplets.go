package digitalocean

import (
	"context"
	"log"
	"strconv"

	"github.com/digitalocean/godo"
	pkgdo "github.com/jerryhoio/digitalocean/pkg/digitalocean"
	proto "github.com/jerryhoio/digitalocean/pkg/proto/digitalocean"
)

// Droplets service implements DigitalOcean Droplets APIs
type Droplets struct {
	Client *godo.Client
}

// Create Droplet with provided parameters
func (g *Droplets) Create(ctx context.Context, req *proto.CreateRequest, rsp *proto.CreateResponse) error {
	createRequest := &godo.DropletCreateRequest{
		Name:   req.Name,
		Region: req.Region,
		Size:   req.Size,
		Image: godo.DropletCreateImage{
			Slug: req.Image,
		},
	}

	droplet, _, err := g.Client.Droplets.Create(ctx, createRequest)
	if err != nil {
		log.Printf("[ERROR] Droplets.Create failed: %s\n\n", err)
		return err
	}

	rsp.Message = "Droplet created: " + strconv.Itoa(droplet.ID)

	return nil
}

// Get function returns droplet
func (g *Droplets) Get(ctx context.Context, req *proto.GetRequest, rsp *proto.GetResponse) error {
	droplet, _, err := g.Client.Droplets.Get(ctx, int(req.Id))
	if err != nil {
		log.Printf("[ERROR] Droplets.Get failed: %s\n\n", err)
		return err
	}

	mapper := pkgdo.Mapper{}
	rsp.Droplet = mapper.DropletToProto(droplet)

	return nil
}

// List function returns List with provided string - just for fun
func (g *Droplets) List(ctx context.Context, req *proto.ListRequest, rsp *proto.ListResponse) error {
	mapper := pkgdo.Mapper{}
	var list []*proto.Droplet
	rsp.Droplets = list

	opt := &godo.ListOptions{}
	for {
		rDroplets, resp, err := g.Client.Droplets.List(ctx, opt)
		if err != nil {
			return err
		}

		for _, d := range rDroplets {
			nd := mapper.DropletToProto(&d)

			list = append(list, nd)
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

// Delete function deletes droplet
func (g *Droplets) Delete(ctx context.Context, req *proto.DeleteRequest, rsp *proto.DeleteResponse) error {
	_, err := g.Client.Droplets.Delete(ctx, int(req.Id))
	if err != nil {
		log.Printf("[ERROR] Droplets.Delete failed: %s\n\n", err)
		return err
	}

	rsp.Message = "Droplet deleted"

	return nil
}
