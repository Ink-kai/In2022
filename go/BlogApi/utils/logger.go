package utils

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	log.SetOutput(os.Stdout)
}
