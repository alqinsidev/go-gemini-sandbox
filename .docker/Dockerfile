# build
FROM golang:1.22.0-alpine as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o bin/services ./src
RUN CGO_ENABLED=0 go build -o bin/migration ./src/db/migration

FROM alpine:3.19
COPY --from=builder /app/bin/services /
COPY --from=builder /app/bin/migration /
COPY --from=builder /app/.env.example /.env
COPY --from=builder /app/src/db /app/db
EXPOSE 2701
CMD ["/services"]