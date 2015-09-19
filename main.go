package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/sinistersig/aws-practice-dynamodb-add-data/data"
	"log"
)

func main() {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	svc := dynamodb.New(config)

	//Batch write to forums table
	resp, err := data.UploadForumData(svc)

	if err != nil {
		log.Println("An error occurred while writing to the Forums table")
		log.Println(err.Error())
	}

	log.Println(resp)

	//Batch write to Reply Table
	resp, err = data.UploadReplyData(svc)

	if err != nil {
		log.Println("An error occurred while writing to the Reply table")
		log.Println(err.Error())
	}

	log.Println(resp)

	//Upload product catalog
	resp, err = data.UploadProductCatalog(svc)

	if err != nil {
		log.Println("An error occurred while writing to the ProductCatalog table")
		log.Println(err.Error())
	}

	log.Println(resp)

}
