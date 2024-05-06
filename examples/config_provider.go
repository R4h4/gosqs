package example

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/qhenkart/gosqs"
)

func main_with_config_provider() {

	// implement a custom AWS session provider function
	provider := func(c gosqs.Config) (aws.Config, error) {

		// note: this implementation just hardcodes key and secret, but it could do anything
		creds := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider("mykey", "mysecret", ""))
		_, err := creds.Retrieve(context.TODO())
		if err != nil {
			return aws.Config{}, gosqs.ErrInvalidCreds.Context(err)
		}

		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithRegion("us-west-1"),
			config.WithCredentialsProvider(creds),
		)
		if err != nil {
			return aws.Config{}, err
		}

		hostname := "http://localhost:4150"
		cfg.BaseEndpoint = &hostname

		return cfg, nil
	}

	// create the gosqs Config with our custom SessionProviderFunc
	c := gosqs.Config{
		// for emulation only
		// Hostname: "http://localhost:4150",

		AwsConfigProvider: provider,
		TopicARN:          "arn:aws:sns:local:000000000000:dispatcher",
		Region:            "us-west-1",
	}

	//follows the flow to see how a worker should be configured and operate
	initWorker(c)

	//follows the flow to see how an http service should be configured and operate
	initService(c)

}
