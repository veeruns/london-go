package main

import (
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/veeruns/rpcserver/london-go/greeting"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "LONDON_GO_DEMO",
			MagicCookieValue: "Hi",
		},
		Plugins: map[string]plugin.Plugin{
			"greeter": greeting.GreeterPlugin{
				Impl: GreeterEnglish{},
			},
		},
	})
}

type GreeterEnglish struct{}

func (GreeterEnglish) Greet(a string) string {
	return fmt.Sprintf("Hi: %s", a)
}
