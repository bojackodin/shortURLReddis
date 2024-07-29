FROM golang:1.22-alpine AS builder

WORKDIR /shortURLReddis

RUN apk update
RUN apk add make

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN make build

FROM alpine:latest
COPY --from=builder /shortURLReddis/bin/web /bin/
ENTRYPOINT ["/bin/web"]
