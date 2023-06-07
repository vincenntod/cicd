FROM golang

WORKDIR /app

COPY . .

RUN go build -o mini-project-2

CMD ["./mini-project-2"]