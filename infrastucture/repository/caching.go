package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type Caching struct {
	redis *redis.Client
}

func NewCaching(redis *redis.Client) *Caching {
	return &Caching{
		redis: redis,
	}
}

func (c *Caching) CacheNewsFeed(userId int, postScores map[int]time.Time, expiresAt int) error {
	key := fmt.Sprintf("newsfeed:%d", userId)
	ctx := context.TODO()

	if err := c.redis.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("failed to clear old cache: %w", err)
	}

	var zMembers []*redis.Z
	for postId, createdAt := range postScores {
		zMembers = append(zMembers, &redis.Z{
			Score:  float64(createdAt.Unix()),
			Member: postId,
		})
	}

	if err := c.redis.ZAdd(ctx, key, zMembers...).Err(); err != nil {
		return fmt.Errorf("failed to cache newsfeed zset: %w", err)
	}

	if err := c.redis.Expire(ctx, key, time.Duration(expiresAt)*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to set expiration: %w", err)
	}

	return nil
}

func (c *Caching) GetNewsFeedIds(userId int) ([]int, error) {
	key := fmt.Sprintf("newsfeed:%d", userId)
	ctx := context.TODO()

	postList, err := c.redis.ZRevRange(ctx, key, 0, -1).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to get newsfeed from redis: %w", err)
	}

	postIds := make([]int, 0, len(postList))
	for _, val := range postList {
		var id int
		fmt.Sscanf(val, "%d", &id)
		postIds = append(postIds, id)
	}

	return postIds, nil
}

func (c *Caching) CachePost(userId int, postId int, scores float64, expiresAt int) error {
	key := fmt.Sprintf("newsfeed:%d", userId)
	ctx := context.TODO()

	members := &redis.Z{
		Score:  scores,
		Member: postId,
	}

	if err := c.redis.ZAdd(ctx, key, members).Err(); err != nil {
		return fmt.Errorf("failed to cache newsfeed zset: %w", err)
	}

	if err := c.redis.Expire(ctx, key, time.Duration(expiresAt)*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to set expiration: %w", err)
	}

	return nil
}
