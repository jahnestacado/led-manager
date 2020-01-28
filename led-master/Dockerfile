FROM golang AS build
WORKDIR /go/src/led-master
COPY main.go app.go
RUN go build -o app .


FROM ubuntu
WORKDIR /led-master
COPY --from=0 /go/src/led-master/app .
CMD ["./app"]
