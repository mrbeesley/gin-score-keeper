class player:
    def __init__(self, name, score):
        self.name = name
        self.score = score
        self.score_tracking = []
        self.round = 0
    
    def update_score(self, score):
        self.round += 1
        self.score_tracking.append({
            'round': self.round,
            'score': score
        })
        self.score += score

    def display_score(self):
        return f'{self.name}: {self.score}'
    
    def detailed_score(self):
        current_sum = 0
        for item in self.score_tracking:
            s = item['score']
            r = item['round']
            current_sum += s
            print(f'{self.name} | {r} | {s} | {current_sum}')