package main

import (
	"fmt"

	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:           "dev",
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Print("error")

	}
	svc := s3.New(sess)
	//get s3 items to resp
	source := "taskawsbucket-hosting"
	other := "task-dest"
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(source)})
	for _, item := range resp.Contents {
		//fmt.Println("Name:         ", *item.Key)
		ob1 := source + "/" + (*item.Key)
		_, err = svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(other),
			CopySource: aws.String(url.PathEscape(ob1)), Key: aws.String(*item.Key)})
		//fmt.Println("Copying")
	}

}
