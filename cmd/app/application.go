package app

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/KRR19/number-search/internal/domain/numbersearch"
	"github.com/KRR19/number-search/internal/infra/config"
	"github.com/KRR19/number-search/internal/infra/filestore"
	"github.com/KRR19/number-search/internal/infra/rest"
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
	if err := store.ReadFromFile(cfg.FilePath()); err != nil {
		panic(err)
	}

	ns := numbersearch.NewService(logger, store, cfg)

	handler := rest.NewHandler(logger, ns)

	mux := createHandlerMux(handler)

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

func createHandlerMux(handler *rest.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(rest.V1GetNumberPositionPath, handler.GetNumberPosition)
	mux.HandleFunc(rest.V2GetNumberPositionPath, handler.V2GetNumberPosition)
	return mux
}

func newConfig() *config.Config {
	vp := viper.New()
	var filePath string
	fileName := ".env"
	if _, err := os.Stat("../" + fileName); !os.IsNotExist(err) {
		filePath = "../" + fileName
	} else if _, err := os.Stat("./" + fileName); !os.IsNotExist(err) {
		filePath = fileName
	}

	if filePath != "" {
		vp.SetConfigFile(filePath)
		_ = vp.ReadInConfig()
	} else {
		vp.SetDefault("PORT", os.Getenv("PORT"))
		vp.SetDefault("LOG_LEVEL", os.Getenv("LOG_LEVEL"))
		vp.SetDefault("VARIATION", os.Getenv("VARIATION"))
		vp.SetDefault("FILE_PATH", os.Getenv("FILE_PATH"))
	}

	cfg := config.NewConfig(vp)
	return cfg
}
