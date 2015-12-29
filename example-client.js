'use strict';


/**
* Example node UDP client that talks to the Johnny Cache server
*/

const dgram = require('dgram');
const client = dgram.createSocket('udp4');

const config = makeMessage({
  "method": "SET_NAMESPACE",
  "namespace": "projects"
});

const post = {
  "namespace": "projects",
  "method": "SET",
  "data": {
    "name": "bob",
    "dateStart": "Today"
  }
}

const post2 = makeMessage({
  "namespace": "projects",
  "method": "SET",
  "data": {
    "name": "Something Cool",
    "dateStart": "Today"
  }
});

function makeMessage(obj) {
  return new Buffer(JSON.stringify(obj))
}

function send(message) {
  client.send(message, 0, message.length, '2000', '127.0.0.1', (err, bytes) => {
    if (err) throw err;
    console.log('UDP message sent to 127.0.0.1');  
  });
}

client.on('listening', function () {
    var address = client.address();
    console.log('UDP Server listening on ' + address.address + ":" + address.port);
});

client.on('message', (msg, rinfo) => {
  console.log(JSON.parse(msg.toString()));
});

// Send initial conntection config
send(config);

// Save some data
setTimeout(() => {
  send(makeMessage(post))
}, 2000)

setTimeout(() => {
  send(post2)
}, 4000)

// Get some data
setTimeout(() => {
  send(makeMessage({
    "namespace": "projects",
    "method": "GET",
    "query": {
      "dateStart": "Today"
    }
  }))
}, 6000)






