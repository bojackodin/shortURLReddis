package model

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type URLRepository interface {
	Add(ctx context.Context, originalURL, shortKey string) error
	Get(ctx context.Context, shortKey string) (string, error)
}

type URLPair struct {
	Original string `bson:"original_url"`
	Short    string `bson:"short_url"`
}
type redisClient struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) URLRepository {
	return &redisClient{
		client: client,
	}
}

func (m *redisClient) Add(ctx context.Context, originalURL, shortKey string) error {
	err := m.client.Set(ctx, shortKey, originalURL, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m *redisClient) Get(ctx context.Context, shortKey string) (string, error) {
	url, err := m.client.Get(ctx, shortKey).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
