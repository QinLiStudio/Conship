package app

import (
	"context"
	"fmt"
	"github.com/QinLiStudio/Conship/internal/app/config"
	"github.com/QinLiStudio/Conship/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type options struct {
	ConfigFile string
}

type Option func(*options)

func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	config.MustLoad(o.ConfigFile)

	logger.WithContext(ctx).Info(fmt.Sprintf("Start server,#run_mode %s,#pid %d", config.C.RunMode, os.Getpid()))

	err := InitLogger()
	if err != nil {
		return nil, err
	}

	injector, injectorCleanFunc, err := BuildInjector()
	if err != nil {
		return nil, err
	}

	httpServerCleanFunc := InitHTTPServer(ctx, injector.Engine)

	return func() {
		injectorCleanFunc()
		httpServerCleanFunc()
	}, nil
}

func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Info(fmt.Sprintf("HTTP server is running on: %s.", addr))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Error(err.Error())
		}
	}
}

func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Info(fmt.Sprintf("Receive signal [%s]", sig.String()))
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.WithContext(ctx).Info("Server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
