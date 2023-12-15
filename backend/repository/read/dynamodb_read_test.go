package read

import (
	"backend/config"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/stretchr/testify/assert"
)

// Ubicación del yaml con la config
var conf, _ = config.NewConfig("/Users/Claudia/Documents/Golang Projects/Amaris_Golang_test/backend/config/config.yaml")

func TestCreateTableIfNotExists(t *testing.T) {
	type args struct {
		d *dynamodb.Client
	}
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
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
			got := CreateTableIfNotExists(tt.args.d)
			assert.NotNil(t, got)
		})
	}
}

func TestNewRepositoryRead(t *testing.T) {
	type args struct {
		key    string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    *RepositoryRead
		want1   *dynamodb.Client
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				key:    conf.Aws.Key,
				secret: conf.Aws.Secret,
			},
			want:    &RepositoryRead{},
			want1:   &dynamodb.Client{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := NewRepositoryRead(tt.args.key, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRepositoryRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.NotNil(t, got)
			assert.NotNil(t, got1)
		})
	}
}

func TestRepositoryRead_Find(t *testing.T) {
	type fields struct {
		client  *dynamodb.Client
		db      string
		timeout time.Duration
	}
	type args struct {
		id string
	}
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
	item := map[string]types.AttributeValue{
		"id":     &types.AttributeValueMemberS{Value: "1"},
		"points": &types.AttributeValueMemberS{Value: "100"},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]types.AttributeValue
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: dynamodbClient,
			},
			args: args{
				id: "1",
			},
			want:    item,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryRead{
				Client:  tt.fields.client,
				db:      tt.fields.db,
				timeout: tt.fields.timeout,
			}
			_, err := r.Find(tt.args.id)
			assert.Nil(t, err)

		})
	}
}

func TestRepositoryRead_Update(t *testing.T) {
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
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RepositoryRead{
				Client:  tt.fields.client,
				db:      tt.fields.db,
				timeout: tt.fields.timeout,
			}
			_, err := r.Update(tt.args.id, tt.args.points)
			assert.Nil(t, err)
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
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
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
			err := putItem(tt.args.d, tt.args.item)
			assert.Nil(t, err)
		})
	}
}

func Test_tableExists(t *testing.T) {
	type args struct {
		d    *dynamodb.Client
		name string
	}
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
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

func TestSeedItems(t *testing.T) {
	type args struct {
		dynamodbClient *dynamodb.Client
	}
	_, dynamodbClient, _ := NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
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
