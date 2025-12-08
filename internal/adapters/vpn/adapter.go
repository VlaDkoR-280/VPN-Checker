package vpn

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"strings"
)

type Adapter struct {
	ns string
}

func InitAdapter(ctx context.Context, prefixName string) (*Adapter, error) {
	a := &Adapter{
		ns: fmt.Sprintf("%s-vpn-checker", prefixName),
	}

	out, err := exec.CommandContext(ctx, "sudo", "ip", "netns", "add", a.ns).CombinedOutput()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link add veth1 type veth peer name veth2").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link add veth3 type veth peer name veth4").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth2 netns", a.ns).Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth4 netns", a.ns).Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link add name br-vpn type bridge").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "addr add 192.168.89.1/24 dev br-vpn").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set br-vpn up").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth1 master br-vpn").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth3 master br-vpn").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth1 up").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "link set veth3 up").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "echo", "echo 1 > /proc/sys/net/ipv4/ip_forward").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "iptables", "-t nat -A POSTROUTING -o eth0 -j MASQUERADE").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "iptables", "-A FORWARD -i br-vpn -o eth0 -j ACCEPT").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "iptables", "-A FORWARD -i eth0 -o br-vpn -j ACCEPT").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "netns exec", a.ns, "ip link set veth2 up").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "netns exec", a.ns, "ip link set veth4 up").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	out, err = exec.CommandContext(ctx, "ip", "netns exec", a.ns, "ip addr add 192.168.89.2/24 dev veth2").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}
	if out != nil {
		return nil, errors.Wrap(err, string(out))
	}

	out, err = exec.CommandContext(ctx, "ip", "netns exec", a.ns, "ip route add default via 192.168.89.1").Output()
	if err != nil {
		if out != nil {
			return nil, errors.Wrap(err, string(out))
		}
		return nil, errors.WithStack(err)
	}

	return a, nil
}

func (a *Adapter) TestBaseConnection(ctx context.Context) error {
	_, err := exec.CommandContext(ctx, "ip", "link set veth1 down").Output()
	return errors.WithStack(err)

}

func (a *Adapter) GetIPAddress(ctx context.Context) (string, error) {
	out, err := exec.CommandContext(ctx, "ip", "netns exec", a.ns, "").Output()
	if err != nil {
		return "", errors.WithStack(err)
	}
	return strings.TrimSpace(string(out)), nil
}
