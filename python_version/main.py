from gin_game import game
from gin_player import player

run_game = True
add_players = True
players = []
counter = 1

while add_players:
    name = input(f"Enter the name of player {counter}: (enter done to finish)\n")
    if name == "done" or name == "":
        add_players = False
    else:
        players.append(player(name, 0))
        counter += 1
        print(f"{name} added to the game.")

g = game(players)

while run_game:
    if g.closed_game:
        run_game = False
    else:
        g.display_scores()
        for p in players:
            score = input(f"Enter the score for {p.name}: \n")
            p.update_score(int(score))
        g.close_round()

print("=============================================================")
print("Final Score: \n")
g.display_scores()
details = input("Show detailed scores? (y/n) \n")
if details == "y":
    for p in players:
        p.detailed_score()
input("Game over, press enter to exit.")