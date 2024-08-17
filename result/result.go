// Package result provides concurrency utilities for accumulating results and managing errors.
//
// Example
//
//	import (
//		result "github.com/neurocode-io/go-pkgs/result"
//		"context"
//	)
//
//	func worker() ([]int, error) {
//		// worker logic here
//		return []int{1, 2, 3}, nil
//	}
//
//	ctx := context.Background()
//	group, ctx := result.WithErrorsThreshold[int](ctx, 2)
//	group.Go(worker)
//	results, err := group.Wait()
//	if err != nil {
//		handle error
//	}
//	results == []int{1, 2, 3}
package result

import (
	"context"
	"sync"
)

type errorWithUnwrap interface {
	error
	Unwrap() []error
}

type multiError struct {
	errs []error
}

type Group[T any] struct {
	cancel    func()
	results   []T
	errs      []error
	wg        sync.WaitGroup
	mutex     sync.Mutex
	threshold int
}

/*
WithErrorsThreshold initializes a new Group[T] with a threshold for error tolerance.

Example

	ctx := context.Background()
	group, ctx := result.WithErrorsThreshold[int](ctx, 2)
	// group is now ready to execute with a threshold of 2 errors
*/
func WithErrorsThreshold[T any](ctx context.Context, threshold int) (Group[T], context.Context) {
	if threshold < 1 {
		panic("threshold must be greater than or equal to 1")
	}

	ctx, cancel := context.WithCancel(ctx)

	return Group[T]{cancel: cancel, threshold: threshold}, ctx
}

/*
Go starts a goroutine that performs a given function and handles its results and errors.

Example

	group.Go(func() ([]int, error) {
		return []int{1, 2, 3}, nil
	})
	// The results will be accumulated and errors managed based on the Group's settings.
*/
func (g *Group[T]) Go(f func() ([]T, error)) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		res, err := f()
		g.processResult(res, err)
	}()
}

func (g *Group[T]) processResult(res []T, err error) {
	if err != nil {
		g.handleErrors(err)
	}

	g.appendResults(res)
}

func (g *Group[T]) handleErrors(err error) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.threshold == 0 || len(g.errs) < g.threshold {
		g.errs = append(g.errs, err)
	}

	if len(g.errs) == g.threshold {
		if g.cancel != nil {
			g.cancel()
		}
	}
}

func (g *Group[T]) appendResults(res []T) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.results = append(g.results, res...)
}

/*
Wait blocks until all tasks have completed and returns the accumulated results and any errors.

Example

	results, err := group.Wait()
	if err != nil {
		// handle error
	}
	// results contains all accumulated results from the group operations
*/
func (g *Group[T]) Wait() ([]T, errorWithUnwrap) {
	g.wg.Wait()
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.cancel != nil {
		g.cancel()
	}

	if len(g.errs) == 0 {
		return g.results, nil
	}

	return g.results, &multiError{errs: g.errs}
}

func (me *multiError) Error() string {
	var b []byte
	for i, err := range me.errs {
		if i > 0 {
			b = append(b, '\n')
		}
		b = append(b, err.Error()...)
	}

	return string(b)
}

func (me *multiError) Unwrap() []error {
	return me.errs
}
