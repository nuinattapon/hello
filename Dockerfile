FROM golang:alpine AS builder
RUN apk --update add ca-certificates tzdata

WORKDIR /app
COPY . /app/
RUN go mod tidy
RUN CGO_ENABLED=0 go build -o /app/hello

FROM alpine
# RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Bangkok
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime

# Create a non-root user
# RUN adduser -D -s /bin/sh hello

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/hello /app/
COPY template.html /app/

# Change ownership to the non-root user
# RUN chown -R hello:hello /app

# Switch to the non-root user
# USER hello

EXPOSE 80
ENTRYPOINT ["/app/hello"]
