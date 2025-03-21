FROM golang:bullseye AS builder

RUN apt-get update && apt-get install -y --no-install-recommends git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go build -o /app/cmd/site
# EXPOSE 8080
# ENTRYPOINT ["/app/cmd/site"]
FROM scratch

COPY --from=builder /app/cmd/site /site


EXPOSE 8080
ENTRYPOINT ["/site"]

