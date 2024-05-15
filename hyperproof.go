package hyperproof

import (
	"github.com/sirupsen/logrus"
)

const (
	Name    = "hyperproof"
	Usage   = "hyperproof platform client"
	EnvName = "HYPERPROOF"
)

var (
	Logger *logrus.Logger = logrus.New()
)
