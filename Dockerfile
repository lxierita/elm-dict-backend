FROM golang:1.14

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go build -o main .
EXPOSE 8000
CMD [ "/app/main" ]