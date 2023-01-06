FROM golang:1.19-alpine
WORKDIR /server
COPY . .

RUN go mod tidy && go build -o /server/server ./cmd/server/

RUN make ensure
CMD /server/server
