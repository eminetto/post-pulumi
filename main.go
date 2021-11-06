package main

import (
	"github.com/eminetto/post-pulumi/iac"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return iac.FargateRun(ctx)
	})
}
