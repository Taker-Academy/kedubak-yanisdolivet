FROM golang:1.22.2

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD [ "go", "run", "example/kedubak-yanisdolivet" ]
