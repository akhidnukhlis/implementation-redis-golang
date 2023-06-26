package repository

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// AttendanceRepository interface defines the methods for attendance data operations
type AttendanceRepository interface {
	CheckIn(userID string) error
	GetAttendance(userID string, limit int64) ([]string, error)
	Close() error
}

// RedisAttendanceRepository implements AttendanceRepository using Redis client library
type RedisAttendanceRepository struct {
	client *redis.Client
}

func NewRedisAttendanceRepository(addr, password string) AttendanceRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return &RedisAttendanceRepository{
		client: client,
	}
}

func (r *RedisAttendanceRepository) CheckIn(userID string) error {
	key := fmt.Sprintf("attendance:%s", userID)
	score := float64(time.Now().Unix())

	member := redis.Z{
		Score:  score,
		Member: userID,
	}

	_, err := r.client.ZAdd(key, member).Result()
	if err != nil {
		return fmt.Errorf("Failed to check-in: %s", err)
	}

	return nil
}

func (r *RedisAttendanceRepository) GetAttendance(userID string, limit int64) ([]string, error) {
	key := fmt.Sprintf("attendance:%s", userID)
	return r.client.ZRevRange(key, 0, limit-1).Result()
}

func (r *RedisAttendanceRepository) Close() error {
	return r.client.Close()
}
