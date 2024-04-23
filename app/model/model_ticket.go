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
			Secret:          list[0],
			Status:          list[1],
			TrainName:       list[2],
			TrainNo:         list[3],
			TrainStartTime:  list[8],
			TrainEndTime:    list[9],
			FromStation:     list[6],
			FromStationName: this.Map[list[6]],
			ToStation:       list[7],
			ToStationName:   this.Map[list[7]],
		})
	}
	return result, nil
}

// Ticket 车票信息
type Ticket struct {
	Secret         string `json:"secret"`
	Status         string `json:"status"`
	TrainName      string `json:"trainName"` //列车名称
	TrainNo        string `json:"trainNo"`   //车次
	TrainStartTime string `json:"trainStartTime"`
	TrainEndTime   string `json:"trainEndTime"`

	FromStation     string `json:"fromStation"`
	FromStationName string `json:"fromStationName"`
	ToStation       string `json:"toStation"`
	ToStationName   string `json:"toStationName"`

	/*
		if resSlice[1] == "预订" {
			sd.SecretStr = resSlice[0]
			sd.LeftTicket = resSlice[29]
			sd.StartTime = resSlice[8]
			sd.ArrivalTime = resSlice[9]
			sd.DistanceTime = resSlice[10]
			sd.IsCanNate = resSlice[37]

			sd.SeatInfo = make(map[string]string)
			sd.SeatInfo["特等座"] = resSlice[utils.SeatType["特等座"]]
			sd.SeatInfo["商务座"] = resSlice[utils.SeatType["商务座"]]
			sd.SeatInfo["一等座"] = resSlice[utils.SeatType["一等座"]]
			sd.SeatInfo["二等座"] = resSlice[utils.SeatType["二等座"]]
			sd.SeatInfo["软卧"] = resSlice[utils.SeatType["软卧"]]
			sd.SeatInfo["硬卧"] = resSlice[utils.SeatType["硬卧"]]
			sd.SeatInfo["硬座"] = resSlice[utils.SeatType["硬座"]]
			sd.SeatInfo["无座"] = resSlice[utils.SeatType["无座"]]
			sd.SeatInfo["动卧"] = resSlice[utils.SeatType["动卧"]]
			sd.SeatInfo["软座"] = resSlice[utils.SeatType["软座"]]
		}
	*/

}

func (this *Ticket) String() string {
	return fmt.Sprintf("状态: %s, 车次: %s,时间: %s-%s, 出发站: %s(%s), 到达站: %s(%s)",
		this.Status, this.TrainNo,
		this.TrainStartTime, this.TrainEndTime,
		this.FromStation, this.FromStationName, this.ToStation, this.ToStationName)
}
