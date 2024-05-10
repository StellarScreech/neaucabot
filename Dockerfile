FROM golang:latest
WORKDIR /app
RUN git clone https://github.com/StellarScreech/neaucabot
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

# Define environment variables
ENV SERVER_HOST=40.233.76.67
ENV SERVER_PORT=8080

EXPOSE 8080
CMD ["./main"]
ENTRYPOINT ["top", "-b"]