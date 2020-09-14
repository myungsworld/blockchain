'use strict';

const FabricCAServices = require('fabric-ca-client');
const { CouchDBWallet, X509WalletMixin } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', 'connection_ca.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

async function main(args) {
    try {
        let ca_name = ''
        let url = ''
        switch(args[2]){
            case 'SalesOrg':
                ca_name = 'ca.sales.knucoin.com';
                url = "http://knucoin:knucoin@localhost:9984"
                break;
            case 'CustomerOrg':
                ca_name = 'ca.customer.knucoin.com'
                url = "http://knucoin:knucoin@localhost:10984"
                break;
        }

        const caInfo = ccp.certificateAuthorities[ca_name];
        const caTLSCACerts = caInfo.tlsCACerts.pem;
        const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);
       
        /* CouchDBWallet */
        const wallet = new CouchDBWallet({"url":url})
        const adminExists = await wallet.exists(args[0])
        
        if (adminExists) {
            console.log(`An identity for the admin user(id:${args[0]}) of ${args[2]} already exiss in the wallet`);
            return;
        }

        const enrollment = await ca.enroll({ enrollmentID: args[0], enrollmentSecret: args[1] });
        const identity = X509WalletMixin.createIdentity(args[2], enrollment.certificate, enrollment.key.toBytes());
        await wallet.import(args[0], identity);
        console.log(`Successfully enrolled admin user(id:${args[0]}) of ${args[2]} and imported it into the wallet`);


    } catch (error) {
        console.error(`Failed to enroll admin user(id:${args[0]}) of ${args[2]}: ${error}`);
        process.exit(1);
    }
}

main(['admin', 'adminpw', 'SalesOrg'])
main(['admin', 'adminpw', 'CustomerOrg'])