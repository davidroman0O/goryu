package goryu

import (
	"context"
	"time"

	hero "github.com/deliveryhero/pipeline/v2"
)

// Just like `Drain` except it will listen to cancellation for teardown
func Cancel[Item any](ctx context.Context, cancel func(Item, error), in <-chan Item) <-chan Item {
	return hero.Cancel(ctx, cancel, in)
}

// Wait before sending input channel
func Delay[Item any](ctx context.Context, duration time.Duration, in <-chan Item) <-chan Item {
	return hero.Delay(ctx, duration, in)
}
