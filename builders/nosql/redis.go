package nosql

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/go-redis/redis/v7"
)

type Redis struct {
	Network  string `json:"network,omitempty" desc:"enum: tcp|unix, default: tcp" validate:"-"`
	Addr     string `json:"addr" desc:"redis address" validate:"required"`
	Password string `json:"password" desc:"redis password" validate:"-"`
	DB       int    `json:"db" desc:"database index" validate:"gte=0"`
	PoolSize int    `json:"pool_size,omitempty" desc:"maximum number of socket connections, default: 10 * num_cpu" validate:"gte=0"`
	//MinIdleConns      int    `json:"min_idle_conns,omitempty" desc:"minimum number of idle connections which is useful when establishing new connection is slow" validate:"gte=0"`
	//MaxConnSecond     int    `json:"max_conn_age,omitempty" desc:"connection second at which client retires (closes) the connection" validate:"get=0"`
	//PoolTimeoutSecond int    `json:"pool_timeout_second,omitempty" desc:"mount of time client waits for connection if all connections are busy before returning an error" validate:"gte=0"`
	//IdleTimeoutSecond int    `json:"idle_timeout_second" desc:"amount of time after which client closes idle connections"`
	client *redis.Client
}

func (r Redis) Build() (pipeline.Handler, error) {
	client := redis.NewClient(&redis.Options{
		Network:  r.Network,
		Addr:     r.Addr,
		Password: r.Password,
		DB:       0,
		PoolSize: r.PoolSize,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	fmt.Println("redis ping:", pong)
	return &Redis{
		Network:  r.Network,
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
		PoolSize: r.PoolSize,
		client:   client,
	}, nil
}

func (r Redis) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		reqRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			args := reqRes.Data.([]interface{})
			result, err := r.client.Do(args...).String()
			if err != nil {
				return nil, err
			}
			respRes.Data = result
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
