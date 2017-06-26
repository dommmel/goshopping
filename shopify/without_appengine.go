// +build !appengine

// This file provides glue for making goshopping work without App Engine.

package shopify

import (
	"context"
	"net/http"
)

func withContext(ctx context.Context, req *http.Request) (context.Context, *http.Request) {
	return ctx, req.WithContext(ctx)
}
