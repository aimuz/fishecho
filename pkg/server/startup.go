package server

import (
	"context"

	"github.com/aimuz/fishecho/pkg/register"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Start(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg, ctx := errgroup.WithContext(ctx)

	services := append(register.GetIniter(), register.GetServer()...)
	for _, svc := range services {
		func(svc register.Service) {
			wg.Go(func() error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					if err := svc.Run(ctx); err != nil {
						cancel()
						return err
					}
					logrus.Infoln("Stopped", svc.Name())
					return nil
				}
			})
		}(svc)
	}
	return wg.Wait()
}
