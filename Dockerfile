FROM golang:1.13-alpine

COPY ./go.mod /app/

WORKDIR /app/
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build

EXPOSE 8080
CMD ["./go-k8s-app"]