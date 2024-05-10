FROM golang:1.21.0-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/service cmd/api/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /src/build ./

CMD [ "./service" ]