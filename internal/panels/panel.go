package panels

import (
	"context"
	"golang.org/x/sync/errgroup"
)

type Panel struct {
	adapters []PanelAdapter
}

func InitPanel(adapters []PanelAdapter) Panel {
	return Panel{adapters}
}

func (p *Panel) SetVpnStatus(ctx context.Context, status bool) error {
	g, gCtx := errgroup.WithContext(ctx)
	for _, a := range p.adapters {
		g.Go(func() error {
			return a.SetStatusVPN(gCtx, status)
		})
	}

	return g.Wait()
}

func (p *Panel) SetInformation(ctx context.Context, info string) error {
	g, gCtx := errgroup.WithContext(ctx)
	for _, a := range p.adapters {
		g.Go(func() error {
			return a.SetInformationVPN(gCtx, info)
		})
	}
	return g.Wait()
}
