FROM golang:1.14-alpine as build
COPY . /app
WORKDIR /app
RUN go build -o server

FROM alpine:3.12
WORKDIR /app
COPY --from=build /app/server .
ENTRYPOINT ["./server"]
