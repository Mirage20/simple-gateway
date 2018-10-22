FROM scratch
COPY simple-gateway /
ENTRYPOINT ["/simple-gateway"]
