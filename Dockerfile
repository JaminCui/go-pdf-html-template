FROM makeblock/bsdo:2.0.0 AS builder

COPY ./ /build/

WORKDIR /build

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o go-pdf ./main.go

FROM registry.cn-hangzhou.aliyuncs.com/makeblock/alpine-pdf

COPY --from=builder /build/go-pdf /app/

COPY ./app/cert  /app/cert

WORKDIR /app

EXPOSE 8080

CMD ["./go-pdf"]