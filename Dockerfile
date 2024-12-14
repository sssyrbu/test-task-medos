FROM golang:latest

WORKDIR /app

COPY . .

RUN go mod download

RUN JWT_SECRET_KEY=$(openssl rand -hex 32) && \
    echo "JWT_SECRET_KEY=$JWT_SECRET_KEY" >> /etc/environment

RUN GOOS=linux go build -o main .

ENV JWT_SECRET_KEY=${JWT_SECRET_KEY}

ENTRYPOINT ["./main"]