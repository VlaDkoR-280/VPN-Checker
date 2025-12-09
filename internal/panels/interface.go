package panels

import "context"

type PanelAdapter interface {
	SetStatusVPN(ctx context.Context, status bool) error
	SetInformationVPN(ctx context.Context, info string) error
}
