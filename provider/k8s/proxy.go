package k8s

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/convox/rack/pkg/helpers"
	"github.com/convox/rack/pkg/structs"
	"github.com/convox/stdsdk"
)

func (p *Provider) Proxy(host string, port int, rw io.ReadWriter, opts structs.ProxyOptions) error {
	cn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 3*time.Second)
	if err != nil {
		return err
	}

	if helpers.DefaultBool(opts.TLS, false) {
		cn = tls.Client(cn, &tls.Config{})
	}

	if err := stdsdk.CopyStreamToEachOther(cn, rw); err != nil {
		return err
	}

	return nil
}
