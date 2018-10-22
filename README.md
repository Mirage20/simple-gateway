## Another simple gateway


#### Sample Usage

1. Download the [`routes.json`](routes.json) file

2. Run the gateway using 

        ./simple-gateway -routes=routes.json -port=8080
    Or

        docker run -p 8080:8080 -v <routes.json-directory>:/var mirage20/simple-gateway -routes=/var/routes.json

#### Writing filters

See [`filters`](pkg/gateway/filters) package for samples
