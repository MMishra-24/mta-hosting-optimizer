
FROM golang:latest


WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /mta-hosting-optimzer

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/mta-hosting-optimzer"]