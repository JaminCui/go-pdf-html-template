package app

import (
	"bytes"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
)

// ParseTemplate ParseTemplate
func ParseTemplate(templatePath string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func WKHtml2PDF(body string) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	page := wkhtmltopdf.NewPageReader(strings.NewReader(body))
	page.Allow.Set("/cert/fonts")
	page.Allow.Set("/cert/images")
	page.EnableLocalFileAccess.Set(true)

	pdfg.AddPage(page)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
	pdfg.MarginRight.Set(0)
	pdfg.MarginLeft.Set(0)
	pdfg.MarginTop.Set(0)
	pdfg.MarginBottom.Set(0)
	pdfg.Dpi.Set(400)

	err = pdfg.Create()
	if err != nil {

		return nil, err
	}
	return pdfg.Bytes(), nil
}

func GoTenBerPDF(body string) (*http.Response, error) {
	httpClient := &http.Client{
		Timeout: time.Duration(50) * time.Second,
	}
	client := &gotenberg.Client{Hostname: "http://localhost:3000", HTTPClient: httpClient}
	//file, err := ioutil.ReadFile("app/cert2/index-zh.html")
	//if err != nil {
	//	log.Fatal(err)
	//}
	index, err := gotenberg.NewDocumentFromString("index.html", body)
	if err != nil {
		return nil, err
	}
	req := gotenberg.NewHTMLRequest(index)
	req.PaperSize(gotenberg.A4)
	req.Margins(gotenberg.NoMargins)
	req.Scale(0.75)

	//err = client.Store(req, "pdf-export/stored.pdf")
	//if err != nil {
	//	log.Fatal(err)
	//}
	return client.Post(req)
}
