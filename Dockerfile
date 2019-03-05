FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
VOLUME /app
EXPOSE 8081
CMD [ "/app/main" ]