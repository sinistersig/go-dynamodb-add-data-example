package data

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func UploadProductCatalog(svc *dynamodb.DynamoDB) (*dynamodb.BatchWriteItemOutput, error) {
	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			"ProductCatalog": {
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("1101"),
							},
							"Title": {
								S: aws.String("Book 101 Title"),
							},
							"ISBN": {
								S: aws.String("111-1111111111"),
							},
							"Authors": {
								SS: []*string{
									aws.String("Author1"),
								},
							},
							"Price": {
								N: aws.String("2"),
							},
							"Dimensions": {
								S: aws.String("8.5 x 11.0 x 0.5"),
							},
							"PageCount": {
								N: aws.String("500"),
							},
							"InPublication": {
								N: aws.String("1"),
							},
							"ProductCategory": {
								S: aws.String("Book"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("102"),
							},
							"Title": {
								S: aws.String("Book 102 Title"),
							},
							"ISBN": {
								S: aws.String("222-2222222222"),
							},
							"Authors": {
								SS: []*string{
									aws.String("Author1"),
									aws.String("Author2"),
								},
							},
							"Price": {
								N: aws.String("20"),
							},
							"Dimensions": {
								S: aws.String("8.5 x 11.0 x 0.8"),
							},
							"PageCount": {
								N: aws.String("600"),
							},
							"InPublication": {
								N: aws.String("1"),
							},
							"ProductCategory": {
								S: aws.String("Book"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("103"),
							},
							"Title": {
								S: aws.String("Book 103 Title"),
							},
							"ISBN": {
								S: aws.String("333-3333333333"),
							},
							"Authors": {
								SS: []*string{
									aws.String("Author1"),
									aws.String("Author2"),
								},
							},
							"Price": {
								N: aws.String("2000"),
							},
							"Dimensions": {
								S: aws.String("8.5 x 11.0 x 1.5"),
							},
							"PageCount": {
								N: aws.String("600"),
							},
							"InPublication": {
								N: aws.String("0"),
							},
							"ProductCategory": {
								S: aws.String("Book"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("201"),
							},
							"Title": {
								S: aws.String("18-Bike-201"),
							},
							"Description": {
								S: aws.String("201 Description"),
							},
							"BicycleType": {
								S: aws.String("Road"),
							},
							"Brand": {
								S: aws.String("Mountain A"),
							},
							"Price": {
								N: aws.String("100"),
							},
							"Gender": {
								S: aws.String("M"),
							},
							"Color": {
								SS: []*string{
									aws.String("Red"),
									aws.String("Black"),
								},
							},
							"ProductCategory": {
								S: aws.String("Bicycle"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("202"),
							},
							"Title": {
								S: aws.String("21-Bike-202"),
							},
							"Description": {
								S: aws.String("202 Description"),
							},
							"BicycleType": {
								S: aws.String("Road"),
							},
							"Brand": {
								S: aws.String("Brand-Company A"),
							},
							"Price": {
								N: aws.String("200"),
							},
							"Gender": {
								S: aws.String("M"),
							},
							"Color": {
								SS: []*string{
									aws.String("Green"),
									aws.String("Black"),
								},
							},
							"ProductCategory": {
								S: aws.String("Bicycle"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"Id": {
								N: aws.String("203"),
							},
							"Title": {
								S: aws.String("19-Bike-203"),
							},
							"Description": {
								S: aws.String("203 Description"),
							},
							"BicycleType": {
								S: aws.String("Road"),
							},
							"Brand": {
								S: aws.String("Brand-Company B"),
							},
							"Price": {
								N: aws.String("300"),
							},
							"Gender": {
								S: aws.String("W"),
							},
							"Color": {
								SS: []*string{
									aws.String("Green"),
									aws.String("Black"),
									aws.String("Red"),
								},
							},
							"ProductCategory": {
								S: aws.String("Bicycle"),
							},
						},
					},
				},
				&dynamodb.WriteRequest{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{

							"Id": {
								N: aws.String("205"),
							},
							"Title": {
								S: aws.String("20-Bike-205"),
							},
							"Description": {
								S: aws.String("205 Description"),
							},
							"BicycleType": {
								S: aws.String("Hybrid"),
							},
							"Brand": {
								S: aws.String("Brand-Company C"),
							},
							"Price": {
								N: aws.String("500"),
							},
							"Gender": {
								S: aws.String("B"),
							},
							"Color": {
								SS: []*string{
									aws.String("Red"),
									aws.String("Black"),
								},
							},
							"ProductCategory": {
								S: aws.String("Bicycle"),
							},
						},
					},
				},
			},
		},
	}

	return svc.BatchWriteItem(params)
}
