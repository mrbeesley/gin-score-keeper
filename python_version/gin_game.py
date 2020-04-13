class game:
    def __init__(self, players):
        self.players = players
        self.round = 1
        self.wild = {
            1: "King",
            2: "Queen",
            3: "Jack",
            4: "Ten",
            5: "Nine",
            6: "Eight",
            7: "Seven",
            8: "Six",
            9: "Five",
            10: "Four",
            11: "Three",
            12: "Two",
            13: "Ace"
        }
        self.closed_game = False

    def display_scores(self):
        current_wild =  'Ace'
        if self.round > 0 and self.round < 14: 
            current_wild = self.wild[self.round]
        
        print('------------------------------------')
        print('------------------------------------')
        print(f'Round:      {self.round}')
        print(f'Wild Card:  {current_wild}')
        
        for p in self.players:
            print(p.display_score())
        
        print('------------------------------------')
        print('------------------------------------')

    def close_round(self):
        self.round += 1
        if self.round == 14:
            self.closed_game = True