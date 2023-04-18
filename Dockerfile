#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/danilopimenta/family-tree-app
COPY . .
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o family-tree-app ./cmd/httpserver.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/danilopimenta/family-tree-app/family-tree-app ./family-tree-app

RUN ["chmod", "+x", "./family-tree-app"]

LABEL Name=family-tree-app
EXPOSE 8080

ENTRYPOINT ["./family-tree-app"]