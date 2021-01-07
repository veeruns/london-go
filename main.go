package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	plugin "github.com/hashicorp/go-plugin"

	"github.com/veeruns/rpcserver/london-go/greeting"
)

func main() {
	language := "english"
	argu := "Test"
	if len(os.Args) > 1 {
		language = os.Args[1]
		argu = strings.Join(os.Args[2:], " ")
	}

	// Note this is BAD, but demonstrates the concept.
	pluginCmd := fmt.Sprintf("./greeter-%s", language)

	log.SetOutput(ioutil.Discard)

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  1,
			MagicCookieKey:   "LONDON_GO_DEMO",
			MagicCookieValue: "Hi",
		},
		Plugins: map[string]plugin.Plugin{
			"greeter": new(greeting.GreeterPlugin),
		},
		Cmd: exec.Command(pluginCmd),
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		log.Fatal(err)
	}

	// We should have a Greeter now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	greeter := raw.(greeting.Greeter)
	test := greeter.Greet(argu)
	fmt.Printf("%s\n", test)
}
