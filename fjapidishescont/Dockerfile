FROM golang:latest
ADD . /golang/src/fjapidishescont
COPY fjapidishescont.ini /
ENV GOPATH $home/golang
RUN go get github.com/go-redis/redis
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go install fjapidishescont
ENTRYPOINT /golang/bin/fjapidishescont
EXPOSE 1650
