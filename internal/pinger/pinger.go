package pinger

import (
	"fmt"

	"github.com/bitBlanket/go-ping-batch/internal/config"
)

type Pinger struct {
	cfg *config.Config
}

func NewPinger(cfg *config.Config) *Pinger {
	return &Pinger{cfg: cfg}
}

func (p *Pinger) Run() error {
	// TODO: 这里未来实现 ICMP 报文
	fmt.Printf("Pinging %s with count=%d ...\n", p.cfg.Target, p.cfg.Count)
	return nil
}
