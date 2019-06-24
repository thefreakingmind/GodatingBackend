FROM golang:1.10 AS builder
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/github.com/thefreakingmind12/godating 
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o / .
FROM scratch
COPY --from=builder /server ./
ENTRYPOINT ["./server"]

