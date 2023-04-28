package goryu

import (
	"context"
	"time"

	hero "github.com/deliveryhero/pipeline/v2"
)

// If you need to call outside of the application and except an error
// otherwise use Decompose
func Process[Input, Output any](
	ctx context.Context,
	processor hero.Processor[Input, Output],
	in <-chan Input) <-chan Output {
	return hero.Process(ctx, processor, in)
}

func NewProcess[Input, Output any](
	process func(ctx context.Context, i Input) (Output, error),
	cancel func(i Input, err error),
) hero.Processor[Input, Output] {
	return hero.NewProcessor(process, cancel)
}

func ProcessBatch[Input, Output any](
	ctx context.Context,
	maxSize int,
	maxDuration time.Duration,
	processor hero.Processor[[]Input, []Output],
	in <-chan Input,
) <-chan Output {
	return hero.ProcessBatch(ctx, maxSize, maxDuration, processor, in)
}

func ProcessBatchConcurrently[Input, Output any](
	ctx context.Context,
	concurrently,
	maxSize int,
	maxDuration time.Duration,
	processor hero.Processor[[]Input, []Output],
	in <-chan Input,
) <-chan Output {
	return hero.ProcessBatchConcurrently(ctx, concurrently, maxSize, maxDuration, processor, in)
}

func ProcessConcurrently[Input, Output any](ctx context.Context, concurrently int, p hero.Processor[Input, Output], in <-chan Input) <-chan Output {
	return hero.ProcessConcurrently(ctx, concurrently, p, in)
}
