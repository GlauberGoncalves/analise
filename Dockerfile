FROM golang

LABEL maintainer="Glauber <sistemas.glauber@gmail.com>"

WORKDIR /app/src/api
ENV GOPATH=/app
COPY . /app/src/api
RUN go build main.go
ENTRYPOINT ["./main"]
EXPOSE 8080