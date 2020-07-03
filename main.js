const SHA256 = require('crypto-js/sha256');

class Block {
    constructor(index, timestamp, data, previousHash ='') {
        this.index = index;
        this.timestamp = timestamp;
        this.data = data;
        this.previousHash = previousHash;
        this.hash = this.calculateHash();
        this.nonce = 0;
    }
    calculateHash(){
        return SHA256(this.index + this.previousHash + this.timestamp + JSON.stringify(this.data) + this.nonce).toString();
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
        this.difficulty = 4;
    }

    createGenesisBlock(){
        return new Block(0,"07/02/2020",'Genesis block','0');
    }

    getLatestBlock(){
        return this.chain[this.chain.length - 1];
    }

    addBlock(newBlock){
        newBlock.previousHash = this.getLatestBlock().hash;
        newBlock.mineBlock(this.difficulty);
        this.chain.push(newBlock);
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
        }
        return true;
    }
}

let myungsCoin = new BlockChain();
console.log('이상훈바보');
console.log('Mining Block 1...');
myungsCoin.addBlock(new Block(1, "07/03/2020",{amount:4}));
console.log('Mining Block 2...');
myungsCoin.addBlock(new Block(2, "07/03/2020",{amount:10}));


