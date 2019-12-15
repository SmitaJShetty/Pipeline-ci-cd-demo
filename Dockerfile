
FROM golang:latest as base

LABEL maintainer="smita<smita.narasimha@gmail.com>"

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN cd cmd && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../build/main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=base /app/build/main .

EXPOSE 8080

CMD ["./main"]