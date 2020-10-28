FROM golang

ADD . /go/src/app
CMD [" export GIN_MODE=debug"]
WORKDIR /go/src/app
COPY * ./
#RUN go mod init github.com/anovanmaximuz/go-jwt
#RUN go get github.com/anovanmaximuz/go-jwt
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
#WORKDIR /go/src/app
#COPY * ./
#RUN go build -o main.go
EXPOSE 8180
CMD ["./main"]
