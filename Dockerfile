FROM golang:1.22-alpine as build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" \
    -o stress_test ./cmd/stress_cli/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /app/stress_test .
ENTRYPOINT [ "./stress_test" ]