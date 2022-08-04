package main

import (
	"flag"
	"log"
	"os"
	"rockt/repo"
	"rockt/repo/sqlite"
	"rockt/router"
	"rockt/task"
)

func main() {
	var (
		datadir string
		db      repo.Repository
		err     error
	)

	if db, err = sqlite.NewRepository(); err != nil {
		log.Println("Error initializing sqlite database")
		os.Exit(1)
	}

	db.Create()

	flag.StringVar(&datadir, "d", "", "Data directory")
	flag.Parse()

	if datadir == "" {
		log.Println("Pass in the data directory location")
		os.Exit(1)
	}

	// check whether the directory exist
	if _, err := os.Stat(datadir); os.IsNotExist(err) {
		log.Println("Directory does not exist")
		os.Exit(1)
	}

	go task.RunTask(datadir, db)
	g := router.SetupRouter(datadir, db)
	g.Run(":8279")
}
