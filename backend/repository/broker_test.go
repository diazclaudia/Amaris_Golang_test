package repository

import (
	"backend/config"
	"backend/repository/read"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

var conf, _ = config.NewConfig("/Users/Claudia/Documents/Golang Projects/Amaris_Golang_test/backend/config/config.yaml")

func TestBroker_Publish(t *testing.T) {
	type fields struct {
		subscribers map[string][]*Subscriber
		mutex       sync.Mutex
	}
	type args struct {
		topic   string
		payload interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				subscribers: make(map[string][]*Subscriber),
			},
			args: args{
				topic:   "1",
				payload: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				subscribers: tt.fields.subscribers,
				mutex:       tt.fields.mutex,
			}
			b.Publish(tt.args.topic, tt.args.payload)
		})
	}
}

func TestBroker_Subscribe(t *testing.T) {
	type fields struct {
		subscribers map[string][]*Subscriber
		mutex       sync.Mutex
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Subscriber
	}{
		{
			name: "success",
			fields: fields{
				subscribers: make(map[string][]*Subscriber),
			},
			args: args{
				topic: "1",
			},
			want: &Subscriber{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				subscribers: tt.fields.subscribers,
				mutex:       tt.fields.mutex,
			}
			got := b.Subscribe(tt.args.topic)
			assert.NotNil(t, got)
		})
	}
}

func TestBroker_Unsubscribe(t *testing.T) {
	type fields struct {
		subscribers map[string][]*Subscriber
		mutex       sync.Mutex
	}
	type args struct {
		topic      string
		subscriber *Subscriber
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				subscribers: make(map[string][]*Subscriber),
			},
			args: args{
				topic:      "1",
				subscriber: &Subscriber{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Broker{
				subscribers: tt.fields.subscribers,
				mutex:       tt.fields.mutex,
			}
			b.Unsubscribe(tt.args.topic, tt.args.subscriber)
		})
	}
}

func TestNewBroker(t *testing.T) {
	tests := []struct {
		name string
		want *Broker
	}{
		{
			name: "success",
			want: &Broker{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewBroker()
			assert.NotNil(t, got)
		})
	}
}

func TestPublishData(t *testing.T) {
	type args struct {
		id       string
		points   string
		repoRead *read.RepositoryRead
	}
	_, dynamodbClient, _ := read.NewRepositoryRead(conf.Aws.Key, conf.Aws.Secret)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				id:     "1",
				points: "2",
				repoRead: &read.RepositoryRead{
					Client: dynamodbClient,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PublishData(tt.args.id, tt.args.points, tt.args.repoRead)
		})
	}
}
