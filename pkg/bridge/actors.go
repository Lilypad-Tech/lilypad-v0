package bridge

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// An Actor receives events on its input channel, runs the passed action for
// each received, and posts the results to its output channel.
//
// An Actor takes two contexts – one for itself and one for its action. If its
// own context is cancelled, it will allow the action to complete and then
// close, allowing a graceful shutdown. If its action's context is cancelled,
// the action is immediately ended and the actor closes.
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

// A TwoActor is like an Actor except the action it is given produces two
// values, and each are posted to a separate channel.
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

// An ErrorActor is like an Actor, except if its passed action returns an error
// it will post the input action to an error channel.
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

// Flatten is an Actor that receives a slice of values and sends each one
// individually to its output channel.
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

// Pesist is a placeholder for saving events to the database.
func Persist[E Event](ctx context.Context, in <-chan E, repo Repository, out chan<- E) {
	ctx = log.Ctx(ctx).With().Str("action", "Persist").Logger().WithContext(ctx)
	Actor(ctx, ctx, in, func(ctx context.Context, in E) E {
		err := repo.Save(in)
		level := map[bool]zerolog.Level{true: zerolog.DebugLevel, false: zerolog.ErrorLevel}[err == nil]
		log.Ctx(ctx).WithLevel(level).Stringer("id", in.OrderId()).Err(err).Msg("Saved")
		return in
	}, out)
}

// Fetch is a placeholder for retrieving events from the database.
func Fetch[E Event](ctx context.Context, repo Repository, state OrderState, out chan<- []E) {
	log.Ctx(ctx).Debug().Msg("Fetch started")

	events, err := repo.Reload(state)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Send()
	} else {
		es := make([]E, 0, len(events))
		for _, e := range events {
			if ec, ok := e.(E); ok {
				es = append(es, ec)
			} else {
				log.Ctx(ctx).Error().Msg("failed to cast event")
			}
		}
		log.Ctx(ctx).Debug().Int("count", len(es)).Msg("Fetch finished")
		out <- es
	}
}

// Discard consumes and ignores everything on its input channel.
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
