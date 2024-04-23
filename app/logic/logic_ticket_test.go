package logic

import (
	"testing"
	"time"
)

func TestGetTicketList(t *testing.T) {
	list, err := GetTicketList(time.Date(2024, 5, 1, 0, 0, 0, 0, time.Local), "WEI", "EAY")
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range list {
		t.Log(v)
	}
}
