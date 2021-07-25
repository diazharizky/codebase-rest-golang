# run `docker-compose up builder` first
FROM alpine:3.10.2

ARG TZ=Asia/Jakarta
ENV TZ=$TZ

RUN apk add --no-cache ca-certificates tzdata

COPY bin/codebase-rest-golang /usr/local/bin

ENTRYPOINT ["/usr/local/bin/codebase-rest-golang"]
