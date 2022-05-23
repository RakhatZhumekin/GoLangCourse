FROM golang:alpine

RUN mkdir /app

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . /app

ENV PORT 11037

RUN CGO_ENABLED=0 go build

EXPOSE 11037

CMD ["./Final"]