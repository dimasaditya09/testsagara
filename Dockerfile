FROM golang:1.15


WORKDIR /

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .


EXPOSE 8000

RUN go build -o testsagara

CMD ["./testsagara"]