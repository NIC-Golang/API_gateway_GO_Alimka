
FROM golang:1.21.0

WORKDIR /Users/valtzmanmagnus/Desktop/GO/Jane-ecommerce-Test-repo

RUN go install github.com/air-verse/air@latest

COPY . .
RUN go mod tidy

