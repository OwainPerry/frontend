FROM golang:1.12.3-alpine3.9 as builder
RUN mkdir /app 
ENV GODIR /go
ENV WORKDIR ${GODIR}/src/github.com/operry/frontend
ENV GOPATH ${GODIR}
WORKDIR ${WORKDIR}
COPY . ${WORKDIR}
RUN go build -o /app/frontend .

FROM alpine:3.9
COPY --from=builder /app/frontend /frontend
EXPOSE 9090
CMD ["/frontend"]