FROM harbor.meitu-int.com/ops/alpine:v3.7-1

MAINTAINER lb6@meitu.com

WORKDIR /www/api-ssopa.meitu-int.com
COPY ssopa-api ssopa-api
ADD conf conf
RUN chmod +x ssopa-api

# api
EXPOSE 8080

CMD ["./ssopa-api", "--env", "prod"]
