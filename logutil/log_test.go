package logutil

import (
	"testing"

	"github.com/pkg/errors"
)

func TestVerbosef(t *testing.T) {
	Level = FATAL
	Verbosef("verbose 1")
	Level = VERBOSE
	Verbosef("verbose 2")
}

func TestDebugf(t *testing.T) {
	Level = FATAL
	Debugf("debug 1")
	Level = DEBUG
	Debugf("debug 2")
}

func TestInfof(t *testing.T) {
	Level = FATAL
	Infof("info 1")
	Level = INFO
	Infof("info 2")
}

func TestWarnf(t *testing.T) {
	Level = FATAL
	Warnf("warn 1")
	Level = WARN
	Warnf("warn 2")
}

func TestWarn(t *testing.T) {
	Level = FATAL
	Warn(errors.Wrap(errors.New("warn1"), "wrap"))
	Level = WARN
	Warn(errors.Wrap(errors.New("warn2"), "wrap"))
}

func TestErrf(t *testing.T) {
	Level = FATAL
	Errf("err 1")
	Level = ERR
	Errf("err 2")
}

func TestErr(t *testing.T) {
	Level = FATAL
	Err(errors.Wrap(errors.New("err1"), "wrap"))
	Level = ERR
	Err(errors.Wrap(errors.New("err2"), "wrap"))
}
