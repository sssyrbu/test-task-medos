FROM golang:latest
WORKDIR /app
RUN JWT_SECRET_KEY=$(openssl rand -hex 32) && \
    echo "JWT_SECRET_KEY=$JWT_SECRET_KEY" > .env
COPY . .
RUN go mod download
RUN GOOS=linux go build -o main .
ENTRYPOINT ["./main"]