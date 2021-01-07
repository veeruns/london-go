package greeting

import "net/rpc"

type GreeterRPC struct{ client *rpc.Client }

func (g *GreeterRPC) Greet(input string) string {
	var resp string
	err := g.client.Call("Plugin.Greet", input, &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}
