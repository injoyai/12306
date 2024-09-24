package logic

import (
	"testing"
	"time"
)

func TestGetTicketList(t *testing.T) {
	/*
		WEI 成都东
		EAY 西安北
		HGH 杭州东
		VRH 温州南
	*/
	list, err := GetTicketList(time.Date(2024, 5, 17, 0, 0, 0, 0, time.Local), "HGH", "VRH")
	if err != nil {
		t.Error(err)
		return
	}
	for _, v := range list {
		t.Log(v)
	}
}
