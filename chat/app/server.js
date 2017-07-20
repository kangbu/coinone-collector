const io = require('socket.io-client')
const prom = require('prom-client');
const promServer  = require('express')();

const chatCounter = new prom.Counter({ name: 'chat_count', help: 'chat count' });

promServer.get('/metrics', (req, res) => {
    res.set('Content-Type', prom.register.contentType)
    res.end(prom.register.metrics())
})

prom.collectDefaultMetrics()
promServer.listen(9200)

console.log('Prmetheus server is listening 9200 port...'

socket = io('https://push.coinone.co.kr/chat', {transports: ['websocket']});

socket.on('connect', function(){
    console.log('Connected to chat server.')
});

socket.on('disconnect', function(){
    console.log('Disconnected from chat server.')
});

socket.on('previous_messages', function(d) {
});

socket.on('new_message', function(msg) {
    chatCounter.inc()
});
