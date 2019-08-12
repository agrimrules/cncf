FROM alpine:3.10 as build
COPY . /app/
WORKDIR /app/
RUN apk --no-cache add go=1.12.6-r0 gcc g++ git
RUN CGO_ENABLED=1 GOOS=linux go build -o grpc-server ./services/grpc/main.go

FROM alpine:3.10
RUN apk --no-cache add ca-certificates curl=7.65.1-r0
WORKDIR /app/
COPY --from=build /app/grpc-server .
EXPOSE 50051
CMD ["./grpc-server"]
