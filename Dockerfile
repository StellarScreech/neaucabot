FROM golang:latest
LABEL maintainer="Atakhan Lazarev & Aziret Masalbekov <lazarev-a-auca-2022@gmail.com>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
ENTRYPOINT ["top", "-b"]