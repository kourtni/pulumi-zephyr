package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket)
		_, err := createInfrastructure(ctx)
		if err != nil {
			return err
		}

		return nil
	})
}

// Simple struct to hold the infrastructure resources
type infrastructure struct {
	bucket *s3.BucketV2
}

// Testable function for creating infrastructure resources
func createInfrastructure(ctx *pulumi.Context) (*infrastructure, error) {
	// Create an AWS resource (S3 Bucket)
	bucket, err := s3.NewBucketV2(ctx, "my-bucket", nil)
	if err != nil {
		return nil, err
	}

	// Export the name of the bucket
	ctx.Export("bucketName", bucket.ID())

	return &infrastructure{
		bucket: bucket,
	}, nil
}