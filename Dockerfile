FROM golang:alpine AS builder

RUN apk update && apk add --no-cache 'git=~2'

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/goginapp/
COPY . .

RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/main .

FROM alpine:3

WORKDIR /

# Copy our static executable.
COPY --from=builder /go/main /go/main

ENV PORT 3333
ENV GIN_MODE release
EXPOSE 3333

WORKDIR /go

# Run the Go Gin binary.
ENTRYPOINT ["/go/main"]
