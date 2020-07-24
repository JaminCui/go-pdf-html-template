package app

import (
	"net/http"
)

func GoTenBergPdfTest() (*http.Response, error) {
	templatePath := "app/cert/index-zh.html"

	templateData := map[string]interface{}{
		"BgImage":      "http://cloud.makex.cc/public/cert/images/bg-2020-zh.png",
		"GameName":     "总决赛",
		"ThemeImage":   "http://cloud.makex.cc/public/cert/images/ultimateWarrior-icon-zh.png",
		"PrizeLang":    "zh",
		"Prize0":       "一等奖",
		"Prize1":       "prize title",
		"TeamNo":       "X10086",
		"TeamName":     "中国移动",
		"StudentGroup": "学生姓名列表",
		"TeacherGroup": "老师姓名列表",
		"GameEndAt":    "2020-07-06",
		"BackImage":    "http://cloud.makex.cc/public/cert/images/makex-2020-back.png",
	}

	body, err := ParseTemplate(templatePath, templateData)
	if err != nil {
		return nil, err
	}

	return GoTenBerPDF(body)
}
