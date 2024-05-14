package hyperproof

import (
	"github.com/sirupsen/logrus"
)

const (
	Name    = "hyperproof"
	Usage   = "hyperproof description"
	EnvName = "HYPERPROOF"
)

var (
	Logger *logrus.Logger = logrus.New()
)
