package class

import "regexp"

// chainToString 消息链转文本
func chainToString(chain []Chain) (fullText string) {

	for _, element := range chain {
		fullText += element.Text
	}
	return

}

// isLegitAnswer 是否为合法答案 [选择题]
func isLegitAnswer(answer string) (ok bool) {

	ok, _ = regexp.MatchString("[a-zA-Z]", answer)
	return

}
