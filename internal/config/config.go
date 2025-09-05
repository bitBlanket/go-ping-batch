package config

import (
	"flag"
	"time"
)

type Config struct {
	Count    int           // ping 次数
	Interval time.Duration // 每次间隔
	Timeout  time.Duration // 超时时间
	Target   string        // 目标地址
}

// ParseFlags 解析命令行参数
func ParseFlags() *Config {
	cfg := &Config{}

	flag.IntVar(&cfg.Count, "c", 4, "Number of echo requests to send")
	flag.DurationVar(&cfg.Interval, "i", time.Second, "Interval between requests")
	flag.DurationVar(&cfg.Timeout, "w", time.Second*3, "Timeout for each request")

	flag.Parse()

	// 最后一个参数作为目标地址
	if flag.NArg() < 1 {
		flag.Usage()
	}
	cfg.Target = flag.Arg(0)

	return cfg
}
