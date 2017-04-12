FROM golang
 
ADD ./go/src/currencymicroservice /go/src/currencymicroservice
WORKDIR /go/src/currencymicroservice
RUN go get github.com/gorilla/mux
RUN go build -o microservice .
CMD ["/go/src/microservice"]
EXPOSE 8089