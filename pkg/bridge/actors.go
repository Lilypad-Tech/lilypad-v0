package bridge

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
)

func Actor[EIn any, EOut any](
	ctx, actionCtx context.Context,
	in <-chan EIn,
	action func(context.Context, EIn) EOut,
	out chan<- EOut,
) {
	log.Ctx(ctx).Debug().Msg("Started")
	defer log.Ctx(ctx).Debug().Msg("Stopped")

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
	log.Ctx(ctx).Debug().Msg("Started")
	defer log.Ctx(ctx).Debug().Msg("Stopped")

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
	log.Ctx(ctx).Debug().Msg("Started")
	defer log.Ctx(ctx).Debug().Msg("Stopped")

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

func Persist[E Event](ctx context.Context, in <-chan E, out chan<- E) {
	ctx = log.Ctx(ctx).With().Str("action", "Persist").Logger().WithContext(ctx)
	Actor(ctx, ctx, in, func(ctx context.Context, in E) E {
		// Persist...
		log.Ctx(ctx).Debug().Stringer("id", in.OrderId()).Str("state", fmt.Sprintf("%T", in)).Msg("Saved")
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
