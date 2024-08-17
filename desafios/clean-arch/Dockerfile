FROM  golang:1.23.0-alpine3.20 AS stage

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch

COPY --from=stage /app/main /main
COPY --from=stage /app/.env ./

CMD ["/main"]