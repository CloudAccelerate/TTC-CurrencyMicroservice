FROM golang
 
ADD . /go/src/currencymicroservice
RUN go install currencymicroservice
ENTRYPOINT /go/bin/basic_web_server
 
EXPOSE 8089

