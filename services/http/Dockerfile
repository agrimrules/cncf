FROM alpine:3.10 as build
COPY . /app/
WORKDIR /app/
RUN apk --no-cache add go=1.12.8-r0 gcc g++ git
RUN CGO_ENABLED=1 GOOS=linux go build ./services/http/server.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates curl=7.65.1-r0
WORKDIR /app/
COPY --from=build /app/server .
COPY ./services/http/config.toml .
EXPOSE 8080
CMD ["./server"]
