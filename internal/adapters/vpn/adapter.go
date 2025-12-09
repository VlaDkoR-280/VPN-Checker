package vpn

import (
	"context"
	"github.com/pkg/errors"
	"os/exec"
)

type Adapter struct {
	ns string
}

func InitAdapter(nsName string) *Adapter {
	return &Adapter{ns: nsName}
}

func (a *Adapter) CheckStatus(ctx context.Context) error {
	out, err := exec.CommandContext(ctx, "ip", "netns", "exec", a.ns, "curl", "ifconfig.me").CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}
	return nil
}
