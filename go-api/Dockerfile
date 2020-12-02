# # build stage
FROM golang:1.14.12-alpine3.12 as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o go-api

# deploy stage
FROM alpine:3.12


WORKDIR /root/
COPY --from=builder /go/src/app/go-api .

CMD ["./go-api"]
