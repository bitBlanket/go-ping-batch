package main

import (
	"fmt"
	"os"

	"github.com/bitBlanket/go-ping-batch/internal/config"
	"github.com/bitBlanket/go-ping-batch/internal/pinger"
)

func main() {
	// 1. 解析命令行参数
	cfg := config.ParseFlags()

	// 2. 打印一下配置，确认解析没问题
	fmt.Printf("Target: %s, Count: %d, Interval: %v, Timeout: %v\n",
		cfg.Target, cfg.Count, cfg.Interval, cfg.Timeout)

	// 3. 调用 Pinger
	p := pinger.NewPinger(cfg)
	if err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
