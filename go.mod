module github.com/veeruns/rpcserver/london-go

go 1.15

require (
	github.com/hashicorp/go-plugin v1.4.0
	github.com/jen20/london-go v0.0.0-20170117122103-3f9d965246d5
)

replace "github.com/veeruns/rpcserver/london-go/greeting" => "../greeting"