FROM golang:alpine as Builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o orion_k main.go

FROM alpine:latest
WORKDIR /app
COPY --from=Builder /app/orion_k/ /app/

CMD ["/app/orion_k"]