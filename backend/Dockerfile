FROM golang:1.18.3-buster

WORKDIR /test-be/src/
COPY ./go.mod ./go.sum ./main.go ./
RUN go build -o test-be ./main.go
EXPOSE 8090
CMD ["./test-be"]
