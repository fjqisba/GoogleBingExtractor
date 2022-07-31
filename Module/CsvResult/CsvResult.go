package CsvResult

import (
	"GoogleBingExtractor/Model"
	"encoding/csv"
	"os"
)

var(
	title = []string{
		"关键字","标题","省份","城市","邮编","电话","邮箱","官网"}
)


type CsvResult struct {
	hFile *os.File
	hCsvWriter *csv.Writer
}

func (this *CsvResult)Close()  {
	if this.hFile != nil{
		this.hFile.Close()
		this.hFile = nil
	}
}


//保存爬取的数据

func (this *CsvResult)SaveScrapeData(scrapeList []Model.ScraperData)  {
	for _,eScrapeData := range scrapeList{
		this.hCsvWriter.Write([]string{
			eScrapeData.Category,eScrapeData.Title,eScrapeData.State,eScrapeData.City,
			eScrapeData.PostalCode,eScrapeData.Phone,eScrapeData.Email,eScrapeData.Website})
	}
	this.hCsvWriter.Flush()
}

func NewCsvResult(filePath string)(*CsvResult,error)  {
	hFile,err := os.Create(filePath)
	if err != nil{
		return nil,err
	}
	hCsvWriter := csv.NewWriter(hFile)
	hCsvWriter.Write(title)
	hCsvWriter.Flush()
	return &CsvResult{hFile: hFile,hCsvWriter: hCsvWriter},nil
}
