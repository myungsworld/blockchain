

let myungsCoin = new BlockChain();
myungsCoin.createTransaction(new Transaction('address1','address2',100));
myungsCoin.createTransaction(new Transaction('address2','address1',50));

console.log('\n Starting the miner...');
myungsCoin.minePendingTransactions('myungs-address');

console.log('\n Balance of myung is ',myungsCoin.getBalanceOfAddress('myungs-address'));

console.log('\n Starting the miner...');
myungsCoin.minePendingTransactions('myungs-address');

console.log('\n Balance of myung is ',myungsCoin.getBalanceOfAddress('myungs-address'));
