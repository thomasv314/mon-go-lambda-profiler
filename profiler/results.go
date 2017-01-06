package profiler

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/thomasv314/mongo-tools/common/bsonutil"
	"log"
)

func (p Profiler) ResultsAsJSON() (jsonBytes []byte, err error) {
	result, err := p.Results()

	if err != nil {
		return
	}

	jsonArr := make([]interface{}, len(result))

	for r := range result {
		asJson, err := bsonutil.GetBSONValueAsJSON(result[r])
		if err != nil {
			panic(err) // unsure what to do yet
		}

		jsonArr[r] = asJson
	}

	jsonBytes, err = json.Marshal(jsonArr)
	return
}

func (p Profiler) UploadResultsToS3() {
	payload, err := p.ResultsAsJSON()

	if err != nil {
		log.Println("Failed to get results..", err)
		return
	}

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
}
