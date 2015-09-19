package data

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

func UploadForumData(svc *dynamodb.DynamoDB) (*dynamodb.BatchWriteItemOutput, error) {
	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"Forum": {
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Name": {
								S: aws.String("Amazon DynamoDB"),
							},
							"Category": {
								S: aws.String("Amazon Web Services"),
							},
							"Threads": {
								N: aws.String("0"),
							},
							"Messages": {
								N: aws.String("0"),
							},
							"Views": {
								N: aws.String("1"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Name": {
								S: aws.String("Amazon S3"),
							},
							"Category": {
								S: aws.String("Amazon Web Services"),
							},
							"Threads": {
								N: aws.String("0"),
							},
						},
					},
				},
			},
		},
	}

	return svc.BatchWriteItem(params)
}
