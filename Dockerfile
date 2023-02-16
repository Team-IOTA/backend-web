FROM golang:1.18

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go get

COPY *.go ./

RUN go build -o /classifie

EXPOSE 8088

CMD [ "/classifie" ]