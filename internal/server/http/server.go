package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	URLStubPattern       = "/path/{intStubParam:[0-9]+}/{stringStubParam:[a-zA-Z-]+}/{anyStubParam:.+}"
	IntStubParamField    = "intStubParam"
	StringStubParamField = "stringStubParam"
	AnyStubParamField    = "anyStubParam"
)

type Config interface {
	GetHTTPHost() string
	GetHTTPPort() string
}

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type Application interface {
	StubMethod(intStubParam int, stringStubParam string, anyStubParam string, headers map[string][]string) ([]byte, error)
}

type Server struct {
	Logger Logger
	Server *http.Server
}

var (
	ErrParameterParseIntStubParam = errors.New("unable to parse int stub param")
	ErrStubMethodExec             = errors.New("unable to execute stub method")
	ErrResponseWrite              = errors.New("unable to write a response")
)

type Handler struct {
	App    Application
	Logger Logger
}

func New(config Config, logger Logger, app Application) *Server {
	handler := &Handler{
		App:    app,
		Logger: logger,
	}

	router := mux.NewRouter()
	router.HandleFunc(URLStubPattern, handler.stubHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    net.JoinHostPort(config.GetHTTPHost(), config.GetHTTPPort()),
		Handler: router,
	}

	return &Server{
		Logger: logger,
		Server: server,
	}
}

func (h *Handler) stubHandler(w http.ResponseWriter, r *http.Request) {
	intParam, err := strconv.Atoi(mux.Vars(r)[IntStubParamField])
	if err != nil {
		SendBadGatewayStatus(w, h, fmt.Errorf("%w: %s", ErrParameterParseIntStubParam, err))
		return
	}

	bytes, err := h.App.StubMethod(intParam, mux.Vars(r)[StringStubParamField], mux.Vars(r)[AnyStubParamField], r.Header)
	if err != nil {
		SendBadGatewayStatus(w, h, err)
		return
	}

	w.Header().Set("Content-Type", http.DetectContentType(bytes))
	w.Header().Set("Content-Length", strconv.Itoa(len(bytes)))
	if _, err := w.Write(bytes); err != nil {
		h.Logger.Error(fmt.Errorf("%w: %s", ErrStubMethodExec, err.Error()))
	}
}

func SendBadGatewayStatus(w http.ResponseWriter, h *Handler, err error) {
	w.WriteHeader(http.StatusBadGateway)
	if n, e := w.Write([]byte(err.Error())); e != nil {
		h.Logger.Error(fmt.Errorf("%w: trying to write %d bytes: %s", ErrResponseWrite, n, e.Error()))
	}
	h.Logger.Error(err.Error())
}

func (s *Server) Start() error {
	return s.Server.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
