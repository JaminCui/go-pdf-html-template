package pdf_generator

import (
	"bytes"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"html/template"
	"log"
	"strings"
)

type PDFHtmlTemplate struct {
	data     map[string]interface{}
	body     string
	filePath string
}

func NewPDFHtmlTemplate() *PDFHtmlTemplate {
	return &PDFHtmlTemplate{}
}
func (pdfHtmlTemplate *PDFHtmlTemplate) ParseTemplate(templatePath string, data interface{}) error {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	pdfHtmlTemplate.body = buf.String()
	return nil
}

func (pdfHtmlTemplate *PDFHtmlTemplate) GeneratePDF(pdfPath string) (bool, error) {
	//t := time.Now().Unix()

	//err1 := ioutil.WriteFile("cloneTemplate/"+strconv.FormatInt(int64(t), 10)+".html", []byte(pdfHtmlTemplate.body), 0644)
	//if err1 != nil {
	//	panic(err1)
	//}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(pdfHtmlTemplate.body))
	//page := wkhtmltopdf.NewPage("cloneTemplate/" + strconv.FormatInt(int64(t), 10) + ".html")

	pdfg.AddPage(page)

	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.MarginRight.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)

	pdfg.Dpi.Set(400)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(pdfPath)
	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
