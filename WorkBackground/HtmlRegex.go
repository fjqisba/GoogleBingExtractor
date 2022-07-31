package WorkBackground

import "regexp"

const strRegex_DetectEmail = `[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+`
//非常宽松的电话匹配规则
const strRegex_DetectPhone = `\:[\+\(\s\-\d\)]{3}[\(\s.\-\/\d\)]{3,20}`

var(
	regex_DetectEmail *regexp.Regexp
	regex_DetectPhone *regexp.Regexp
)

func init()  {
	regex_DetectEmail = regexp.MustCompile(strRegex_DetectEmail)
	regex_DetectPhone = regexp.MustCompile(strRegex_DetectPhone)
}
