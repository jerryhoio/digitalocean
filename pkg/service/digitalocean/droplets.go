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
			nd := &proto.Droplet{
				Id:      int64(d.ID),
				Name:    d.Name,
				Memory:  int32(d.Memory),
				Vcpus:   int32(d.Vcpus),
				Disk:    int32(d.Disk),
				Created: d.Created,
				Region: &proto.Region{
					Name:      d.Region.Name,
					Slug:      d.Region.Slug,
					Available: d.Region.Available,
					Sizes:     d.Region.Sizes,
					Features:  d.Region.Features,
				},
				Image: &proto.Image{
					Id:           int64(d.Image.ID),
					Name:         d.Image.Name,
					Type:         d.Image.Type,
					Distribution: d.Image.Distribution,
					Slug:         d.Image.Slug,
					Public:       d.Image.Public,
					Regions:      d.Image.Regions,
					MinDiskSize:  int32(d.Image.MinDiskSize),
					Created:      d.Image.Created,
				},
				Size: &proto.Size{
					Slug:         d.Size.Slug,
					Memory:       int32(d.Size.Memory),
					Vcpus:        int32(d.Size.Vcpus),
					Disk:         int32(d.Size.Disk),
					PriceMonthly: d.Size.PriceMonthly,
					PriceHourly:  d.Size.PriceHourly,
					Regions:      d.Size.Regions,
					Available:    d.Size.Available,
					Transfer:     d.Size.Transfer,
				},
				Networks: &proto.Networks{},
			}

			if d.Kernel != nil {
				nd.Kernel = &proto.Kernel{
					Id:      int64(d.Kernel.ID),
					Name:    d.Kernel.Name,
					Version: d.Kernel.Version,
				}
			}

			if d.NextBackupWindow != nil {
				nd.NextBackupWindow = &proto.BackupWindow{
					Start: d.NextBackupWindow.Start.String(),
					End:   d.NextBackupWindow.End.String(),
				}
			}

			if len(d.Networks.V4) > 0 {
				for _, x := range d.Networks.V4 {
					nd.Networks.V4 = append(nd.Networks.V4, &proto.NetworkV4{
						IPAddress: x.IPAddress,
						Netmask:   x.Netmask,
						Gateway:   x.Gateway,
						Type:      x.Type,
					})
				}
			}

			if len(d.Networks.V6) > 0 {
				for _, x := range d.Networks.V6 {
					nd.Networks.V6 = append(nd.Networks.V6, &proto.NetworkV6{
						IPAddress: x.IPAddress,
						Netmask:   int32(x.Netmask),
						Gateway:   x.Gateway,
						Type:      x.Type,
					})
				}
			}

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
