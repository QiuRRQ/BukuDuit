package helper

import (
	"bukuduit-go/helpers/messages"
	"errors"
)

// SmsChannel ....
type SmsChannel struct {
	ID          int    `db:"id"`
	Channel     string `db:"channel"`
	Description string `db:"description"`
}

var (
	// SmsChannelList ...
	SmsChannelList = []SmsChannel{
		{
			ID:          1,
			Channel:     "1",
			Description: "",
		},
		{
			ID:          2,
			Channel:     "2",
			Description: "",
		},
	}
)

// GetSmsChannelLength ...
func GetSmsChannelLength() int {
	return len(SmsChannelList)
}

// GetSmsChannelNextID ...
func GetSmsChannelNextID(id int) (res int) {
	res = id + 1
	if res > GetSmsChannelLength() {
		res = 1
	}

	return res
}

// GetSmsChannelByID ...
func GetSmsChannelByID(id int) (res SmsChannel, err error) {
	for _, r := range SmsChannelList {
		if r.ID == id {
			return r, err
		}
	}

	return res, errors.New(messages.DataNotFound)
}

// GetNextSmsChannelByID ...
func GetNextSmsChannelByID(id int) (res SmsChannel, err error) {
	nextID := GetSmsChannelNextID(id)

	return GetSmsChannelByID(nextID)
}
