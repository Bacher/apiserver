FROM debian
COPY build/apiserver /bin/apiserver
ENV DB_HOST "go-postgres"
ENV RPC_ADDR "frontserver-1:9999"
CMD ["/bin/apiserver"]