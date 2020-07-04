const SHA256 = require('crypto-js/sha256');

class Transaction{
    constructor(fromAddress, toAddress, amount){
        this.fromAddress = fromAddress;
        this.toAddress = toAddress;
        this.amount = amount;
    }
}

class Block {
    constructor(timestamp, transactions, previousHash ='') {
        this.timestamp = timestamp;
        this.transactions = transactions;
        this.previousHash = previousHash;
        this.hash = this.calculateHash();
        this.nonce = 0;
    }
    calculateHash(){
        return SHA256(this.index + this.previousHash + this.timestamp + JSON.stringify(this.transactions) + this.nonce).toString();
    }

//Proof of Work
//컴퓨터 연산이 쥰내이 빨라서 빠르게 바꿀수없게 이걸 적용 Mining 이라고도 함 00을 많이 넣을수록 시간이 오래 걸리게 되고 비트코인은 한 블록이 10분마다 만들어지게 하는걸 목적으로 함
    mineBlock(difficulty){
        while(this.hash.substring(0,difficulty) !== Array(difficulty + 1).join("0")){
            this.nonce++;
            this.hash = this.calculateHash();
        }
        console.log("Block minded " + this.hash);
    }
}

class BlockChain {
    constructor(){
        this.chain = [this.createGenesisBlock()];
        this.difficulty = 2;
        this.pendingTransactions = [];
        this.miningReward = 100;
    }

    createGenesisBlock(){
        return new Block("07/02/2020",'Genesis block','0');
    }

    getLatestBlock(){
        return this.chain[this.chain.length - 1]
    }

    //mining의 난이도와 보상의 차이를 결정하는 곳 
    minePendingTransactions(miningRewardAddresss){
    
        let block = new Block(Date.now(), this.pendingTransactions);
        block.mineBlock(this.difficulty); 

        console.log('Block successfully mined');
        this.chain.push(block);
    //블록이 만들어지면 다시 초기화
        this.pendingTransactions= [
            new Transaction(null, miningRewardAddresss , this.miningReward)
        ];
    }

    createTransaction(transaction){
        this.pendingTransactions.push(transaction);
    }

    getBalanceOfAddress(address){
        let balance = 0;

        for(const block of this.chain){
            for(const trans of block.transactions){
                if(trans.fromAddress === address){
                    balance -= trans.amount;
                }

                if(trans.toAddress === address){
                    balance += trans.amount;
                }
            }
        }
        return balance;
    }

    //유효성 검사
    isChainValid() {
        for(let i = 1; i<this.chain.length ; i++){
            const currentBlock = this.chain[i];
            const previousBlock = this.chain[i -1];
            // 지금 블록의 해시값이 이상한지 검사
            if(currentBlock.hash !== currentBlock.calculateHash()){
                return false;
            }

            if(currentBlock.previousHash !== previousBlock.hash){
                return false;
            }
        };
        return true;
    }
}

let myungsCoin = new BlockChain();
myungsCoin.createTransaction(new Transaction('address1','address2',100));
myungsCoin.createTransaction(new Transaction('address2','address1',50));

console.log('\n Starting the miner...');
myungsCoin.minePendingTransactions('myungs-address');

console.log('\n Balance of myung is ',myungsCoin.getBalanceOfAddress('myungs-address'));

console.log('\n Starting the miner...');
myungsCoin.minePendingTransactions('myungs-address');

console.log('\n Balance of myung is ',myungsCoin.getBalanceOfAddress('myungs-address'));
