FROM golang:latest
WORKDIR /go/src/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /go/bin/app
RUN chmod +x wait-for-it.sh
EXPOSE 3000
CMD [ "./wait-for-it.sh" , "postgres:5432" , "--strict" , "--timeout=300" , "--", "/go/bin/app" ]