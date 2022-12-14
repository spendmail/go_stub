package app

import (
	"errors"
	"fmt"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

type Config interface {
}

type Application struct {
	Logger Logger
	Config Config
}

var (
	ErrRequest = errors.New("request error")
)

func New(logger Logger, config Config) (*Application, error) {
	return &Application{
		Logger: logger,
		Config: config,
	}, nil
}

func (app *Application) StubMethod(intStubParam int, stringStubParam string, anyStubParam string, headers map[string][]string) ([]byte, error) {

	resultBytes := []byte(fmt.Sprintf("%s, %s!", stringStubParam, anyStubParam))
	return resultBytes, nil
}
