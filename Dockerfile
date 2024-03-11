# syntax=docker/dockerfile:1

FROM golang:1.22

# DST for copy
WORKDIR /app

# Copy deps and download
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./

# Build stuffs
RUN CGO_ENABLED=0 GOOS=linux go build -o /nprode_exporter

# Now run
CMD ["/nprode_exporter"]
