package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/BlaineEXE/go-cobra-with-interrupt/cmd/myapp"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	myapp.ExecuteContext(ctx)
}
