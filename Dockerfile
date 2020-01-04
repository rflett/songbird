FROM golang:1.13-alpine
WORKDIR /app
COPY . .
RUN apk add git --no-cache
RUN go get github.com/dghubble/go-twitter/twitter github.com/dghubble/oauth1 github.com/joho/godotenv
RUN go build -o songbird
CMD ["./songbird"]
