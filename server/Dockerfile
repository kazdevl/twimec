FROM golang:1.17-alpine as BUILDER

ENV GOCS=linux
ENV GOACH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /root
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build main.go

FROM alpine:3.14 as PROD
COPY --from=BUILDER /root/main ./
RUN mkdir config && mkdir assets
COPY --from=BUILDER /root/config/ ./config
COPY --from=BUILDER /root/assets/ ./assets