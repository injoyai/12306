package logic

import (
	"fmt"
	"github.com/injoyai/12306/app/model"
	"github.com/injoyai/goutil/net/http"
	"time"
)

// GetTicketList 查询车票信息
func GetTicketList(date time.Time, from, to string) ([]*model.Ticket, error) {

	//请求地址
	u := "https://kyfw.12306.cn/otn/leftTicket/query"
	//query的参数有顺序要求,否则会请求失败
	u += fmt.Sprintf("?leftTicketDTO.train_date=%s&leftTicketDTO.from_station=%s&leftTicketDTO.to_station=%s&purpose_codes=ADULT",
		date.Format("2006-01-02"),
		from,
		to,
	)

	result := &struct {
		model.Resp
		Data model.TicketResp `json:"data"`
	}{}
	resp := http.Url(u).
		SetHeader("Cookie", "RAIL_DEVICEID=;").
		SetHeader("Accept-Language", "en").
		Bind(result).
		//Debug().
		Get()

	if resp.Err() != nil {
		return nil, resp.Err()
	}

	return result.Data.Decode()
}
