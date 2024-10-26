package app

import (
	"log/slog"
	"net/http"
	"os"

	v1 "github.com/KRR19/number-search/internal/infra/http/v1"
)

type CustomHandler struct{}

func (h *CustomHandler) Enabled() bool                    { return true }
func (h *CustomHandler) Log(level slog.Level, msg string) {}

type Application struct {
	mux    *http.ServeMux
	logger *slog.Logger
}

func NewApplication() *Application {
	loger := newLogger()

	v1Handler := v1.NewHandler()

	mux := createHandlerMux(v1Handler)

	return &Application{
		logger: loger,
		mux:    mux,
	}
}

func (a *Application) ServeHTTP() error {
	port := ":8080"
	a.logger.Info("starting http server on port " + port)

	return http.ListenAndServe(port, a.mux)
}

func newLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	slog.SetDefault(logger)

	return logger
}

func createHandlerMux(v1Handler *v1.Handler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(v1.GetNumberPositionPath, v1Handler.GetNumberPosition)
	return mux
}
