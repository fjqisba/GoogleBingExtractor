package WorkBackground

import (
	"GoogleBingExtractor/Model"
	"GoogleBingExtractor/Module/PhantomJS"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type BingExtractor struct {
	Url string
}

func handleClassAlgo(doc *goquery.Document)Model.ScraperData {

	nodeHRef := doc.Find("a[href]")
	strAlgo := nodeHRef.Text()
	strWebsite := nodeHRef.AttrOr("href","")
	strKeyWord := doc.Find("[class='b_caption']").Text()
	if strKeyWord == ""{
		strKeyWord = doc.Find("[class='b_snippet']").Text()
	}
	strEmail := regex_DetectEmail.FindString(strKeyWord)
	strPhone := regex_DetectPhone.FindString(strKeyWord)
	strPhone = strings.ReplaceAll(strPhone,":","")
	strPhone = strings.TrimSpace(strPhone)
	return Model.ScraperData{Title: strAlgo,Website: strWebsite,Phone: strPhone,Email: strEmail}
}

func handleClassTopBorder(doc *goquery.Document)Model.ScraperData  {
	strAlgo := doc.Find("[class='b_algo']").Text()
	strWebsite := doc.Find("[class='b_attribution b_clearfix ']").Text()
	strKeyWord := doc.Find("[class='rwrl rwrl_pri rwrl_padref']").Text()
	strEmail := regex_DetectEmail.FindString(strKeyWord)
	strPhone := regex_DetectPhone.FindString(strKeyWord)
	strPhone = strings.ReplaceAll(strPhone,":","")
	strPhone = strings.TrimSpace(strPhone)
	return Model.ScraperData{Title: strAlgo,Website: strWebsite,Phone: strPhone,Email: strEmail}
}

func (this *BingExtractor)ExtractBingUrl(bingUrl string)([]Model.ScraperData,error)  {
	this.Url = bingUrl
	bingPageContent := PhantomJS.GetPageHtml(bingUrl)
	if bingPageContent == ""{
		return nil,errors.New("no html")
	}
	return this.ExtractBingHtml(bingPageContent)
}

func (this *BingExtractor)ExtractBingHtml(bingHtml string)([]Model.ScraperData,error)  {
	doc,err := goquery.NewDocumentFromReader(strings.NewReader(bingHtml))
	if err != nil{
		return nil,err
	}
	vec_Result := doc.Find("#b_results")
	if vec_Result == nil{
		return nil,errors.New("no b_results")
	}
	if len(vec_Result.Nodes) == 0{
		return nil,errors.New("no result_nodes")
	}
	var retDataList []Model.ScraperData
	doc = goquery.NewDocumentFromNode(vec_Result.Nodes[0])
	for _,eQueryResult := range doc.Children().Nodes{
		node_li := goquery.NewDocumentFromNode(eQueryResult)
		className := node_li.AttrOr("class","")
		var tmpScraperData Model.ScraperData
		if className == "b_ans b_top b_topborder"{
			tmpScraperData = handleClassTopBorder(node_li)
		}
		if className == "b_algo"{
			tmpScraperData = handleClassAlgo(node_li)
		}
		if tmpScraperData.Website == ""{
			continue
		}
		retDataList = append(retDataList, tmpScraperData)
	}
	return retDataList,nil
}