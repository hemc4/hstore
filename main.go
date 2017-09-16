package hstore

import (
	"fmt"
	"log"
	"time"
)

const (
	folder  string = "/Users/hemc/exchdata/"
	fileExt string = "txt"
)

type HStore struct {
	Name string
}

func (hStore *HStore) StoreIt(line string) {
	storeFile := hStore.getFilePath()
	addString(storeFile, line)
}

func (hStore *HStore) getFilePath() string {
	var fileFullName = hStore.Name + getHourlyName()
	log.Println("file :" + fileFullName)
	return folder + fileFullName + "." + fileExt
}

func getHourlyName() string {
	t := time.Now()
	appendDate := fmt.Sprintf("_%d_%02d_%02d", t.Year(), t.Month(), t.Day())
	appendHour := fmt.Sprintf("_%02d", t.Day())
	return appendDate + appendHour
}
