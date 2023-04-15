FROM golang:1.19

WORKDIR /app

ARG ENV=sample
ENV ENV=${ENV}

COPY go.mod ./
RUN go mod download

EXPOSE 5000

COPY . ./

RUN go build -o /gpt-twitter-integration

CMD [ "/gpt-twitter-integration" ]