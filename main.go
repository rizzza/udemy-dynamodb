package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/pkg/errors"
)

type DynamoClient struct {
	cli          *dynamodb.DynamoDB
	table, index string
}

const (
	AWSDefaultRegion = "us-east-1"
)

type DynamoClientExternal interface {
	Test()
}

func NewDynamoClient() (DynamoClientExternal, error) {
	cli := dynamodb.New(
		session.New(&aws.Config{Region: aws.String(AWSDefaultRegion)}),
		aws.NewConfig().WithRegion(AWSDefaultRegion))

	client := &DynamoClient{
		cli:   cli,
		table: "",
		index: "",
	}

	return client, client.Validate(context.Background())
}

func (d *DynamoClient) Test() {
	log.Println("Nothing to see here!")
}

func (d *DynamoClient) Validate(ctx context.Context) error {
	//try to list tables to validate the connection is working on startup
	listTablesInput := &dynamodb.ListTablesInput{
		Limit: aws.Int64(1),
	}

	_, err := d.cli.ListTablesWithContext(ctx, listTablesInput)
	return errors.Wrap(err, "Valdiate ListTablesWithContext")
}

func main() {
	cli, err := NewDynamoClient()
	if err != nil {
		panic("no client for you!")
	}

	cli.Test()
}
