package marusia

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"
)

// Config ...
type Config struct {
	useSSL            bool
	certFile, keyFile string
	addr              string
	webhookURL        string
}

// NewConfig ...
func NewConfig(useSSL bool, certFile, keyFile, addr, webhookURL string) *Config {
	return &Config{
		useSSL:     useSSL,
		certFile:   certFile,
		keyFile:    keyFile,
		addr:       addr,
		webhookURL: webhookURL,
	}
}

// Skill ...
type Skill struct {
	config       *Config
	dialogRouter *DialogRouter
	logger       *slog.Logger
}

// NewSkill ...
func NewSkill(c *Config, dr *DialogRouter) *Skill {
	return &Skill{
		config:       c,
		dialogRouter: dr,
		logger:       slog.Default(),
	}
}

func (s *Skill) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("new request",
		"method", r.Method,
		"url", r.URL.String(),
		"host", r.Host,
		"user_agent", r.UserAgent(),
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		s.logger.Error("fault decode response body", slog.String("error", err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req Request
	if err := json.Unmarshal(body, &req); err != nil {
		s.logger.Error("fault parsing response to struct", slog.String("error", err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	call := strings.ToLower(req.OriginalUtterance())
	df, err := s.dialogRouter.Select(call)
	if err != nil {
		s.logger.Error("fault get dialog function", slog.String("error", err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var resp Response
	resp.LoadSession(&req)

	df(context.Background(), &req, &resp)

	data, err := json.Marshal(resp)
	if err != nil {
		s.logger.Error("fault create json response err", slog.String("error", err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(data); err != nil {
		s.logger.Error("fault write response err", slog.String("error", err.Error()))
		log.Printf("fault sending client response err: %s", err)
	}

}

// ListenAndServe ...
func (s *Skill) ListenAndServe() error {
	s.logger.Info("starting server", slog.Any("config", s.config))

	http.HandleFunc(s.config.webhookURL, corsMiddleware(s.ServeHTTP))

	if !s.config.useSSL {
		return http.ListenAndServe(s.config.addr, nil)
	}

	return http.ListenAndServeTLS(s.config.addr, s.config.certFile, s.config.keyFile, nil)
}
