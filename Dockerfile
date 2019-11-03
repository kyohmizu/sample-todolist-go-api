FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ .
RUN go build .

CMD ["./go-api"]

# --- #

FROM golang:latest

COPY --from=builder /app/go-api /usr/local/bin/

EXPOSE 9999
CMD ["go-api"]

