package umongo

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type clientOpts struct {
	requireSuccessfulPing    *context.Context
	additionalClientOpts     []*options.ClientOptions
	registerEmbeddedDocument bool
}

type applyOption func(*clientOpts)

func WithRequireSuccessfulPing(ctx context.Context) applyOption {
	return func(o *clientOpts) {
		o.requireSuccessfulPing = &ctx
	}
}

func WithRegisterEmbeddedDocument(enable bool) applyOption {
	return func(o *clientOpts) {
		o.registerEmbeddedDocument = enable
	}
}

func WithClientOption(opts *options.ClientOptions) applyOption {
	return func(o *clientOpts) {
		o.additionalClientOpts = append(o.additionalClientOpts, opts)
	}
}
