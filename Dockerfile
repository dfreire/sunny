FROM golang:1.6

RUN wget https://github.com/Masterminds/glide/releases/download/0.10.2/glide-0.10.2-linux-amd64.tar.gz
RUN tar xvf glide-0.10.2-linux-amd64.tar.gz
RUN mv linux-amd64/glide /usr/bin/

COPY . /go/src/github.com/dfreire/sunny
WORKDIR /go/src/github.com/dfreire/sunny
RUN go build -o sunny main.go

