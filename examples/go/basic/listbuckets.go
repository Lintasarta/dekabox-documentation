/*
 * This example to show demonstrate how to get bucket name and creation date
 * from SDK
 */

package basic

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

func ListBuckets() {
	endpoint := os.Getenv("ENDPOINT")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("auto"),
		Endpoint: aws.String(endpoint),
	}))

	config := aws.NewConfig()
	config.WithDisableEndpointHostPrefix(*aws.Bool(true))
	config.WithS3ForcePathStyle(*aws.Bool(true))
	client := s3.New(sess, config)

	input := s3.ListBucketsInput{}

	result, err := client.ListBuckets(&input)
	if err != nil {
		log.Fatalf("error : %s", err.Error())
	}
	for _, content := range result.Buckets {
		log.Infof("Name %s", *content.Name)
		log.Infof("Creation Date %s", *content.CreationDate)
	}
}
