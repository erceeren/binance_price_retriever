FROM golang:1.19 AS build_stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY * ./
RUN go build -o binance_price_retriever

FROM alpine:3.9
COPY --from=build_stage /app/binance_price_retriever /app
CMD ["/app/binance_price_retriever"]