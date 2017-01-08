package profiler

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func (p Profiler) uploadToS3(payload []byte) (err error) {
	sess, err := session.NewSession()

	if err != nil {
		log.Println("failed to create session,", err)
		return
	}

	svc := s3.New(sess)

	params := &s3.PutObjectInput{
		Bucket: aws.String("thomas-vendetta-sandb0x"),
		Key:    aws.String("test-collection-profile.json"), // Required
		Body:   bytes.NewReader(payload),
	}

	_, err = svc.PutObject(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		log.Println(err.Error())
		return
	}

	log.Println("Uploaded to S3")
	return
}
