package main

import (
	"sync"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/stretchr/testify/assert"
)

type mocks int

// NewResource mocks the creation of any resource.
func (mocks) NewResource(args pulumi.MockResourceArgs) (string, resource.PropertyMap, error) {
	return args.Name + "_id", args.Inputs, nil
}

// Call mocks function calls.
func (mocks) Call(args pulumi.MockCallArgs) (resource.PropertyMap, error) {
	return args.Args, nil
}

func TestCreateInfrastructure(t *testing.T) {
	err := pulumi.RunErr(func(ctx *pulumi.Context) error {
		infra, err := createInfrastructure(ctx)
		assert.NoError(t, err)

		var wg sync.WaitGroup
		wg.Add(1)

		pulumi.All(
			infra.bucket.URN(),
			infra.bucket.Tags,
		).ApplyT(func(all []interface{}) error {
			urn := all[0].(pulumi.URN)
			tags := all[1].(map[string]string)

			expectedBucketUrn := "urn:pulumi:stack::zephyr::aws:s3/bucketV2:BucketV2::my-bucket"
			assert.Equalf(t, pulumi.URN(expectedBucketUrn), urn, "URN name mismatch %v", urn)
			assert.Emptyf(t, tags, "bucket should have no tags; found: %v", tags)
			wg.Done()
			return nil
		})

		wg.Wait()
		return nil
	}, pulumi.WithMocks("zephyr", "stack", mocks(0)))
	assert.NoError(t, err)
}