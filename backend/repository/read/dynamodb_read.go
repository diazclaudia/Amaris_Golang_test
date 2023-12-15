package read

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const tableName = "points"

type RepositoryRead struct {
	Client  *dynamodb.Client
	db      string
	timeout time.Duration
}

type TableRead struct {
	client  *dynamodb.Client
	db      string
	timeout time.Duration
}

func (r RepositoryRead) Find(id string) (map[string]types.AttributeValue, error) {
	item, err := r.Client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("not found the item")
	}
	return item.Item, err
}

func NewRepositoryRead(key, secret string) (*RepositoryRead, *dynamodb.Client, error) {
	// Create an AWS configuration with the loaded credentials
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
	)
	cfg.Region = "us-east-1"
	if err != nil {
		fmt.Println("Error loading AWS configuration:", err)
		return nil, nil, err
	}

	// Create a DynamoDB client
	dynamodbClientRead := dynamodb.NewFromConfig(cfg)
	repo := &RepositoryRead{
		Client: dynamodbClientRead,
	}
	return repo, dynamodbClientRead, nil
}

func tableExists(d *dynamodb.Client, name string) bool {
	tables, err := d.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Fatal("ListTables failed", err)
	}
	for _, n := range tables.TableNames {
		if n == name {
			return true
		}
	}
	return false
}

func CreateTableIfNotExists(d *dynamodb.Client) bool {
	if tableExists(d, tableName) {
		return false
	}
	_, err := d.CreateTable(context.TODO(), buildCreateTableInput())
	if err != nil {
		log.Fatal("CreateTable failed", err)
	}
	return true
}

func buildCreateTableInput() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	}
}

func SeedItems(dynamodbClient *dynamodb.Client) {
	CreateTableIfNotExists(dynamodbClient)

	for i := 1; i < 3; i++ {
		item := map[string]types.AttributeValue{
			"id":     &types.AttributeValueMemberS{Value: strconv.Itoa(i)},
			"points": &types.AttributeValueMemberS{Value: "100"},
		}
		err := putItem(dynamodbClient, item)
		if err != nil {
			log.Fatal("failed to put item", err)
		}
	}
}

func putItem(d *dynamodb.Client, item map[string]types.AttributeValue) error {
	_, err := d.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	})
	return err
}

func (r RepositoryRead) Update(id, points string) (*string, error) {
	item, err := r.Client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: aws.String("set points = :points"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":points": &types.AttributeValueMemberS{Value: points},
		},
	})

	if err != nil {
		return nil, err
	}
	if item == nil {
		return nil, errors.New("not found the item")
	}
	status := "updated"
	return &status, err
}
