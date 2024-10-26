FROM golang:1.23-alpine3.20

WORKDIR /app

COPY . /app

RUN go get .

RUN go test ./...

RUN go build -o lecture-cicd-go-app ./main.go

EXPOSE 8080

CMD [ "/app/lecture-cicd-go-app" ]
