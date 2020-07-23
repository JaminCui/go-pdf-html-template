package main

import (
	"fmt"
	pdf_generator "go-pdf-html-template/pdf-generator"
)

func main() {
	PDFTemplateTest()
}

func PDFTemplateTest() {
	templatePath := "app/cert/index-zh.html"

	//path for download pdf
	outputPath := "storage/index-en.pdf"

	pdft := pdf_generator.NewPDFHtmlTemplate()

	templateData := struct {
		BgImage      string
		GameName     string
		ThemeImage   string
		PrizeLang    string
		Prize0       string
		Prize1       string
		TeamNo       string
		TeamName     string
		StudentGroup string
		TeacherGroup string
		GameEndAt    string
		BackImage    string
	}{
		BgImage:      "http://cloud.makex.cc/public/cert/images/bg-2020-zh.png",
		GameName:     "总决赛",
		ThemeImage:   "http://cloud.makex.cc/public/cert/images/ultimateWarrior-icon-zh.png",
		PrizeLang:    "zh",
		Prize0:       "一等奖",
		Prize1:       "prize title",
		TeamNo:       "X10086",
		TeamName:     "中国移动",
		StudentGroup: "苏老板",
		TeacherGroup: "测试大佬",
		GameEndAt:    "2020-07-06",
		BackImage:    "http://cloud.makex.cc/public/cert/images/makex-2020-back.png",
	}

	if err := pdft.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := pdft.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
