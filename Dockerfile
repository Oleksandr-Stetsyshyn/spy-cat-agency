FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY controllers/*.go ./controllers/
COPY models/*.go ./models/
COPY middleware/*.go ./middleware/

RUN go build -o /main
EXPOSE 8080
CMD [ "/main" ]