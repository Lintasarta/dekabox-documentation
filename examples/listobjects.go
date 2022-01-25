package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

func main() {
	endpoint := os.Getenv("ENDPOINT")
	bucket := os.Getenv("BUCKET")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("auto"),
		Endpoint: aws.String(endpoint),
	}))

	config := aws.NewConfig()
	config.WithDisableEndpointHostPrefix(*aws.Bool(true))
	config.WithS3ForcePathStyle(*aws.Bool(true))
	client := s3.New(sess, config)

	x := s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	}
	result, err := client.ListObjects(&x)
	if err != nil {
		log.Fatalf("error : %s", err.Error())
	}

	for _, content := range result.Contents {
		log.Infof("DisplayName %s", *content.Key)
		log.Infof("Storage Type %s", *content.StorageClass)
		log.Infof("Size %s", *content.Size)
	}
}
