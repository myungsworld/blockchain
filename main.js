const SHA256 = require('crypto-js/sha256');

class Block {
    constructor(index, timestamp, data, previousHash ='') {
        this.index = index;
        this.timestamp = timestamp;
        this.data = data;
        this.previousHash = previousHash;
        this.hash= this.calculateHash();
    }
    calculateHash(){
        return SHA256(this.index + this.previousHash + this.timestamp + JSON.stringify(this.data)).toString();
    }
}

class BlockChain {
    constructor(){
        this.chain = [this.createGenesisBlock()];
    }

    createGenesisBlock(){
        return new Block(0,"07/02/2020",'Genesis block','0');
    }

    getLatestBlock(){
        return this.chain[this.chain.length - 1];
    }

    addBlock(newBlock){
        newBlock.previousHash = this.getLatestBlock().hash;
        newBlock.hash = newBlock.calculateHash();
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
myungsCoin.addBlock(new Block(1, "07/03/2020",{amount:4}));
myungsCoin.addBlock(new Block(2, "07/03/2020",{amount:10}));


console.log('is blockchain valid?',myungsCoin.isChainValid());
myungsCoin.chain[1].data = {amount:200};
myungsCoin.chain[1].hash = myungsCoin.chain[1].calculateHash();
console.log('is blockchain valid?',myungsCoin.isChainValid());


//JSON.stringify(문자열로 변환할 값,null이면 모든속성을 포함,마지막 숫자는 space개수)
 console.log(JSON.stringify(myungsCoin,null,4));