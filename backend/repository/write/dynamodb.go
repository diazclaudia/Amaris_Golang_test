package write

import (
	"backend/repository/read"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	repoBroker "backend/repository"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const tableName = "points"

type repository struct {
	client  *dynamodb.Client
	db      string
	timeout time.Duration
}

type Table struct {
	client  *dynamodb.Client
	db      string
	timeout time.Duration
}

func NewRepository(serverURL, db, port string, timeout int) (*repository, *dynamodb.Client, error) {
	mongoClient, err := NewClient(serverURL, port)
	repo := &repository{
		client:  mongoClient,
		db:      db,
		timeout: time.Duration(timeout) * time.Second,
	}
	if err != nil {
		return nil, nil, err
	}

	return repo, mongoClient, nil
}

func NewClient(url, port string) (*dynamodb.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: fmt.Sprintf("%s:%v", url, port)}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "key", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
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

func (r repository) Update(id, points, key, secret string) (*string, error) {
	item, err := r.client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
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
	repoRead, _, err := read.NewRepositoryRead(key, secret)
	if err != nil {
		panic("error creating dynamodb client")
	}
	repoBroker.PublishData(id, points, repoRead)
	return &status, err
}
