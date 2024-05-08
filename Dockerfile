FROM golang:1.21.0-alpine AS builder

WORKDIR /src

COPY ./src .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/service

FROM alpine

WORKDIR /build

COPY --from=builder /src/build ./

CMD [ "./service" ]

