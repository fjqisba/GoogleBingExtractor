package Model

import (
	"fmt"
	"net/url"
)

type CollectionTask struct {
	//任务
	TaskId int
	//关键字
	Category string
	//国家
	Country string
	//省份
	State string
	//城市
	City string
	//邮政编码
	ZipCode string
}

//https://www.bing.com/search?q=led%20strip+Abejuela+04691+email&count=30&t=web
//关键字 + 城市名 + 邮编

func (this *CollectionTask)BuildBingRequest()string  {
	category := url.QueryEscape(this.Category)
	state := url.QueryEscape(this.State)
	city := url.QueryEscape(this.City)
	zipCode := url.QueryEscape(this.ZipCode)
	if this.ZipCode == ""{
		return fmt.Sprintf("https://www.bing.com/search?count=30&q=%s+%s+%s+email&t=web",category,state,city)
	}
	return fmt.Sprintf("https://www.bing.com/search?count=30&q=%s+%s+%s+email&t=web",category,city,zipCode)
}