package WorkBackground

import (
	"GoogleBingExtractor/LogManager"
	"GoogleBingExtractor/Model"
	"GoogleBingExtractor/Module/CsvResult"
	"GoogleBingExtractor/Module/EmailMiner"
	"context"
	"encoding/json"
	"log"
	"time"
)

//处理单个爬取任务
//hCsvResult用于写出结果
//eZipCodeTask是爬取任务
//返回true表示任务停止,返回false表示任务继续
func handleZipCodeTask(ctx context.Context,hCsvResult *CsvResult.CsvResult,eZipCodeTask *Model.CollectionTask)bool  {

	bingUrl := eZipCodeTask.BuildBingRequest()
	var bingExt BingExtractor
	scrapeList,err := bingExt.ExtractBingUrl(bingUrl)
	if err != nil{
		LogManager.ErrorLogger.Error("提取BingUrl错误:",bingUrl,err)
		return false
	}
	//判断用户是否点击了结束任务
	if ctx.Err() != nil{
		return true
	}
	log.Println("extract bing html finished:",bingUrl)
	if len(scrapeList) == 0{
		return false
	}
	//还原任务数据
	for i:=0;i<len(scrapeList);i++{
		scrapeList[i].Category = eZipCodeTask.Category
		scrapeList[i].State = eZipCodeTask.State
		scrapeList[i].City = eZipCodeTask.City
		scrapeList[i].PostalCode = eZipCodeTask.ZipCode
	}

	//探索邮箱
	for i:=0;i<len(scrapeList);i++{
		if scrapeList[i].Email == ""{
			log.Println("guess email:",scrapeList[i].Website)
			emailList := EmailMiner.GetEmail(ctx,scrapeList[i].Website)
			emailBytes,err := json.Marshal(emailList)
			if err != nil{
				scrapeList[i].Email = string(emailBytes)
			}
		}
	}

	//写出数据
	hCsvResult.SaveScrapeData(scrapeList)
	return false
}

func StartWorkEngine(ctx context.Context,workTaskList []Model.CollectionTask,updateFunc func())  {
	defer func() {
		updateFunc()
	}()

	filePath := "./csv/" + time.Now().Format("20060102150405") + ".csv"
	hCsvResult,err := CsvResult.NewCsvResult(filePath)
	if err != nil{
		log.Println("create csv error",err)
		return
	}
	defer hCsvResult.Close()
	for _,eZipCodeTask := range workTaskList{
		select {
		case <- ctx.Done():
			return
		default:
			handleZipCodeTask(ctx,hCsvResult,&eZipCodeTask)
		}
	}
}