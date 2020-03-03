FROM golang:1.14-buster AS builder
RUN apt-get update \
  && apt-get upgrade -y
RUN apt-get install -y git make upx
WORKDIR /app
COPY . .
RUN make deps
RUN make build_production
RUN make compress_production

FROM gcr.io/distroless/base-debian10
ENV PATH=/bin
COPY --from=builder /app/bin/semver_linux_amd64 /bin/semver
ENTRYPOINT ["/bin/semver"]