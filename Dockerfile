FROM golang:1.26.2

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

# I will need to remove the run with air
# RUN go build -o app ./cmd/main.go

EXPOSE 8080

CMD ["air"] 

# I will learn this mean in a bit