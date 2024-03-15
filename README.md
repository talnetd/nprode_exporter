# nprode_exporter
Simple network beats to monitor your concerned endpoints over the network peer or internet

# Why?
Modern applications depends on several third-party services (sso and payment gateways etc)
and it becomes necessity to monitor those endpoints as well. This tool will monitor them 
for you at specified ports and push the metrics back to a Prometheus Push Gateway.

## Getting Started
You just need go installed in your system.

## Dependencies
Install go deps.
```
go get
```

## Install
For now I have worked out a simple docker setup and a configuration file but you can also
build it yourself and deployed it. Later will make releases for popular formats such as .deb.
```
docker build --tag nprode_exporter .
```

Or build it on your machine.
```
go build -o ./nprode_exporter
```

## Run
Currently this repo come with a sample yaml file. Should be clear enough to create your own
by copying it. To run the program simply pass the flag --config /to/your/sample/yaml
```
./nprode_exporter --config /to/your/sample/yaml
```

Or modify Dockerfile and adjust the config path and run the docker image as container instead.
```
docker run -it -d --name nprode nprode_exporter:latest
```

And you might need a dummy pushgateway then ...
```
docker pull prom/pushgateway
docker run -d -p 9091:9091 prom/pushgateway
```


## Authors
TAL:@talnetd
