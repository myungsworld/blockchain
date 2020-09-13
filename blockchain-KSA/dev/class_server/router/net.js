var express = require('express');
var router = express.Router();

var instruction = [ 'pwd','docker-compose -f docker-compose.yml up -d ca.example.com orderer.example.com peer0.org1.example.com couchdb' ,  'docker ps','docker rm -f $(docker ps -aq)', 'docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel create -o orderer.example.com:7050 -c mychannel -f /etc/hyperledger/configtx/channel.tx', 'docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@org1.example.com/msp" peer0.org1.example.com peer channel join -b mychannel.block','docker exec peer0.org1.example.com peer channel list'];
var test = [ 'pwd','ls','cd ~/fabric-samples/basic-network', 'docker exec peer0.org1.example.com peer channel list'];

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('net', { title: 'fabcar example' });
});

router.get('/:id', function(req, res, next) {
  const id = req.params.id;
  console.log(id);

  var exec = require('child_process').exec,
    child;

  child = exec(instruction[id-1], {cwd: '/home/bstudent/fabric-samples/basic-network'}, function (error, stdout, stderr) {
      console.log('stdout: ' + stdout);
      console.log('stderr: ' + stderr);
      if (error !== null) {
          console.log('exec error: ' + error);
      }
      var obj = JSON.stringify({
        anStdout:stdout,
        anStderr:stderr
      })
      res.send(obj);
  });

  
  //res.render('net', { title: 'fabcar example' });
});


module.exports = router;