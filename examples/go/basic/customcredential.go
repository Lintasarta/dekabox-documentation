/*
 * This example to show how to use custom credential in apps
 * usually S3 SDK search for file ~/.aws/credential in system
 */

package basic

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	log "github.com/sirupsen/logrus"
)

func CustomCredential() {
	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	endpoint := os.Getenv("ENDPOINT")
	credential := credentials.NewStaticCredentials(accessKey, secretKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credential,
		Region:      aws.String("auto"),
		Endpoint:    aws.String(endpoint),
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
