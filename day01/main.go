package main

import (
	"context"
	"strconv"

	"github.com/hienduyph/goss/iox"
	"github.com/hienduyph/goss/logger"
	"github.com/hienduyph/goss/streamx"
)

//

func main() {
	ctx := context.Background()
	r, e := iox.Read(ctx, "day01/input.txt")
	logger.FatalIf(e, "read file failed")
	defer r.Close()
	stream, e := streamx.ReadLines(ctx, r)
	logger.FatalIf(e, "construct stream failed")
	prev := 0
	idx := 0
	result := 0
	for item := range stream.Stream {
		val, e := strconv.Atoi(string(item))
		logger.FatalIf(e, "construct stream failed")
		if val >= prev && idx > 0 {
			result += 1
		}
		idx++
		prev = val
	}
	logger.Info("larger than prev", "Result", result)
}
