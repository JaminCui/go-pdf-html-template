package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExportPDFTemplate(c *gin.Context) {
	content, err := WKHtml2PDFTest()
	if err != nil {
		fmt.Println(err)
	}
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.pdf")
	c.Header("Content-Type", "application/octet-stream")

	c.Writer.Write(content)
}

func ExportPDF(c *gin.Context) {
	resp, err := GoTenBergPdfTest()
	if err != nil {
		fmt.Println(err)
	}
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.pdf")
	c.Header("Content-Type", "application/octet-stream")

	resp.Write(c.Writer)
	//c.Data(http.StatusOK, "application/octet-stream", content)
}
