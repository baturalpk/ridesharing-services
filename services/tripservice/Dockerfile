FROM golang:1.18-alpine AS building

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . ./

RUN go build -o /build

# ---
FROM alpine:latest

WORKDIR /

COPY --from=building /build /exec

EXPOSE 50001

ENTRYPOINT [ "/exec" ]
