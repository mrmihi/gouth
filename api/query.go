package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gouth/errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type Resp interface{ Response }

type Query[R Resp] struct {
	cli    Client
	Method string
	Path   string
	Body   []byte
	Form   url.Values
}

func NewQuery[R Resp](c Client) Query[R] {
	return Query[R]{
		cli:    c,
		Method: http.MethodGet,
		Form:   url.Values{},
	}
}

func (q Query[R]) WithMethod(m string) Query[R] { q.Method = m; return q }
func (q Query[R]) WithPath(format string, v ...any) Query[R] {
	q.Path = fmt.Sprintf(format, v...)
	return q
}
func (q Query[R]) WithBody(body any) Query[R]            { b, _ := json.Marshal(body); q.Body = b; return q }
func (q Query[R]) WithFormData(key, val string) Query[R] { q.Form.Set(key, val); return q }

func (q Query[R]) Do(ctx context.Context) (R, error) {
	var zero R

	fullURL, err := url.JoinPath(q.cli.getBaseURL(), q.Path)
	if err != nil {
		slog.Log(ctx, slog.LevelError, "invalid URL",
			"base", q.cli.getBaseURL(), "path", q.Path, "err", err)
		return zero, err
	}

	var body io.Reader
	switch {
	case len(q.Form) > 0:
		body = bytes.NewBufferString(q.Form.Encode())
	case len(q.Body) > 0:
		body = bytes.NewBuffer(q.Body)
	}

	slog.Log(ctx, slog.LevelInfo, "sending HTTP request",
		"method", q.Method, "url", fullURL)

	rsp, err := q.cli.do(ctx, q.Method, fullURL, body)
	if err != nil {
		slog.Log(ctx, slog.LevelError, "request failed", "err", err)
		return zero, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(rsp.Body)

	slog.Log(ctx, slog.LevelDebug, "received HTTP response",
		"method", q.Method, "url", fullURL, "status", rsp.StatusCode)

	switch rsp.StatusCode {
	case http.StatusOK, http.StatusCreated:
	case http.StatusUnauthorized, http.StatusForbidden:
		return zero, errors.ErrInvalidAuth
	case http.StatusNotFound:
		return zero, errors.ErrNotFound
	case http.StatusTooManyRequests:
		return zero, errors.ErrTooManyRequests
	default:
		return zero, errors.ErrInternalAPICall
	}

	if err := json.NewDecoder(rsp.Body).Decode(&zero); err != nil && err != io.EOF {
		buf, _ := io.ReadAll(rsp.Body)
		wrapped := fmt.Errorf("decode response: %w – body: %s", err, buf)
		slog.Log(ctx, slog.LevelError, "JSON decode error", "err", wrapped)
		return zero, wrapped
	}

	return zero, nil
}
