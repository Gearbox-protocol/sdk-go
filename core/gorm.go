package core

import (
	"fmt"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func NewDBClient(con string, config *gorm.Config) *gorm.DB {
	urls := func () []string {
		s:= strings.Split(con, "@")
		front := s[0]
		s = strings.Split(s[1], "/")
		mid, end := s[0], s[1]
		urls := []string{}
		for _, host := range strings.Split(mid, ",") {
			url := fmt.Sprintf("%s@%s/%s", front, host, end)
			urls = append(urls , url)
		}
		return urls
	}()
	db, err := selectDB(urls, config)
	log.CheckFatal(err)
	return db
}

func newDB(url string, cof *gorm.Config) (*gorm.DB, error){
	return gorm.Open(postgres.Open(url), cof)
}

func selectDB(urls []string,  cof *gorm.Config) (db *gorm.DB, err error) {
	data := &struct {
		RecoveryMode bool `gorm:"column:pg_is_in_recovery"`
	}{}
	for ind:=0;  ind< len(urls); ind++ { // for each url, try 2 times
		for i:=0;i< 2 ; i++ {
			db, err = newDB(urls[ind], cof)
			if !isServerDown(err)  {
				err = db.Raw("SELECT pg_is_in_recovery();").First(data).Error
				if err != nil || data.RecoveryMode {
					log.Info("Recovery mode", ind)
					break
				}
				log.Info("using gorm db ind", ind)
				return db, err
			}
		}
	}
	return 
}

func isServerDown(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "i/o timeout") ||strings.Contains(err.Error(), "timeout: context deadline exceeded") )
}