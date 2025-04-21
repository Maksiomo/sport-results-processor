package server

import (
	"context"
)

type Runnable interface {
	Start()
	Stop(context.Context)
}
