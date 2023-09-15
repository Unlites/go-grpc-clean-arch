FROM golang:1.20.8-alpine3.18 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o /airport ./cmd/airport/main.go

FROM scratch

COPY --from=builder airport /bin/airport

CMD ["/bin/airport"]