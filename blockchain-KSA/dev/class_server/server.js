// ExpressJS Setup
const express = require('express');
const app = express();
var bodyParser = require('body-parser');

const fs = require('fs');
const path = require('path');

// Constants
const PORT = 8080;
const HOST = '0.0.0.0';

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'ejs');

// use static file
app.use(express.static(path.join(__dirname, 'public')));

// configure app to use body-parser
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

// Router
var indexRouter = require('./router/index');
var networkRouter = require('./router/net');
var chaincodeRouter = require('./router/cc');
 var webserverRouter = require('./router/websrv');
 var webclientRouter = require('./router/webclnt');

app.use('/', indexRouter);
app.use('/net', networkRouter);
app.use('/cc', chaincodeRouter);
 app.use('/websrv', webserverRouter);
 app.use('/webclnt', webclientRouter);

// server start
app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);
