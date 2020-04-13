package models

import (
	"encoding/json"
	"github.com/Cactush/go-gin/pkg/logging"
)

type Option struct {
	Answer  string   `json:"answer"`
	KeyWord []string `json:"key_word"`
}
type Questionbank struct {
	ID         int    `json:"id"`
	LovetypeId int    `json:"-"`
	Topic      string `json:"topic"`
	Option_1   string `json:"-"`
	Option_2   string `json:"-"`
	Option_3   string `json:"-"`
	Option_4   string `json:"-"`
}

func (q *Questionbank) GetOption() (map[int]string, error) {
	var err error
	var temp_option_list []Option
	str := [][]byte{[]byte(q.Option_1), []byte(q.Option_2), []byte(q.Option_3), []byte(q.Option_4)}
	for _, value := range str {
		option := Option{}
		err = json.Unmarshal(value, &option)
		temp_option_list = append(temp_option_list, option)
	}
	var options_dict map[int]string = make(map[int]string)
	for index, value := range temp_option_list {
		logging.Info(index, value)
		options_dict[index+1] = value.Answer
	}
	return options_dict, err
}
