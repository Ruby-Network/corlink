FROM golang:1.21.6

WORKDIR /usr/src/app

COPY . .

RUN make

EXPOSE 8080

CMD ["./bin/corlink-server", "start"]
