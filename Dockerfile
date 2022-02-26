FROM golang:1.16.4-buster AS builder

# ARG VERSION=dev
# you could give this a default value as well


WORKDIR /go/src/app
COPY src/go.mod .
COPY src/go.sum .

RUN go mod download

COPY src/ .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go 

FROM debian:buster-slim
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
ARG MONGOURI_ARG="mongodb://root:vrWKB7Pxav@192.168.8.244:27017/?authSource=admin&readPreference=primary&ssl=false" 
ENV MONGOURI=$MONGOURI_ARG
CMD ["main"]