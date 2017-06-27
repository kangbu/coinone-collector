import json
import urllib2
import time

from prometheus_client import start_http_server, Gauge

CURRENCY_LAST = Gauge('currency_last', '', ['currency'])
CURRENCY_FIRST = Gauge('currency_first', '', ['currency'])
CURRENCY_HIGH = Gauge('currency_high', '', ['currency'])
CURRENCY_LOW = Gauge('currency_low', '', ['currency'])
CURRENCY_VOLUME = Gauge('currency_volume', '', ['currency'])


class Ticker(object):

    def __init__(self, d):
        self.currency = d['currency']
        self.last = int(d['last'])
        self.first = int(d['first'])
        self.high = int(d['high'])
        self.low = int(d['low'])
        self.volume = float(d['volume'])

    def __repr__(self):
        return '[%s] last: %d, first: %d, high: %d, low: %d, volume: %f' % (self.currency, self.last, self.first, self.high, self.low, self.volume)

    def metric(self):
        CURRENCY_LAST.labels(self.currency).set(self.last)
        CURRENCY_FIRST.labels(self.currency).set(self.first)
        CURRENCY_HIGH.labels(self.currency).set(self.high)
        CURRENCY_LOW.labels(self.currency).set(self.low)
        CURRENCY_VOLUME.labels(self.currency).set(self.volume)


def load_data():
    res = urllib2.urlopen('https://api.coinone.co.kr/ticker?currency=all')
    return json.load(res)


def main():
    start_http_server(8000)

    while True:
        try:
            data = load_data()
        except:
            continue

        for currency in ['btc', 'eth', 'etc', 'xrp']:
            if currency not in data:
                continue

            ticker = Ticker(data[currency])
            ticker.metric()

        time.sleep(10)


if __name__ == '__main__':
    main()
