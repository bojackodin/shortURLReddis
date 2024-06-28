package model

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type URLRepository interface {
	Add(originalURL, shortKey string) error
	Get(shortKey string) (string, error)
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

func (m *redisClient) Add(originalURL, shortKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.client.Set(ctx, shortKey, originalURL, 1*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}

func (m *redisClient) Get(shortKey string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	url, err := m.client.Get(ctx, shortKey).Result()
	if err != nil {
		return "", err
	}
	return url, nil
}
