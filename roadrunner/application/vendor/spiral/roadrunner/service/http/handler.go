package http

import (
	"github.com/pkg/errors"
	"github.com/spiral/roadrunner"
	"net/http"
	"strconv"
	"sync"
)

const (
	// EventResponse thrown after the request been processed. See ErrorEvent as payload.
	EventResponse = iota + 500

	// EventError thrown on any non job error provided by road runner server.
	EventError
)

// ErrorEvent represents singular http error event.
type ErrorEvent struct {
	// Request contains client request, must not be stored.
	Request *http.Request

	// Error - associated error, if any.
	Error error
}

// ResponseEvent represents singular http response event.
type ResponseEvent struct {
	// Request contains client request, must not be stored.
	Request *Request

	// Response contains service response.
	Response *Response
}

// Handler serves http connections to underlying PHP application using PSR-7 protocol. Context will include request headers,
// parsed files and query, payload will include parsed form dataTree (if any).
type Handler struct {
	cfg *Config
	rr  *roadrunner.Server
	mul sync.Mutex
	lsn func(event int, ctx interface{})
}

// Listen attaches handler event watcher.
func (h *Handler) Listen(l func(event int, ctx interface{})) {
	h.mul.Lock()
	defer h.mul.Unlock()

	h.lsn = l
}

// mdwr serve using PSR-7 requests passed to underlying application. Attempts to serve static files first if enabled.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// validating request size
	if h.cfg.MaxRequest != 0 {
		if length := r.Header.Get("content-length"); length != "" {
			if size, err := strconv.ParseInt(length, 10, 64); err != nil {
				h.handleError(w, r, err)
				return
			} else if size > h.cfg.MaxRequest*1024*1024 {
				h.handleError(w, r, errors.New("request body max size is exceeded"))
				return
			}
		}
	}

	req, err := NewRequest(r, h.cfg.Uploads)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	req.Open()
	defer req.Close()

	p, err := req.Payload()
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	rsp, err := h.rr.Exec(p)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	resp, err := NewResponse(rsp)
	if err != nil {
		h.handleError(w, r, err)
		return
	}

	h.handleResponse(req, resp)
	resp.Write(w)
}

// handleError sends error.
func (h *Handler) handleError(w http.ResponseWriter, r *http.Request, err error) {
	h.throw(EventError, &ErrorEvent{Request: r, Error: err})

	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

// handleResponse triggers response event.
func (h *Handler) handleResponse(req *Request, resp *Response) {
	h.throw(EventResponse, &ResponseEvent{Request: req, Response: resp})
}

// throw invokes event handler if any.
func (h *Handler) throw(event int, ctx interface{}) {
	h.mul.Lock()
	defer h.mul.Unlock()

	if h.lsn != nil {
		h.lsn(event, ctx)
	}
}
