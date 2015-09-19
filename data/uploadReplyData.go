package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func UploadReplyData(svc *dynamodb.DynamoDB) (*dynamodb.BatchWriteItemOutput, error) {
	oneDayAgo := time.Now().AddDate(0, 0, -1).String()
	sevenDaysAgo := time.Now().AddDate(0, 0, -7).String()
	fourteenDaysAgo := time.Now().AddDate(0, 0, -14).String()
	twentyOneDaysAgo := time.Now().AddDate(0, 0, -21).String()

	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"Reply": {
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Id": {
								S: aws.String("Amazon DynamoDB#DynamoDB Thread 1"),
							},
							"ReplyDateTime": {
								S: aws.String(fourteenDaysAgo),
							},
							"Message": {
								S: aws.String("DynamoDB Thread 1 Reply 2 text"),
							},
							"PostedBy": {
								S: aws.String("User B"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Id": {
								S: aws.String("Amazon DynamoDB#DynamoDB Thread 2"),
							},
							"ReplyDateTime": {
								S: aws.String(twentyOneDaysAgo),
							},
							"Message": {
								S: aws.String("DynamoDB Thread 2 Reply 3 text"),
							},
							"PostedBy": {
								S: aws.String("User B"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Id": {
								S: aws.String("Amazon DynamoDB#DynamoDB Thread 2"),
							},
							"ReplyDateTime": {
								S: aws.String(sevenDaysAgo),
							},
							"Message": {
								S: aws.String("DynamoDB Thread 2 Reply 2 text"),
							},
							"PostedBy": {
								S: aws.String("User A"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Id": {
								S: aws.String("Amazon DynamoDB#DynamoDB Thread 2"),
							},
							"ReplyDateTime": {
								S: aws.String(oneDayAgo),
							},
							"Message": {
								S: aws.String("DynamoDB Thread 2 Reply 1 text"),
							},
							"PostedBy": {
								S: aws.String("User A"),
							},
						},
					},
				},
			},
		},
	}

	return svc.BatchWriteItem(params)
}
