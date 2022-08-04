package task

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"rockt/model"
	"rockt/repo"
	"strings"
	"time"
)

//RunTask execute file processing
func RunTask(directory string, db repo.Repository) {
	items, _ := ioutil.ReadDir(directory)
	for _, item := range items {
		if !item.IsDir() {
			log.Println("Processing : ", item.Name())
			dirFile := directory + "/" + item.Name()

			if _, err := os.Stat(dirFile); os.IsNotExist(err) {
				log.Println(" data file not found")
				return
			}

			var modelArray []model.Datarecord

			readFile(dirFile, func(s string) {
				split := strings.Split(s, " ")

				if len(split) == 3 {
					// check date format, if not correct skip the data
					_, err := time.Parse(time.RFC3339, split[0])
					if err == nil {
						modelArray = append(modelArray, model.Datarecord{
							DateISO8601:  split[0],
							EmailAddress: split[1],
							SessionID:    split[2],
							FileName:     item.Name(),
						})
					}
				}
			})

			db.BulkInsert(modelArray)
		}
	}
}

//readFile to read file line by line and call handle function for each line
func readFile(filePath string, handle func(string)) error {
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
