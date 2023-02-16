package batch

import "github.com/hidenari-yuda/go-grpc-clean/domain/config"

type Batch struct {
	cfg config.Config
}

func NewBatch(cfg config.Config) *Batch {
	return &Batch{cfg: cfg}
}

func (b *Batch) Start() {}
