package kvstore

import "context"

type KvStore interface {
	Has(ctx context.Context, key string) (bool, error)
	Get(ctx context.Context, key string, value any) error
	GetAll(ctx context.Context, values map[string]any) error
	GetMany(ctx context.Context, keys []string, values map[string]any) error
	Set(ctx context.Context, key string, value any) error
	SetMany(ctx context.Context, data map[string]any) error
	Delete(ctx context.Context, key string) error
	DeleteMany(ctx context.Context, keys []string) error
	DeleteAll(ctx context.Context) error
	Transaction(ctx context.Context, fn func(tx KvStore) error) error
}
