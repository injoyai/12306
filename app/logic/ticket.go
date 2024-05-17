package logic

import (
	"encoding/json"
	"fmt"
	"github.com/injoyai/12306/app/model"
	"github.com/injoyai/conv"
	"github.com/injoyai/goutil/net/http"
	"github.com/injoyai/goutil/oss"
	"github.com/injoyai/goutil/times"
	"github.com/injoyai/logs"
	"strings"
	"time"
)

// GetTicketList 查询车票信息
func GetTicketList(date time.Time, from, to string) (_ model.Tickets, err error) {

	//校验日期
	if err = CheckDate(date); err != nil {
		return
	}

	//校验站台名称
	if from, err = CheckStationName(from); err != nil {
		return
	}
	if to, err = CheckStationName(to); err != nil {
		return
	}

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
		//Bind(result).
		//Debug().
		Get()

	if resp.Err() != nil {
		return nil, resp.Err()
	}

	resp.Bind(&result)
	//if err := json.Unmarshal(resp.GetBody(), &result); err != nil {
	//	return err
	//}

	return result.Data.Decode()
}

func GetStationName() (map[string]string, error) {

	m := make(map[string]string)
	keuValid := "valid"
	cacheFilename := oss.UserInjoyDir("/12306/cache/stationName.json")

	//尝试从文件获取
	bs, err := oss.Read(cacheFilename)
	if err == nil {
		if json.Unmarshal(bs, &m) == nil && time.Now().Unix()-conv.Int64(m[keuValid]) < 0 {
			delete(m, keuValid)
			return m, nil
		}
	}

	logs.Debug("更新站台名称数据...")

	url := "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.9303"

	resp := http.Url(url).Get()
	if resp.Err() != nil {
		return nil, resp.Err()
	}

	/*
		响应格式如下
		var station_names ='@bjb|北京北|VAP|beijingbei|bjb|0|0357|北京|||@bjd|北京东|BOP|beijingdong|bjd|1|0357|北京|||@...'
	*/

	for _, v := range strings.Split(resp.GetBodyString(), "@") {
		/*
			得到大概如下格式
			bjb|北京北|VAP|beijingbei|bjb|0|0357|北京|||
		*/
		if list := strings.Split(v, "|"); len(list) >= 3 {
			m[list[1]] = list[2]
		}

	}

	//设置有效期
	m[keuValid] = conv.String(time.Now().AddDate(0, 0, 1).Unix())
	oss.New(cacheFilename, m)
	return m, nil
}

func CheckDate(date time.Time) error {
	if times.IntegerDay(time.Now()).Sub(date) > 0 {
		return fmt.Errorf("日期必须大于等于今天")
	}
	return nil
}

func CheckStationName(name string) (string, error) {
	nameMap, err := GetStationName()
	if err != nil {
		return "", err
	}
	if value := nameMap[name]; len(value) > 0 {
		name = value
	}

	for _, v := range nameMap {
		if v == name {
			return name, nil
		}
	}

	return "", fmt.Errorf("站台名称有误: " + name)
}
