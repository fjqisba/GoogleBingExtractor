package ChromeBrowser

import (
	"GoogleBingExtractor/Model"
	"context"
	"github.com/chromedp/chromedp"
	"os"
)

var(
	browser context.Context
	browserCancel context.CancelFunc
)

func QueryBingResult(request string)([]Model.ScraperData,error)  {
	taskCtx, cancel := chromedp.NewContext(browser)
	defer cancel()
	var html string
	err := chromedp.Run(taskCtx,
		chromedp.Navigate(request),
		chromedp.OuterHTML("body",&html,chromedp.ByQuery))
	if err != nil{
		return nil,err
	}
	return nil,nil
}

func init()  {
	os.Setenv("google-chrome","./chrome/chrome.exe")
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.Flag("headless",false),
	)
	browser, browserCancel = chromedp.NewExecAllocator(context.Background(), opts...)
}