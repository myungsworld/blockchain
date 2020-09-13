var express = require('express');
var router = express.Router();

var instruction = [ 'npm install','node enrollAdmin.js','node registerUser.js', 'tree wallet', 'node query.js', 'node invoke.js',' node server.js &'];

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('websrv', { title: 'fabcar example' });
});

router.get('/:id', function(req, res, next) {
  const id = req.params.id;
  console.log(id);

  var exec = require('child_process').exec,
    child;

  child = exec(instruction[id-1], {cwd: '/home/bstudent/fabric-samples/fabcar/javascript'}, function (error, stdout, stderr) {
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