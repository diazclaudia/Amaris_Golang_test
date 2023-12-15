package write

import (
	"backend/config"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var conf, _ = config.NewConfig("/Users/Claudia/Documents/Golang Projects/Amaris_Golang_test/backend/config/config.yaml")

func TestCreateTableIfNotExists(t *testing.T) {
	type args struct {
		d *dynamodb.Client
	}
	_, dynamodbClient, _ := NewRepository("http://localhost", "my-dynamodb", "8000", 5000)
	dynamodbClient.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				d: dynamodbClient,
			},
			want: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTableIfNotExists(tt.args.d); got != tt.want {
				t.Errorf("CreateTableIfNotExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewClient(t *testing.T) {
	type args struct {
		url  string
		port string
	}
	tests := []struct {
		name    string
		args    args
		want    *dynamodb.Client
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				url:  "http://localhost",
				port: "8000",
			},
			want:    &dynamodb.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.url, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
		})
	}
}

func TestNewRepository(t *testing.T) {
	type args struct {
		serverURL string
		db        string
		port      string
		timeout   int
	}
	tests := []struct {
		name    string
		args    args
		want    *repository
		want1   *dynamodb.Client
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				serverURL: "http://localhost",
				db:        "my-dynamodb",
				port:      "8000",
				timeout:   5000,
			},
			want:    &repository{},
			want1:   &dynamodb.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NewRepository(tt.args.serverURL, tt.args.db, tt.args.port, tt.args.timeout)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.NotNil(t, got1)
		})
	}
}

func TestSeedItems(t *testing.T) {
	type args struct {
		dynamodbClient *dynamodb.Client
	}
	_, dynamodbClient, _ := NewRepository("http://localhost", "my-dynamodb", "8000", 5000)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				dynamodbClient: dynamodbClient,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SeedItems(tt.args.dynamodbClient)
		})
	}
}

func Test_buildCreateTableInput(t *testing.T) {
	tests := []struct {
		name string
		want *dynamodb.CreateTableInput
	}{
		{
			name: "success",
			want: &dynamodb.CreateTableInput{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildCreateTableInput(); got != nil {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_putItem(t *testing.T) {
	type args struct {
		d    *dynamodb.Client
		item map[string]types.AttributeValue
	}
	_, dynamodbClient, _ := NewRepository("http://localhost", "my-dynamodb", "8000", 5000)
	item := map[string]types.AttributeValue{
		"id":     &types.AttributeValueMemberS{Value: "1"},
		"points": &types.AttributeValueMemberS{Value: "100"},
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				d:    dynamodbClient,
				item: item,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := putItem(tt.args.d, tt.args.item); (err != nil) != tt.wantErr {
				t.Errorf("putItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repository_Update(t *testing.T) {
	type fields struct {
		client  *dynamodb.Client
		db      string
		timeout time.Duration
	}
	type args struct {
		id     string
		points string
	}
	want := "updated"
	_, dynamodbClient, _ := NewRepository("http://localhost", "my-dynamodb", "8000", 5000)
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *string
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: dynamodbClient,
			},
			args: args{
				id:     "1",
				points: "200",
			},
			want:    &want,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository{
				client:  tt.fields.client,
				db:      tt.fields.db,
				timeout: tt.fields.timeout,
			}
			got, err := r.Update(tt.args.id, tt.args.points, conf.Aws.Key, conf.Aws.Secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tableExists(t *testing.T) {
	type args struct {
		d    *dynamodb.Client
		name string
	}
	_, dynamodbClient, _ := NewRepository("http://localhost", "my-dynamodb", "8000", 5000)
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				d: dynamodbClient,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tableExists(tt.args.d, tt.args.name); got != tt.want {
				t.Errorf("tableExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
