package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	profile := flag.String("p", "defualt", "Mentoin profile")
	source_buc := flag.String("s", "", "Enter source bucket name")
	dest_buc := flag.String("d", "", "Enter destination bucket name")
	flag.Parse()
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile:           *profile,
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		fmt.Print("error")

	}
	svc := s3.New(sess)
	//get s3 items to resp
	// source := "taskawsbucket-hosting"
	// other := "task-dest"
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(*source_buc)})
	for _, item := range resp.Contents {
		//fmt.Println("Name:         ", *item.Key)
		ob1 := *source_buc + "/" + (*item.Key)
		_, err = svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(*dest_buc),
			CopySource: aws.String(url.PathEscape(ob1)), Key: aws.String(*item.Key)})
		//fmt.Println("Copying")
	}

}
