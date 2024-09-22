FROM  golang:1.23.0-alpine3.20 AS build
WORKDIR /app

COPY go.mod go.sum ./

RUN apk --no-cache add tzdata ca-certificates 

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
COPY --from=build /app/.env ./
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./cloudrun"]