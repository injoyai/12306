package model

import (
	"errors"
	"fmt"
	"strings"
)

type Resp struct {
	Code    int    `json:"httpstatus"` //状态码
	Succ    bool   `json:"status"`     //是否成功
	Message string `json:"message"`    //返回信息
}

type TicketResp struct {
	Result []string          `json:"result"` //用|隔开的车辆信息集合
	Flag   string            `json:"flag"`   //未知
	Level  string            `json:"level"`  //未知
	Map    map[string]string `json:"map"`    //站台对应的中文名
}

func (this *TicketResp) Decode() ([]*Ticket, error) {
	result := []*Ticket(nil)
	for _, v := range this.Result {
		list := strings.Split(v, "|")
		if len(list) <= 37 {
			return nil, errors.New("解析车票信息失败")
		}

		result = append(result, &Ticket{
			Secret:         list[0],
			Status:         list[1],
			TrainName:      list[2],
			TrainNo:        list[3],
			TrainStartTime: list[8],
			TrainEndTime:   list[9],
			TrainSpendTime: list[10],

			FromStation:     list[6],
			FromStationName: this.Map[list[6]],
			ToStation:       list[7],
			ToStationName:   this.Map[list[7]],

			Alternate: list[37],
			SeatType: map[string]string{
				"特等座": list[25],
				"商务座": list[32],
				"一等座": list[31],
				"二等座": list[30],
				"软卧":   list[23],
				"硬卧":   list[28],
				"动卧":   list[33],
				"软座":   list[24],
				"硬座":   list[29],
				"无座":   list[26],
			},
		})
	}
	return result, nil
}

// Ticket 车票信息
type Ticket struct {
	Secret         string `json:"secret"`
	Status         string `json:"status"`         //状态
	TrainName      string `json:"trainName"`      //列车名称
	TrainNo        string `json:"trainNo"`        //车次
	TrainStartTime string `json:"trainStartTime"` //出发时间
	TrainEndTime   string `json:"trainEndTime"`   //到达时间
	TrainSpendTime string `json:"trainSpendTime"` //总耗时时间

	FromStation     string `json:"fromStation"`     //起始站
	FromStationName string `json:"fromStationName"` //起始站名称
	ToStation       string `json:"toStation"`       //终点站
	ToStationName   string `json:"toStationName"`   //终点站名称

	Alternate string            `json:"alternate"` //候补
	SeatType  map[string]string `json:"seatType"`  //座位类型,一等座,二等座...

	/*
		sd.LeftTicket = resSlice[29]
		sd.IsCanNate = resSlice[37]
	*/

}

func (this *Ticket) String() string {
	return fmt.Sprintf("状态: %s, 车次: %s, 站点: %s-%s, 时间: %s-%s, %s, %v",
		this.Status, this.TrainNo,
		this.FromStationName, this.ToStationName,
		this.TrainStartTime, this.TrainEndTime,
		this.TrainSpendTime,
		this.SeatType,
		//this.FromStation, this.FromStationName, this.ToStation, this.ToStationName,
	)
}
