package batch

import "github.com/hidenari-yuda/go-docker-template/domain/config"

type Batch struct {
	cfg config.Config
}

func NewBatch(cfg config.Config) *Batch {
	return &Batch{cfg: cfg}
}

func (b *Batch) Start() {}
