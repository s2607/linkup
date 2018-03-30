FROM golang
#works on like 1.7 or 1.8 incase it stops working
ADD . /go/src/gitea.cscapstone.us/swiley/linkup
RUN go get -v github.com/mattn/go-sqlite3
RUN go install gitea.cscapstone.us/swiley/linkup
RUN cd  /go/src/gitea.cscapstone.us/swiley/linkup && ./mkcert.sh
RUN mkdir /app && cp /go/src/gitea.cscapstone.us/swiley/linkup/server.* /app
WORKDIR /app
ENTRYPOINT  /go/bin/linkup
EXPOSE 8080
