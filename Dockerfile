FROM mondoolabs/golang:1.15.3-alpine
RUN mkdir /my_todo
ADD . /my_todo
WORKDIR /my_todo
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
#RUN go build -o main .
#EXPOSE 33333
## Our start command which kicks off
## our newly created binary executable
## CMD ["/app/main"]