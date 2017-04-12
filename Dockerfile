FROM golang
 
ADD ./go/src/currencymicroservice /go/src/currencymicroservice
RUN go get github.com/gorilla/mux
RUN go install currencymicroservice
ENTRYPOINT /go/bin/basic_web_server
 
EXPOSE 8089