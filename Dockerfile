FROM golang:1.17.2 as build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOOS=linux CGO_ENABLE=0 go build -o /backend


FROM alpine
COPY --from=build /backend ./backend
CMD ["./backend"]