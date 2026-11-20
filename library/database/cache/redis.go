package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-biu/library/config"
	"log"
	"time"
)

var client *redis.Client

func InitRedisConn() *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf("%s:%d", config.YamlConf.Redis.Host, config.YamlConf.Redis.Port),
		Password:    config.YamlConf.Redis.Password,
		DB:          0,
		PoolSize:    50,
		DialTimeout: 10 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatal("NewRedisConnError:", err)
	}

	return client
}

func RedisClient() *redis.Client {
	return client
}

//type RedisConn struct {
//	redis.Conn
//}
//
//func RedisConnect() redis.Conn {
//	host := fmt.Sprintf("%s:%d", config.YamlConf.Redis.Host, config.YamlConf.Redis.Port)
//	conn, err := redis.Dial("tcp", host)
//	if err != nil {
//		panic(err)
//	}
//	return conn
//}
//
//func (r *RedisConn) Set(key string, value interface{}, expire int) bool {
//	_, err := r.Do("SET", key, value)
//	if err != nil {
//		return false
//	}
//	r.Do("EXPIRE", key, expire)
//	return true
//}
//
//func (r *RedisConn) Get(key string) interface{} {
//	reply, err := r.Do("GET", key)
//	if err != nil {
//		panic(err)
//	}
//	return reply
//}
//
//// 按key值删除redis
//func (r *RedisConn) Delete(key string) {
//	_, err := r.Do("DEL", key)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (r *RedisConn) Zadd(key string, value int, field string) {
//	_, err := r.Do("ZADD", key, value, field)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func (r *RedisConn) GetSortSetLength(key string) interface{} {
//	reply, err := r.Do("ZCARD", key)
//	if err != nil {
//		panic(err)
//	}
//	return reply
//}
//
//func (r *RedisConn) GetRankSort(key string) {
//
//}
//
//func (r *RedisConn) Zunionstore(key string, args ...[]string) {
//	var weights []int
//	for _, _ = range args {
//		weights = append(weights, 1)
//	}
//	r.Do("ZUNIONSTORE", key, len(args), args, "WEIGHTS", weights)
//}

//var Redis *redis.Pool
//
//func NewRedis() error {
//	host := fmt.Sprintf("%s:%d", config.YamlConf.Redis.Host, config.YamlConf.Redis.Port)
//	Redis = &redis.Pool{
//		// 最大空闲连接数
//		MaxIdle: 5,
//		// 在给定时间内，允许分配的最大连接数（当为零时，没有限制）
//		MaxActive: 50,
//		// 在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
//		IdleTimeout: 240 * time.Second,
//		Dial: func() (redis.Conn, error) {
//			c, err := redis.Dial("tcp", host)
//			if err != nil {
//				panic(err)
//			}
//			return c, nil
//		},
//	}
//	return nil
//}
//
//func Set(key string, data interface{}, time int) (bool, error) {
//	conn := Redis.Get()
//	defer conn.Close()
//	value, err := json.Marshal(data)
//	if err != nil {
//		return false, err
//	}
//	reply, err := redis.Bool(conn.Do("SET", key, value))
//	conn.Do("EXPIRE", key, time)
//	return reply, err
//}
//
//func Get(key string) ([]byte, error) {
//	conn := Redis.Get()
//	defer conn.Close()
//	reply, err := redis.Bytes(conn.Do("GET", key))
//	if err != nil {
//		return nil, err
//	}
//	return reply, nil
//}
//
//func Exists(key string) bool {
//	conn := Redis.Get()
//	defer conn.Close()
//	exists, err := redis.Bool(conn.Do("EXISTS", key))
//	if err != nil {
//		return false
//	}
//	return exists
//}
//
//func Delete(key string) (bool, error) {
//	conn := Redis.Get()
//	defer conn.Close()
//	return redis.Bool(conn.Do("DEL", key))
//}
//
//func Incr(key string, time int) int {
//	conn := Redis.Get()
//	defer conn.Close()
//	value, err := redis.Int(conn.Do("INCR", key))
//	if err != nil {
//		return 0
//	}
//	conn.Do("EXPIRE", key, time)
//	return value
//}
//
//func SetSort(key string, value string) (bool, error) {
//	conn := Redis.Get()
//	defer conn.Close()
//	reply, err := redis.Bool(conn.Do("ZADD", key, value))
//	if err != nil {
//		return false, err
//	}
//	return reply, nil
//}
