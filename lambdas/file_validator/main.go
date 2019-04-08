package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"strings"
)

func isJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

func uploadToS3(content, key string) error {
	if key == "" {
		return errors.New("invalid file key")
	}
	conf := aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))}
	sess, err := session.NewSession(&conf)
	if err != nil {
		return errors.New("error starting s3 key")
	}
	svc := s3manager.NewUploader(sess)

	fmt.Println("Uploading file to S3...")
	_, err = svc.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("DEST_S3_BUCKET")), //"lambda-demo-09-18"
		Key:         aws.String(key),
		ContentType: aws.String("application/json"),
		Body:        strings.NewReader(content),
	})
	return err
}

func upload(e events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if content := e.Body; isJSON(content) {
		err := uploadToS3(content, e.PathParameters["item"])
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       http.StatusText(500),
			}, err
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 201,
			Body:       http.StatusText(201),
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       http.StatusText(400),
		}, errors.New("Not JSON")
	}
}

func main() {
	lambda.Start(upload)
}
