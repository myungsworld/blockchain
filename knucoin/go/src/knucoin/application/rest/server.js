const express = require('express');
const app = express();
const bodyParser = require('body-parser');
app.use(bodyParser.json())

var path = require('path');
var sdk = require('./sdk');
var enroll = require('../sdk/enrollAdmin')
var regist = require('../sdk/registerUsers')

const PORT = 8080;
const HOST = 'localhost';

app.post('/api/enrollAdmin', (req, res) => {

    id = req.body.id
    pw = req.body.pw
    org = req.body.org

    args = [id, pw, org]
    enroll.main(args, res)
})

app.post('/api/registUser', (req, res) => {

    id = req.body.id
    admin = req.body.admin
    org = req.body.org

    args = [id, admin, org]

    regist.main(args, res)
})

app.post('/api/chargeMoney', (req, res) => {
   
    id = req.body.id
    org = req.body.org
    amount = req.body.amount

    user = [id, org]
    args = [id, amount]

    sdk.send(true, user, 'chargeMoney', args, res)
})

app.post('/api/transferMoney', (req, res) => {

    sender = req.body.sender
    receiver = req.body.receiver
    amount = req.body.amount
    org = req.body.org

    user = [sender, org]
    args = [sender, receiver, amount]

    sdk.send(true, user, 'transferMoney', args, res)
})

app.post('/api/getWallet', (req, res) => {

    id = req.body.id
    org = req.body.org

    user = [id, org]
    args = [id]

    sdk.send(false, user, 'getWallet', args, res)
})

app.listen(PORT, HOST);
console.log(`Running on http://${HOST}:${PORT}`);