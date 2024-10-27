package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/KRR19/number-search/internal/domain/numbersearch"
	"github.com/KRR19/number-search/internal/infra/config"
	"github.com/KRR19/number-search/internal/infra/filestore"
	v1 "github.com/KRR19/number-search/internal/infra/http/v1"
	"github.com/spf13/viper"
)

type Application struct {
	mux    *http.ServeMux
	logger *slog.Logger
	cfg    *config.Config
}

func NewApplication() *Application {
	cfg := newConfig()

	logger := newLogger(cfg.LogLevel())

	store := filestore.NewStore()
	if err := store.ReadFromFile("../input.txt"); err != nil {
		panic(err)
	}

	ns := numbersearch.NewService(logger, store, cfg)

	v1Handler := v1.NewHandler(logger, ns)

	mux := createHandlerMux(v1Handler)

	return &Application{
		logger: logger,
		mux:    mux,
		cfg:    cfg,
	}
}

func (a *Application) ServeHTTP() error {
	port := a.cfg.Port()
	a.logger.Info("starting http server on port" + port)

	return http.ListenAndServe(port, a.mux)
}

func newLogger(logLevel string) *slog.Logger {
	var level slog.Level
	switch logLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warning":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)

	return logger
}

func createHandlerMux(v1Handler *v1.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(v1.GetNumberPositionPath, v1Handler.GetNumberPosition)
	return mux
}

func newConfig() *config.Config {
	vp := viper.New()
	vp.SetConfigFile("../.env")
	_ = vp.ReadInConfig()

	vp.SetDefault("PORT", os.Getenv("PORT"))
	vp.SetDefault("LOG_LEVEL", os.Getenv("LOG_LEVEL"))
	vp.SetDefault("VARIATION", os.Getenv("VARIATION"))

	cfg := config.NewConfig(vp)
	return cfg
}
