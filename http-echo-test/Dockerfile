# Compile binary
FROM golang:alpine as compile

# hadolint ignore=DL3045
COPY httpd.go .
RUN CGO_ENABLED=0 go build httpd.go


# Make actual image
FROM scratch

COPY --from=compile /go/httpd /httpd

ENTRYPOINT ["/httpd"]
