FROM golang:alpine as build
WORKDIR /go/src/github.com/MartinForReal/buoy
COPY . .

RUN go install -v ./...

FROM alpine:latest
COPY --from=build /go/bin/buoy /usr/local/bin/
ENTRYPOINT [ "buoy" ]
