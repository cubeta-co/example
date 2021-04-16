FROM golang:1.15-alpine as builder
WORKDIR /app
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o example main.go

FROM scratch
COPY --from=builder /app/example /example
COPY index.html ./index.html
COPY assets/ ./assets/

CMD ["/example"]