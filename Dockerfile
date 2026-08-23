FROM golang:1.26
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /bin/xuzhou-minewater ./cmd/server
EXPOSE 8080
ENTRYPOINT ["/bin/xuzhou-minewater"]
