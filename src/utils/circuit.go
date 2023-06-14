package utils

import (
	"context"
)

type Circuit func(context.Context) (string, error)

/* func myFunction func(ctx context.Context) (string, error) {return "", ""}
wrapped := Breaker(DebouceFirst(myFunction))
response, err := wrapped(ctx) */
