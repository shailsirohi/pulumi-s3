package main

import (
"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/s3"
"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx,"")
		name := conf.Require("name")

		// Create an AWS resource (S3 Bucket)
		bucket, err := s3.NewBucket(ctx, name, nil)
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}