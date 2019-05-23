FROM golang:alpine as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .

FROM alpine
WORKDIR /app
COPY ./html5up-identity ./html5up-identity
COPY --from=builder /app/main .
VOLUME /app
EXPOSE 80
CMD [ "/app/main" ]
