package main

import (
	"flag"
	"fmt"
	"net/url"
	"sync"
	"time"

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
	
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(*source_buc)})
	var wg sync.WaitGroup
	wg.Add(int(*resp.KeyCount))
	for _, item := range resp.Contents {

		go func() {

			defer wg.Done()
			ob1 := *source_buc + "/" + (*item.Key)

			//fmt.Println(ob1)
			_, err = svc.CopyObject(&s3.CopyObjectInput{Bucket: aws.String(*dest_buc),
				CopySource: aws.String(url.PathEscape(ob1)), Key: aws.String(*item.Key)})

		}()
		time.Sleep(5 * time.Millisecond)
	}
	wg.Wait()

}
