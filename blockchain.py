class Blockchain(object):
    def __init__(self):
        self.chain = []
        self.current_transcations = []
    
# 새로운 블록 생성하고 그 블록을 chain에 추가
    def new_block(self):
        pass

    def new_transcation(self, sender, recipient, amount):
        """
        Creates a new transaction to go into the next mined Block
        :param sender: <str> Address of the Sender
        :param recipient: <str> Address of the Recipient
        :param amount: <int> Amount
        :return: <int> The index of the Block that will hold this transaction
        """
        self.current_transcations

# @staticmethod 는 객체 만들 필요없이 클래스이름에서 바로 받아서 실행
# ex) blockchain.hash();
    @staticmethod
    def hash(block):
        pass

    @property
    def last_block(self):
        pass
