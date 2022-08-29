FROM golang:1.15.3-alpine3.12

EXPOSE 9000

RUN apk undate \
	&&apk add --no-cache \
	mysql-client \
	build-base\
	
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 
COPY ./grpc_entrypoint.sh /usr/local/bin/grpc_entrypoint.sh
RUN  /bin/chomd +x/usr/local/bin/grpc-entrypoint.sh

RUN go build cmd/main.go
RUN mv main /usr/local/bin/

CMD["main"]
ENTRYPOINT["grpc_entrypoint.sh"]

