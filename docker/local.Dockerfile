FROM cosmtrek/air

WORKDIR /usr/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
