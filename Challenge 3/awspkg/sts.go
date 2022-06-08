package awspkg

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func PrintCallerIdentity() string {
	// 👇 Create new api client with new session.
	svc := sts.New(session.New())

	// 👇 Call GetCallerIdentity method provided by api client.
	input := &sts.GetCallerIdentityInput{}

	// 👇 Handle error.
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		}
	}

	// 👇 Return string representation of the result.
	return result.String()
}
