FROM golang:latest

ADD . /go/src/app
CMD [" export GIN_MODE=release"]
#RUN go get github.com/anovanmaximuz/go-jwt
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
WORKDIR /go/src/app
COPY * ./
RUN go build -o base 
EXPOSE 8080
CMD ["./base"]
