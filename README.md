# go-pdf-html-template
#build
```cassandraql
docker build -t go-pdf -f Dockerfile .
```
#run
```cassandraql
docker run -p 8080:8080 -it -d go-pdf:latest
```
#curl
```cassandraql
http://127.0.0.1:8080/export/pdfTemplate
```