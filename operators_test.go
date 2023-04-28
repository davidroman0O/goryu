package goryu

import (
	"context"
	"fmt"
	"testing"
)

// go test . -v -count=1 -run $TestSimple$
func TestSimple(t *testing.T) {
	ctx := context.Background()

	agg := Aggregate[int, int](
		ctx,
		Emit[int](ctx, 1, 2, 3, 4, 5, 6),
		2,
		func(ctx context.Context, input ...int) int {
			acc := 0
			for _, v := range input {
				acc += v
			}
			return acc
		})

	arr := Array(ctx, agg)

	if arr[0] != 3 {
		t.Error("should be 3")
	}

	if arr[1] != 7 {
		t.Error("should be 7")
	}

	if arr[2] != 11 {
		t.Error("should be 11")
	}
}

// go test . -v -count=1 -run $TestSimpleSplitStream$
func TestSimpleSplitStream(t *testing.T) {
	ctx := context.Background()

	input := Emit[int](ctx, 1, 2, 3, 4, 5, 6)

	// We have now two goroutines, therefore we need two listeners
	// otherwise the compiler won't like it
	s1, s2 := Duplicate(ctx, input)

	ms1 := Map[int, int](s1, func(a int) int {
		return -1 * a
	})

	emitchan := Emit[int](
		ctx,
		Array(ctx,
			Merge[int](ctx, ms1, s2))...)

	closed := false
	for {
		if closed {
			break
		}
		select {
		case v, ok := <-emitchan:
			if !ok {
				closed = true
				break
			}
			fmt.Println(v)
		}
	}
}
