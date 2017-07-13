# coinone-collector

## Description
This program collects virtual currency data and store them to prometheus as metrics. You can see metrics on [Grafana](https://grafana.com/) dashboard. This is based on [Coinone API](http://doc.coinone.co.kr/).

## Prerequirements
* [Docker](https://www.docker.com/)

## Usage
Very simple! Just execute the program using docker-compose.
```shell
$ git clone https://github.com/odg0318/coinone-collector
$ cd coinone-collector
$ docker-compose up
```

## Available Currencies
* Bitcoin
* Ethereum
* Ethereum Classic
* Ripple

## References
* Coinone: https://coinone.co.kr/ (Thanks!)
* Prometheus: https://prometheus.io/
* Grafana: https://grafana.com/
* Docker: https://www.docker.com/
