FROM golang
#works on like 1.7 or 1.8 incase it stops working
ADD . /go/src/gitea.cscapstone.us/swiley/linkup
RUN go get -v github.com/mattn/go-sqlite3
RUN go install gitea.cscapstone.us/swiley/linkup
RUN cd  /go/src/gitea.cscapstone.us/swiley/linkup && ./mkcert.sh
RUN ls -l /go/src/gitea.cscapstone.us/swiley/linkup/
RUN mkdir /app && cp /go/src/gitea.cscapstone.us/swiley/linkup/server.* /app && cp /go/src/gitea.cscapstone.us/swiley/linkup/*.sh /app&& cp /go/src/gitea.cscapstone.us/swiley/linkup/*.tpl /app &&cp -r /go/src/gitea.cscapstone.us/swiley/linkup/css /app/css&&cp -r /go/src/gitea.cscapstone.us/swiley/linkup/js /app/js&&cp -r /go/src/gitea.cscapstone.us/swiley/linkup/imgs /app/imgs && cp /go/bin/linkup /app/ && cp /go/src/gitea.cscapstone.us/swiley/linkup/*.html /app/
WORKDIR /app
ENTRYPOINT  /app/runapp.sh
EXPOSE 8090
