package TestUnit

import (
	"GoogleBingExtractor/WorkBackground"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func getHtmlContent(request string)string  {
	hReq,err := http.NewRequest("GET",request,nil)
	if err != nil{
		return ""
	}
	hReq.Header.Set("User-Agent","Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)")
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return ""
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return ""
	}
	return string(respBytes)
}


func TestQueryHtml(t *testing.T) {
	var bingExt WorkBackground.BingExtractor

	//requestUrl := "https://www.bing.com/search?q=led%20strip+australia+email&count=30&t=web"
	//bingExt.ExtractBingUrl(requestUrl)
	//return
	//bingHtml := PhantomJS.GetPageHtml(requestUrl)

	hFile,_ := os.Open("D:\\test.html")
	bingHtml, _ := ioutil.ReadAll(hFile)
	hFile.Close()
	bingExt.ExtractBingHtml(string(bingHtml))
}
