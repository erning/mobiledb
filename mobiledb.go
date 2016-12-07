package mobiledb

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strconv"
)

type MobileInfo struct {
	Province   string
	City       string
	ISP        string
	Code       int
	ZipCode    int
	DetailType string
}

type MobileDB struct {
	Prefixes map[int]int
	Indexes  []int
	Values   [][6]int
	Symbols  []string
}

var db MobileDB

func init() {
	b64, _ := base64.StdEncoding.DecodeString(DATA_B64)
	decoder := gob.NewDecoder(bytes.NewReader(b64))
	decoder.Decode(&db)
}

func GetMobileInfo(number string) (*MobileInfo, error) {
	prefix, err := strconv.Atoi(number[:3])
	if err != nil {
		return nil, fmt.Errorf("Unknown prefix")
	}
	index, ok := db.Prefixes[prefix]
	if !ok {
		return nil, fmt.Errorf("Unknown prefix")
	}
	offset, err := strconv.Atoi(number[3:7])
	if err != nil {
		return nil, fmt.Errorf("Unknown number")
	}

	value := db.Values[db.Indexes[index+offset]]
	if value[3] == 0 {
		return nil, fmt.Errorf("Unknown number")
	}

	return &MobileInfo{
		Province:   db.Symbols[value[0]],
		City:       db.Symbols[value[1]],
		ISP:        db.Symbols[value[2]],
		Code:       value[3],
		ZipCode:    value[4],
		DetailType: db.Symbols[value[5]],
	}, nil

}
