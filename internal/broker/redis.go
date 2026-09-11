package broker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const UserCreatedChannel = "user-created"

type UserCreatedEvent struct {
	UserID uuid.UUID `json:"user_id"`
}

type RedisBroker struct {
	client *redis.Client
}

func NewRedisBroker(addr string) (*RedisBroker, error) {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}
	return &RedisBroker{client: client}, nil
}

func (b *RedisBroker) SubscribeUserCreated(ctx context.Context, handler func(context.Context, uuid.UUID) error) error {
	sub := b.client.Subscribe(ctx, UserCreatedChannel)
	defer sub.Close()
	if _, err := sub.Receive(ctx); err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}
	ch := sub.Channel()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-ch:
			var event UserCreatedEvent
			if err := json.Unmarshal([]byte(msg.Payload), &event); err != nil {
				log.Printf("failed to unmarshal event: %v", err)
				continue
			}
			if err := handler(ctx, event.UserID); err != nil {
				log.Printf("failed to handle user created event: %v", err)
			}
		}
	}
}

func (b *RedisBroker) Close() error {
	return b.client.Close()
}
