FROM golang:1.8

WORKDIR /go/src/app
COPY app . 

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 9200 

CMD ["go-wrapper", "run", "app"]
