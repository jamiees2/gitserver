FROM golang:1.19-alpine

RUN apk add git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /gitserver

EXPOSE 5000

CMD [ "/gitserver" ]
