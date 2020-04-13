from gin_player import player
from gin_game import game

class lake_game:
    def __init__(self):
        self.m = player('Michael', 0)
        self.j = player('Jessica', 0)
        self.g = player('George', 0)
        self.gm = game([self.m, self.j, self.g])

    def add(self, ms, js, gs):
        self.m.update_score(ms)
        self.j.update_score(js)
        self.g.update_score(gs)
        self.gm.close_round()
    
    def scores(self):
        self.gm.display_scores()


class noah_game:
    def __init__(self):
        self.m = player('Michael', 0)
        self.j = player('Jessica', 0)
        self.g = player('George', 0)
        self.n = player('Noah', 0)
        self.gm = game([self.m, self.j, self.g, self.n])

    def add(self, ms, js, gs, ns):
        self.m.update_score(ms)
        self.j.update_score(js)
        self.g.update_score(gs)
        self.n.update_score(ns)
        self.gm.close_round()
    
    def scores(self):
        self.gm.display_scores()