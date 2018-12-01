package redis

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
)

var (
	testCfg = Config{
		Addr: "localhost:6379",
	}

	cli = redis.NewClient(testCfg.Options())
)

func TestRedisSSRDBImpl_Put(t *testing.T) {
	type fields struct {
		cli *redis.Client
	}
	type args struct {
		key    string
		value  string
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "put",
			fields: fields{
				cli: cli,
			},
			args: args{
				key:    "test-key",
				value:  "test-value",
				expire: time.Second * 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &RedisSSRDBImpl{
				cli: tt.fields.cli,
			}
			if err := impl.Put(tt.args.key, tt.args.value, tt.args.expire); (err != nil) != tt.wantErr {
				t.Errorf("RedisSSRDBImpl.Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisSSRDBImpl_Get(t *testing.T) {
	type fields struct {
		cli *redis.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "get",
			fields: fields{
				cli: cli,
			},
			args: args{
				key: "test-key",
			},
			want: "test-value",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			impl := &RedisSSRDBImpl{
				cli: tt.fields.cli,
			}
			got, err := impl.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisSSRDBImpl.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisSSRDBImpl.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
