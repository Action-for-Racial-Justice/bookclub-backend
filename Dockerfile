FROM alpine:3.12 as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./bin/bookclub-backend /opt/bookclub-backend/bookclub-backend  

EXPOSE 80

CMD ["/opt/tally/bookclub-backend"]

