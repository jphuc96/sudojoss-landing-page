FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 8081
VOLUME [ "/app" ]
CMD [ "/app/main" ]