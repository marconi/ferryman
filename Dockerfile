FROM google/golang
ADD . /gopath/src/github.com/marconi/ferryman
RUN go get -t github.com/marconi/ferryman
WORKDIR /gopath/src/github.com/marconi/ferryman
