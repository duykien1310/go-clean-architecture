FROM golang:alpine AS build
WORKDIR /app
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN go install github.com/swaggo/swag/cmd/swag@3c5e4861ecb16ed8586df83f9b579c234c697b9e
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN swag init
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /bin/auth_service

FROM scratch
COPY --from=build /app/resource /resource
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/config /config
COPY --from=build /bin/auth_service /bin/auth_service
EXPOSE 8000
ENTRYPOINT ["/bin/auth_service"]
