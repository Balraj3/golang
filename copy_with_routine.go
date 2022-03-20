package main

import (
	"fmt"
	"net/url"
	"sync"
	"time"

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
	var wg sync.WaitGroup
	wg.Add(int(*resp.KeyCount))
	for _, item := range resp.Contents {

		go func() {

			defer wg.Done()
			ob1 := source + "/" + (*item.Key)

			//fmt.Println(ob1)
			_, err = svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(other),
				CopySource: aws.String(url.PathEscape(ob1)), Key: aws.String(*item.Key)})

		}()
		time.Sleep(5 * time.Millisecond)
	}
	wg.Wait()

}
