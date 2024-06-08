package codebuild

import "github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"

func NewDefaultBuildEnv() *awscodebuild.BuildEnvironment {
	return &awscodebuild.BuildEnvironment{
		BuildImage:  awscodebuild.LinuxBuildImage_AMAZON_LINUX_2_ARM_3(),
		ComputeType: awscodebuild.ComputeType_SMALL,
	}
}

func NewUbuntuBuildEnv() *awscodebuild.BuildEnvironment {
	return &awscodebuild.BuildEnvironment{
		BuildImage:  awscodebuild.LinuxBuildImage_STANDARD_7_0(),
		ComputeType: awscodebuild.ComputeType_SMALL,
	}
}

func NewDefaultBuildRuntimes() awscodebuild.BuildSpec {
	return awscodebuild.BuildSpec_FromObject(&map[string]interface{}{
		"phases": map[string]interface{}{
			"install": map[string]interface{}{
				"runtime-versions": map[string]interface{}{
					"nodejs": "20",
					"golang": "1.21",
				},
			},
		},
		"cache": map[string]interface{}{
			"paths": []string{
				"/root/.cache/go-build",
			},
		},
	})
}
