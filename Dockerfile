FROM alpine:3.12 as alpine

RUN apk add -U --no-cache ca-certificates

FROM scratch

COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY ./bin/bookclub-backend /opt/bookclub-backend/bookclub-backend  
COPY config.env . 

EXPOSE 8081

CMD ["/opt/bookclub-backend/bookclub-backend"]

