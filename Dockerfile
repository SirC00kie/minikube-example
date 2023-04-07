FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o minikube-example server/cmd/main.go

ENTRYPOINT [ "./minikube-example" ]