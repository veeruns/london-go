package greeting

import "fmt"

type GreeterRPCServer struct {
	// This is the real implementation
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args string, resp *string) error {
	str := fmt.Sprintf("%v ", args)
	*resp = s.Impl.Greet(str)
	return nil
}
