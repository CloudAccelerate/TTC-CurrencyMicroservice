FROM golang

RUN mkdir /go/src/currencymicroservice
ADD . go/src/currencymicroservice
WORKDIR go/src/currencymicroservice
RUN go build -o microservice .
CMD ["go/src/microservice"]
 
EXPOSE 8089