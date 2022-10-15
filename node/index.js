var cluster = require('cluster');
var http = require('http');
var numCPUs = 32;

if (cluster.isMaster) {
 for (var i = 0; i < numCPUs; i++) {
  cluster.fork();
 }
} else {
 http.createServer(function(req, res) {
  res.writeHead(200, {"Content-Type": "application/json"});
  res.end(JSON.stringify({message: "Hello World"}));
 }).listen(8080);
}
