package bridge

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
)

func Actor[EIn any, EOut any](
	ctx, actionCtx context.Context,
	in <-chan EIn,
	action func(context.Context, EIn) EOut,
	out chan<- EOut,
) {
	log.Ctx(ctx).Debug().Msg("Actor started")
	defer log.Ctx(ctx).Debug().Msg("Actor stopped")

	for {
		select {
		case event := <-in:
			out <- action(actionCtx, event)
		case <-ctx.Done():
			return
		case <-actionCtx.Done():
			return
		}
	}
}

func TwoActor[EIn, EOut1, EOut2 any](
	ctx, actionCtx context.Context,
	in <-chan EIn,
	action func(context.Context, EIn) (EOut1, EOut2),
	out1 chan<- EOut1,
	out2 chan<- EOut2,
) {
	log.Ctx(ctx).Debug().Msg("TwoActor started")
	defer log.Ctx(ctx).Debug().Msg("TwoActor stopped")

	for {
		select {
		case event := <-in:
			result1, result2 := action(actionCtx, event)
			out1 <- result1
			out2 <- result2
		case <-ctx.Done():
			return
		case <-actionCtx.Done():
			return
		}
	}
}

func ErrorActor[EIn any, EOut any](
	ctx, actionCtx context.Context,
	in <-chan EIn,
	action func(context.Context, EIn) (EOut, error),
	out chan<- EOut,
	errored chan<- EIn,
) {
	log.Ctx(ctx).Debug().Msg("ErrorActor started")
	defer log.Ctx(ctx).Debug().Msg("ErrorActor stopped")

	for {
		select {
		case event := <-in:
			result, err := action(actionCtx, event)
			if err != nil {
				log.Ctx(ctx).Error().Err(err).Send()
				errored <- event
			} else {
				out <- result
			}
		case <-ctx.Done():
			return
		case <-actionCtx.Done():
			return
		}
	}
}

func Flatten[E any](ctx context.Context, in <-chan []E, out chan<- E) {
	for {
		select {
		case events := <-in:
			for _, event := range events {
				out <- event
			}
		case <-ctx.Done():
			return
		}
	}
}

func Persist[E any](ctx context.Context, in <-chan E, out chan<- E) {
	Actor(ctx, ctx, in, func(ctx context.Context, in E) E {
		// Persist...
		return in
	}, out)
}

func Fetch[E any](ctx context.Context, out chan<- E) {
	log.Ctx(ctx).Debug().Msg("Fetch started")
	// Fetch...
}

func Discard[E any](ctx context.Context, in <-chan E) {
	for {
		select {
		case <-in:
			continue
		case <-ctx.Done():
			return
		}
	}
}

func Retry[E Retryable](ctx context.Context, maxAttempts int, in <-chan E, retry chan<- E, cancel chan<- E) {
	ErrorActor(ctx, ctx, in, func(ctx context.Context, e E) (E, error) {
		if e.Attempts() >= maxAttempts {
			return e, errors.New("over max retry attempts")
		} else {
			e.AddAttempt()
			return e, nil
		}
	}, retry, cancel)
}
