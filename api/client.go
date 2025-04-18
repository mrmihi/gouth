package api

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"time"
)

type Client interface {
	getBaseURL() string
	do(ctx context.Context, method, url string, body io.Reader) (*http.Response, error)
}

type RestClientConfig struct {
	BaseURL string        // e.g. "https://api.example.com"
	Dump    bool          // if true dump full HTTP requests / responses
	Timeout time.Duration // optional: 0 == use http.DefaultClient timeout
}

type RestClient struct {
	RestClientConfig
	httpC *http.Client
	lg    *slog.Logger
}

func NewRestClient(lg *slog.Logger, cfg RestClientConfig) *RestClient {
	if lg == nil {
		lg = slog.Default()
	}

	cl := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	if cfg.Timeout > 0 {
		cl.Timeout = cfg.Timeout
	}

	return &RestClient{
		RestClientConfig: cfg,
		httpC:            cl,
		lg:               lg,
	}
}

func (c *RestClient) do(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if c.Dump {
		if dump, e := httputil.DumpRequestOut(req, true); e == nil {
			c.lg.Log(ctx, slog.LevelDebug, "HTTP request dump", "dump", dump)
		}
	}

	rsp, err := c.httpC.Do(req)
	if err != nil {
		return nil, err
	}

	if c.Dump {
		if dump, e := httputil.DumpResponse(rsp, true); e == nil {
			c.lg.Log(ctx, slog.LevelDebug, "HTTP response dump", "dump", dump)
		}
	}

	c.lg.Log(ctx, slog.LevelDebug, "HTTP round‑trip complete",
		"method", method, "url", url, "status", rsp.StatusCode)

	return rsp, nil
}

func (c *RestClient) getBaseURL() string { return c.BaseURL }
