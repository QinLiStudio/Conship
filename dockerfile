FROM golang:1.18-alpine AS build


WORKDIR /build
COPY . .
RUN go build -o main main.go


FROM alpine:latest
LABEL MAINTAINER="https://ql.sylu.edu.cn"

WORKDIR /app
COPY --from=build /build/config/config.toml ./config/config.toml
COPY --from=build /build/main ./main

EXPOSE 8080
ENTRYPOINT ["/app/main"]