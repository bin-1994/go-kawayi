package utils

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"go-biu/library/database/cache"
	"time"
)

type RedisClient struct {
	Client *redis.Client
}

var ctx = context.Background()

func NewRedis() *RedisClient {
	return &RedisClient{Client: cache.RedisClient()}
}

func (rdb *RedisClient) Get(key string, data any) error {
	value, err := rdb.Client.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(value, &data); err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) Set(key string, data any, exprie time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if _, err := rdb.Client.Set(ctx, key, value, exprie).Result(); err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) Exist(key string) bool {
	result, err := rdb.Client.Exists(ctx, key).Result()
	if err != nil {
		return false
	}
	return result == 1
}

func (rdb *RedisClient) Incr(key string) int64 {
	value, err := rdb.Client.Incr(ctx, key).Result()
	if err != nil {
		return 0
	}
	return value
}

func (rdb *RedisClient) IncrBy(key string, incr int64) int64 {
	value, err := rdb.Client.IncrBy(ctx, key, incr).Result()
	if err != nil {
		return 0
	}
	return value
}

func (rdb *RedisClient) Decr(key string) int64 {
	value, err := rdb.Client.Decr(ctx, key).Result()
	if err != nil {
		return 0
	}
	return value
}

func (rdb *RedisClient) DecrBy(key string, incr int64) int64 {
	value, err := rdb.Client.DecrBy(ctx, key, incr).Result()
	if err != nil {
		return 0
	}
	return value
}

func (rdb *RedisClient) Del(key string) error {
	count, err := rdb.Client.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	if count == 0 {
		return nil
	}
	return nil
}

func (rdb *RedisClient) Expire(key string, expire time.Duration) bool {
	_, err := rdb.Client.Expire(ctx, key, expire).Result()
	if err != nil {
		return false
	}
	return true
}

// hash
func (rdb *RedisClient) HSet(key, field string, data any) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = rdb.Client.HSet(ctx, key, field, value).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) HGet(key, field string, data any) error {
	value, err := rdb.Client.HGet(ctx, key, field).Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(value, &data)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) HGetAll(key string) map[string]string {
	data, err := rdb.Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil
	}
	return data
}

func (rdb *RedisClient) HKeys(key string) []string {
	fields, err := rdb.Client.HKeys(ctx, key).Result()
	if err != nil {
		return nil
	}
	return fields
}

func (rdb *RedisClient) HLen(key string) int64 {
	size, err := rdb.Client.HLen(ctx, key).Result()
	if err != nil {
		return 0
	}
	return size
}

func (rdb *RedisClient) HDel(key string, fields ...string) error {
	err := rdb.Client.HDel(ctx, key, fields...).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) HExists(key string, field string) bool {
	result, err := rdb.Client.HExists(ctx, key, field).Result()
	if err != nil {
		return false
	}
	return result
}

// 添加有序集合成员
func (rdb *RedisClient) ZAdd(key string, score float64, value string) error {
	err := rdb.Client.ZAdd(ctx, key, redis.Z{Score: score, Member: value}).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取有序集合的成员数
func (rdb *RedisClient) ZCard(key string) int64 {
	result, _ := rdb.Client.ZCard(ctx, key).Result()

	return result
}

// 移除区间成员
func (rdb *RedisClient) ZRemRangeByScore(key string, min, max string) error {
	err := rdb.Client.ZRemRangeByScore(ctx, key, min, max).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RedisClient) ZIncrBy(key string, score float64, value string) error {
	err := rdb.Client.ZIncrBy(ctx, key, score, value).Err()
	if err != nil {
		return err
	}
	return nil
}

// ZRangeByScoreWithScores 获取Redis有序集合中指定分数范围内的成员及其分数
func (rdb *RedisClient) ZRangeByScoreWithScores(key string, min, max string) []redis.Z {
	result, err := rdb.Client.ZRangeByScoreWithScores(ctx, key, &redis.ZRangeBy{Min: min, Max: max}).Result()
	if err != nil {
		return nil
	}
	return result
}

func (rdb *RedisClient) ZUnionStore(key string, keys []string, weights []float64) bool {
	err := rdb.Client.ZUnionStore(ctx, key, &redis.ZStore{
		Keys:    keys,
		Weights: weights,
	}).Err()
	if err != nil {
		return false
	}
	return true
}

// ZRevRangeWithScores 获取Redis有序集合中指定排名范围内的成员及其分数（从高到低排序）
func (rdb *RedisClient) ZRevRangeWithScores(key string, start, stop int64) []redis.Z {
	result, err := rdb.Client.ZRevRangeWithScores(ctx, key, start, stop).Result()
	if err != nil {
		return nil
	}
	return result
}

func (rdb *RedisClient) ZRem(key string, member string) error {
	err := rdb.Client.ZRem(ctx, key, member).Err()
	if err != nil {
		return err
	}
	return nil
}
