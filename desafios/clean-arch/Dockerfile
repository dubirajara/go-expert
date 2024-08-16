FROM  golang:1.22.3-alpine3.19 AS stage

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go

FROM scratch

COPY --from=stage /app/main /main
COPY --from=stage /app/.env ./

EXPOSE 8000

CMD ["/main"]