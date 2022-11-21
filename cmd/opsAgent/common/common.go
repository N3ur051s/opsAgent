package common

import (
	"context"
)

var (
	// MainCtx is the main agent context passed to components
	MainCtx context.Context

	// MainCtxCancel cancels the main agent context
	MainCtxCancel context.CancelFunc
)
