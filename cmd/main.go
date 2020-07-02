package main

import (
	"github.com/robertpeteuil/minesweeper/api"
	"github.com/sirupsen/logrus"
)

func main() {
	//log := &logger{logrus.New()}
	log := logrus.StandardLogger()
	log.Infoln("Starting API...")
	if err := api.Start(log); err != nil {
		log.WithError(err).Fatalln("Couldn't start API.")
	}
}
