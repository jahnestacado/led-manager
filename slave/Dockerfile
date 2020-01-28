FROM golang AS build
WORKDIR /go/src/led-slave
COPY main.go app.go
RUN go build -o app .


FROM ubuntu
WORKDIR /led-slave
COPY --from=0 /go/src/led-slave/app .
CMD ["./app"]
