FROM golang:1-alpine as builder

WORKDIR /workspace

RUN apk --no-cache --no-progress add git ca-certificates tzdata \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/*

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o informer

FROM scratch

WORKDIR /

COPY --from=builder /workspace/informer .

EXPOSE 80

ENV GIN_MODE=release

ENTRYPOINT ["/informer"]
