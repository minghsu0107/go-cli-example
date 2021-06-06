FROM golang:1.15.2 AS builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG tag
ARG build_number
ENV DRONE_TAG=$tag
ENV DRONE_BUILD_NUMBER=$build_number

COPY . .
RUN make build-linux

FROM scratch

COPY --from=builder /app/gce /bin/

ENTRYPOINT ["/bin/gce"]