package marusia

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
}

// NewSkill ...
func NewSkill(c *Config, dr *DialogRouter) *Skill {
	return &Skill{
		config:       c,
		dialogRouter: dr,
	}
}

func (s *Skill) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf(`%s %s "%s %s" "%s"`, r.Host, r.RemoteAddr, r.Method, r.URL, r.UserAgent())
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("fault decode response body err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var cr Request
	if err := json.Unmarshal(body, &cr); err != nil {
		log.Printf("fault parsing response to struct err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	call := strings.ToLower(cr.OriginalUtterance())
	df, err := s.dialogRouter.Select(call)
	if err != nil {
		log.Printf("fault get dialog function err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var sr Response
	sr.LoadSession(&cr)
	
	resp := df(&cr)

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("fault create json response err: %s", err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, err := w.Write(data); err != nil {
		log.Printf("fault sending client response err: %s", err)
	}

}

// ListenAndServe ...
func (s *Skill) ListenAndServe() error {
	log.Printf("starting server config: %+v ...", s.config)
	http.HandleFunc(s.config.webhookURL, corsMiddleware(s.ServeHTTP))
	if !s.config.useSSL {
		return http.ListenAndServe(s.config.addr, nil)
	}
	return http.ListenAndServeTLS(s.config.addr, s.config.certFile, s.config.keyFile, nil)
}
