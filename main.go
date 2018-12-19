package main

import (
	"github.com/hashicorp/sentinel-sdk"
	"github.com/hashicorp/sentinel-sdk/framework"
	"github.com/hashicorp/sentinel-sdk/rpc"
	"github.com/scottwinkler/sentinel-example-import/example"
)

func main() {
	rpc.Serve(&rpc.ServeOpts{
		ImportFunc: func() sdk.Import {
			return &framework.Import{Root: example.New()}
		},
	})
}
