FROM golang:alpine as builder
WORKDIR /build
ADD . /build/
RUN CGO_ENABLED=0 go build -o initialize .
FROM gcr.io/distroless/base
WORKDIR "/opt/sob"
COPY config config
COPY --from=builder /build/initialize .
COPY rootCA.pem /etc/ssl/certs/
CMD ["./initialize"]
