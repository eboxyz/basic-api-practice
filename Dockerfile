FROM golang:1.10 
ENV PATH = "$GOPATH/bin:${PATH}"

COPY . src/league-of-leggos
WORKDIR src/league-of-leggos

RUN go build -o main 

EXPOSE 8000

CMD ["./main"]