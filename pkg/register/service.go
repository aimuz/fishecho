package register

import "context"

type RunFn func(ctx context.Context) error

type Service struct {
	name string
	run  RunFn
}

func (s *Service) Name() string {
	return s.name
}

func (s *Service) Run(ctx context.Context) error {
	return s.run(ctx)
}

type runnable interface {
	Run(ctx context.Context) error
}
