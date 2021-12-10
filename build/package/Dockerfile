FROM jamincui/alpine

ARG tmp_api_version=default_value
ENV API_VERSION=$tmp_api_version

ADD competition-tool-api /app/

WORKDIR /app

EXPOSE 8080

CMD ["./go-pdf"]