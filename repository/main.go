package main

import (
	"bufio"
	"flag"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"rockt/repository/model"
	"rockt/repository/repository"
	"rockt/repository/repository/sqlite"
	"strings"
)

func main() {
	var datadir string
	var db repository.Repository
	var err error

	if db, err = sqlite.NewRepository(); err != nil {
		log.Println("Error initializing sqlite inmemory database")
		os.Exit(1)
	}

	db.Create()

	flag.StringVar(&datadir, "d", "", "Data directory")
	flag.Parse()

	if datadir == "" {
		log.Println("Pass in the data directory location")
		os.Exit(1)
	}

	// check for directory existence
	if _, err := os.Stat(datadir); os.IsNotExist(err) {
		log.Println("Directory does not exist")
		os.Exit(1)
	}

	g := SetupRouter(datadir, db)
	g.Run(":8279")
}

func SetupRouter(datadir string, db repository.Repository) *gin.Engine {
	g := gin.Default()
	g.POST("/", postHandler(datadir, db))

	return g
}

func postHandler(datadir string, db repository.Repository) func(c *gin.Context) {
	return func(c *gin.Context) {
		var r model.RequestBody

		if err := c.BindJSON(&r); err != nil {
			c.JSON(400, &model.ResponseError{Message: "invalid request"})
			return
		}

		log.Println("Processing file ", r.Filename)

		dirFile := datadir + "/" + r.Filename

		if _, err := os.Stat(dirFile); os.IsNotExist(err) {
			c.JSON(400, &model.ResponseError{Message: r.Filename + " data file not found"})
			return
		}

		var modelArray []model.Datarecord

		ReadFile(dirFile, func(s string) {
			split := strings.Split(s, " ")

			if len(split) == 3 {
				modelArray = append(modelArray, model.Datarecord{
					DateISO8601:  split[0],
					EmailAddress: split[1],
					SessionID:    split[2],
				})
			}
		})

		db.BulkInsert(modelArray)

		records := db.Query(r.From, r.To)

		// query the database
		c.JSON(200, &records)
	}
}

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	buf := bufio.NewReader(f)

	for {
		line, _, err := buf.ReadLine()
		l := strings.TrimSpace(string(line))
		handle(l)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}
