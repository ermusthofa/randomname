############################
# build executable binary
############################
FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 go build -o backend .

############################
# build service image
############################
FROM alpine
COPY --from=builder /app/backend /app/backend
COPY --from=builder /app/config.yaml /app/config.yaml
WORKDIR /app
ENV SERVER_LISTEN_ADDRESS 8080
CMD ["./backend"]