FROM golang:1.13.5

LABEL maintainer="Minura Iddamalgoda"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD [ "./main" ]
