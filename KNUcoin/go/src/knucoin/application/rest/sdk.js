'use strict';

const { CouchDBWallet, Gateway } = require('fabric-network');
var path = require('path');

async function send(type, user, func, args, res){
    try {
        
        let ccpPath = '';
        let ca_name = ''
        let url = ''
        switch(user[1]){
            case 'SalesOrg':
                url = "http://knucoin:knucoin@localhost:9984"
                ccpPath = path.resolve(__dirname, '..', 'connection_sales.json')
                break;
            case 'CustomerOrg':
                url = "http://knucoin:knucoin@localhost:10984"
                ccpPath = path.resolve(__dirname, '..', 'connection_customer.json')
                break;
        }

        const wallet = new CouchDBWallet({"url":url})

        const userExists = await wallet.exists(user[0]);
        if (!userExists) {            
            console.log(`sdk :An identity for the client user(id:${user[0]}) of ${user[1]} doesn't exists in the wallet`);
            console.log('sdk :Run the registUser.js application before retrying');
            return;
        }
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: user[0], discovery: { enabled: true, asLocalhost: true } });
        const network = await gateway.getNetwork('channelsales1');
        const contract = network.getContract('knucoin-cc');

        if(type){
            await contract.submitTransaction(func, ...args);
            console.log('sdk :Transaction has been submitted');
            await gateway.disconnect();
            res.send('success');
        }else{
            const result = await contract.evaluateTransaction(func, ...args);
            console.log(`sdk :Transaction has been evaluated, result is: ${result.toString()}`);
            res.send(result.toString());
        }
    } catch (error) {
        console.error(`sdk :Failed to submit transaction: ${error}`);
        res.send(`Failed to submit transaction: ${error}`);
    }
}

module.exports = {
    send:send
}
