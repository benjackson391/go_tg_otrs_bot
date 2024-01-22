FROM golang:latest
WORKDIR /app
COPY . .
RUN go build -o gotgbot cmd/app/main.go

ENV GODEBUG="gctrace=1"
ENV GOMAXPROCS=10

RUN echo 'kernel.unprivileged_userns_clone=1' > /etc/sysctl.d/00-local-userns.conf
RUN echo 'net.core.somaxconn=65535' >> /etc/sysctl.conf

EXPOSE 8080

CMD ["./gotgbot"]
