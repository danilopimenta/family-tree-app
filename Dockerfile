#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/github.com/danilopimenta/family-tree-app
COPY . .
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o family-tree-app ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add --no-cache bash
COPY --from=builder /go/src/github.com/danilopimenta/family-tree-app /family-tree-app
COPY --from=builder /go/src/github.com/danilopimenta/family-tree-app/config/config.yml /config.yml

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh ./wait-for-it.sh
RUN ["chmod", "+x", "./wait-for-it.sh"]
LABEL Name=family-tree-app
EXPOSE 8080