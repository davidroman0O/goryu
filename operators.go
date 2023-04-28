package goryu

import (
	"context"
	"time"

	brad "github.com/bradfair/go-pipelines"
	hero "github.com/deliveryhero/pipeline/v2"
	sol "github.com/solsw/chanhelper"
)

/// Gochan compiles (for now) two libraries of helpers into one
/// I use both already so why not merge them all

// Buffer converts a given channel into a buffered channel.
// Buffer creates a buffered channel that will close after the input is closed and the buffer is fully drained
// [xxxx]--->
func Buffer[Item any](
	ctx context.Context,
	input <-chan Item,
	count int) <-chan Item {
	return brad.Buffer(ctx, input, count)
}

// Merge fans multiple channels in to a single channel
// Merges the input channels into a single output channel.
// s1:            -a--b----c--->
// s2:            --w---x-y--z->
// Merge(s1, s2): -aw-b-x-yc-z->
func Merge[Item any](ctx context.Context, inputChannels ...<-chan Item) <-chan Item {
	return brad.Merge(ctx, inputChannels...)
}

// Take one channel of type X and a function that you write to create a channel of type Y WHILE converting multiple items into Y type
func Aggregate[Input any, Output any](
	ctx context.Context,
	input <-chan Input,
	count int,
	f func(ctx context.Context, input ...Input) Output) <-chan Output {
	return brad.AggregatorFunc(ctx, input, count, f)
}

// Take one channel of type X and a function that you write to create a channel of type Y
func Decompose[Input, Output any](
	ctx context.Context,
	input <-chan Input,
	f func(input Input) []Output) <-chan Output {
	return brad.DecomposerFunc(ctx, input, f)
}

// Emit fans is ...Item“ out to a <-chan Item`
// tldr: create a container of data of a type into a channel ready to be consumed
func Emit[Item any](
	ctx context.Context,
	is ...Item) <-chan Item {
	return hero.Emit(is...)
}

// Emit fans is ...Item“ out to a <-chan Item`
// tldr: created a container of data of a type from a function that generate that data
func Emitter[Item any](
	ctx context.Context,
	next func() Item) <-chan Item {
	return hero.Emitter(ctx, next)
}

func Slice[Input any](
	ctx context.Context,
	input <-chan Input,
	count int) []Input {
	return brad.ToSlice(ctx, input, count)
}

func Array[Input any](
	ctx context.Context,
	input <-chan Input) []Input {
	out := []Input{}
	oneClose := false
	for {
		if oneClose {
			break
		}
		select {
		case d1, ok := <-input:
			if !ok {
				oneClose = true
				continue
			}
			out = append(out, d1)
		}

	}
	return out
}

func One[Input any](
	ctx context.Context,
	input <-chan Input) Input {
	var one Input
	select {
	case one = <-input:
	}
	return one
}

// Duplicate a channel into two output channels
// Block the reading until both channel got both items
func Duplicate[Input any](
	ctx context.Context,
	input <-chan Input) (_, _ <-chan Input) {
	return brad.Tee(ctx, input)
}

// Already closed channels merged sending signals one open channel will have their output as one
//
//	sX: -ab--X
//	sY: ---f-X
//	b:        -abf->
//
// ```
// inputChs := make(chan (<-chan bool), 2)
//
// inputCh1 := make(chan bool, 1)
// inputCh1 <- true
// close(inputCh1)
//
// inputCh2 := make(chan bool, 1)
// inputCh2 <- true
// close(inputCh2)
//
// inputChs <- inputCh1
// inputChs <- inputCh2
//
// outputCh := Bridge(ctx, inputChs)
// ```
func Bridge[Input any](ctx context.Context, inputsStream <-chan <-chan Input) <-chan Input {
	return brad.Bridge(ctx, inputsStream)
}

// teardown a channel
func Drain[Item any](in <-chan Item) {
	hero.Drain(in)
}

// Takes a channel with an array OR interface from `Collect` then splits it back out into individual elements
//
// s1: 	    -[abc]--[fghj]-->
// flat:                      -a-b-c-f-g-h-j-->
func Flat[Item any](in <-chan []Item) <-chan Item {
	return hero.Split(in)
}

func Collect[Item any](ctx context.Context, maxSize int, maxDuration time.Duration, in <-chan Item) <-chan []Item {
	return hero.Collect(ctx, maxSize, maxDuration, in)
}

// Take returns a channel that closes after receiving the specified number of elements from the specified input channel.
func Take[Input any](ctx context.Context, input <-chan Input, count int) <-chan Input {
	return brad.Take(ctx, input, count)
}

// It returns two channels. One that blocks with the inputs that pass the filter, and one doesn't block with the inputs that do not pass the filter.
func Filter[Input any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) bool) (_, _ <-chan Input) {
	return brad.FilterFunc(ctx, input, f)
}

// OrDone wraps a channel with a done channel and returns a forwarding channel that closes when either the original channel or the done channel closes.
func OrDone[Input any](done <-chan struct{}, inputs <-chan Input) <-chan Input {
	return brad.OrDone(done, inputs)
}

func Repeat[Output any](ctx context.Context, fn func() Output) <-chan Output {
	return brad.RepeatFunc(ctx, fn)
}

// Executes a function on each item in a channel until the channel is closed or the context is cancelled.
func Sink[Input any](ctx context.Context, inputs <-chan Input, f func(ctx context.Context, input Input)) {
	brad.SinkFunc(ctx, inputs, f)
}

// Uses the given function to transform a channel of generic Inputs to a channel of generic Outputs.
func Transformer[Input, Output any](ctx context.Context, input <-chan Input, f func(ctx context.Context, input Input) Output) <-chan Output {
	return brad.TransformerFunc(ctx, input, f)
}

// PeekAndReceive checks whether the channel has a value and receives the value if possible:
//   - if 'ch' has a value, PeekAndReceive receives the value and returns the value and 'true', 'true';
//   - if 'ch' is open and has no value or the channel is nil, PeekAndReceive returns a [zero value] and 'false', 'true';
//   - if 'ch' is closed and empty, PeekAndReceive returns a [zero value] and 'false', 'false'.
//
// [zero value]: https://go.dev/ref/spec#The_zero_value
func PeekAndReceive[T any](ch <-chan T) (value T, ok, open bool) {
	return sol.PeekAndReceive(ch)
}
