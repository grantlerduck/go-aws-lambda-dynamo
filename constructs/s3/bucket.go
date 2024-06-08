package s3

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewRemoveableBucket(scope constructs.Construct, id string) awss3.IBucket {
	return awss3.NewBucket(scope, jsii.String(id), &awss3.BucketProps{
		EnforceSSL: jsii.Bool(true),
		LifecycleRules: &[]*awss3.LifecycleRule{
			{
				Enabled:    jsii.Bool(true),
				Expiration: cdk.Duration_Days(jsii.Number(7)),
			},
		},
		RemovalPolicy: cdk.RemovalPolicy_DESTROY,
	})
}
