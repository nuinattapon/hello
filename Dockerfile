FROM golang:alpine AS builder
RUN apk --update add ca-certificates tzdata

WORKDIR /app
COPY main.go go.* /app/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /app/hello

FROM alpine
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Bangkok
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/hello /app/
COPY template.html /app/

EXPOSE 80
ENTRYPOINT ["/app/hello"]
