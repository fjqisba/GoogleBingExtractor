package Model

//爬取结果

type ScraperData struct {
	//关键字
	Category string
	//省份
	State string		`json:"state"`
	//城市
	City string			`json:"city"`
	//邮编
	PostalCode string	`json:"postal_code"`
	//标题
	Title string		`json:"title"`
	//网址
	Website string		`json:"website"`
	//电话
	Phone string		`json:"phone"`
	//邮箱
	Email string		`json:"email"`
}