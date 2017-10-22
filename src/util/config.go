package util

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Entry, Dest, Format, Chunk string
	Cleanup, Log               bool
}

func (c *Config) GetConfig(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	datajson := []byte(data)
	err = json.Unmarshal(datajson, c)
	if err != nil {
		return
	}
	println(c)
}
