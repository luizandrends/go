package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type store struct {
	rdb *redis.Client
}

type Store interface {
	SaveShortenedURL(ctx context.Context, _url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

func NewStore(rdb *redis.Client) Store {
	return store{rdb}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	for range 5 {
		code := genCode()
		err := s.rdb.HGet(ctx, "encurtador", code).Err()
		if err == nil {
			continue
		}
		if !errors.Is(err, redis.Nil) {
			return "", fmt.Errorf("Failed to get code from encurtador hashmap: %w", err)
		}

		if err := s.rdb.HSet(ctx, "encurtador", code, _url).Err(); err != nil {
			return "", fmt.Errorf("failed to save shortened URL: %w", err)
		}

		return code, nil
	}

	return "", errors.New("failed to generate an available code")
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	fullUrl, err := s.rdb.HGet(ctx, "encurtador", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get code from encurtador hashmap: %w", err)
	}

	return fullUrl, nil
}
