FROM golang:latest

WORKDIR /app

COPY . ./

RUN make

# Script to wait for a host:port to be available
RUN curl -fsSLO https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x wait-for-it.sh

ENTRYPOINT ["./run.sh"]
