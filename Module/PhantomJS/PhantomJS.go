package PhantomJS

import "os/exec"

func GetPageHtml(url string)string  {
	cmd := exec.Command("./rsrc./phantomjs.exe","./rsrc./page.js",url)
	outPut,err := cmd.Output()
	if err != nil{
		return ""
	}
	return string(outPut)
}