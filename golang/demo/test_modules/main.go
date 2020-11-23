package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("this is a test")
	logrus.Error("this is an error")
	fmt.Println("modules demo")
}
