package app

func WKHtml2PDFTest() ([]byte, error) {
	templatePath := "app/cert/index-zh.html"

	templateData := map[string]interface{}{
		"BgImage":      "/Users/jamin/makeblock-jamin/go-pdf-html-template/app/cert/images/bg-2020-zh.png",
		"GameName":     "总决赛",
		"ThemeImage":   "/Users/jamin/makeblock-jamin/go-pdf-html-template/app/cert/images/premier-icon-zh.png",
		"PrizeLang":    "zh",
		"Prize0":       "一等奖",
		"Prize1":       "prize title",
		"TeamNo":       "X10086",
		"TeamName":     "中国移动",
		"StudentGroup": "苏老板",
		"TeacherGroup": "测试大佬",
		"GameEndAt":    "2020-07-06",
		"BackImage":    "/Users/jamin/makeblock-jamin/go-pdf-html-template/app/cert/images/makex-2020-back.png",
	}

	body, err := ParseTemplate(templatePath, templateData)
	if err != nil {
		return nil, err
	}

	return WKHtml2PDF(body)
}
