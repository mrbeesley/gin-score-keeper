using System;
using System.Text;
using GinScoreKeeper.Contracts;
using GinScoreKeeper.Controllers;
using GinScoreKeeper.Models;

namespace GinScoreKeeper.Manage
{
    public class GameManager
    {
        public GameManager()
        {
            _gm = new GameController();
            _pl = new PlayerController();
        }

        private Game _state;
        private readonly GameController _gm;
        private readonly PlayerController _pl;

        public GameManager SaveGame()
        {
            Console.WriteLine();
            Console.WriteLine("Would you like to save the game data? (y|n)");
            var input = Console.ReadLine();
            if (input.ToLower() == "y")
            {
                //TODO: save the game data
                _gm.SaveGame(_state);

                Console.WriteLine("The data was saved.");
            }
            Console.WriteLine("Thanks for playing!");
            return this;
        }

        public GameManager FinishGame()
        {
            Console.WriteLine("Game Over!!!");
            Console.WriteLine();
            _gm.SetWinner(_state);
            Console.WriteLine($"The winner is - {_state.Winner.Name}");
            Console.WriteLine();
            Console.WriteLine(BuildScorecard(_state));
            return this;
        }

        public GameManager KeepScore()
        {
            while (!_state.GameOver)
            {
                var scoreCard = BuildScorecard(_state);
                Console.WriteLine(scoreCard);
                foreach (var player in _state.Players)
                {
                    Console.WriteLine($"Enter the score for {player.Name}:");
                    var needInput = true;
                    int score = 0;
                    while (needInput)
                    {
                        var input = Console.ReadLine();
                        if (int.TryParse(input, out score))
                        {
                            needInput = false;
                        }
                        else
                        {
                            Console.WriteLine("The score must be a number! Try again");
                        }
                    }
                    _pl.UpdateScore(player, score, _state.Round);
                }
                _gm.CloseRound(_state);
            }
            return this;
        }

        public GameManager SetupPlayers()
        {
            var GettingPlayers = true;
            while (GettingPlayers)
            {
                Console.WriteLine("Enter the name of the player or type done when finished (then press enter):");
                var input = Console.ReadLine();
                if (_state.Players.Count > 1 && (input == "done" || input == ""))
                {
                    GettingPlayers = false;
                }
                else
                {
                    var player = _pl.CreateNewPlayer(input);
                    _state.Players.Add(player);
                    Console.WriteLine($"{player.Name} was added to the game!");
                }
            }
            return this;
        }

        public GameManager StartGame()
        {
            _state = new Game();
            Console.WriteLine("Lets Get this game Started, first we need to know who is playing.");
            return this;
        }


        #region Util Methods

        private string BuildScorecard(Game gm)
        {
            StringBuilder ret = new StringBuilder();
            ret.AppendLine($"{BuildBorder(50, '#')}");
            ret.AppendLine();
            ret.AppendLine($"Round:{BuildSpace("Round")}{gm.Round}");
            ret.AppendLine($"Wild Card:{BuildSpace("Wild Card")}{gm.Wild}");
            ret.AppendLine();
            ret.AppendLine($"{BuildBorder(50, '-')}");
            ret.AppendLine();
            foreach (var player in gm.Players)
            {
                var space = BuildSpace(player.Name);
                ret.AppendLine($"{player.Name}:{space}{player.Score}");
            }
            ret.AppendLine();
            ret.AppendLine($"{BuildBorder(50, '#')}");
            return ret.ToString();
        }

        private string BuildSpace(string name)
        {
            var ret = new StringBuilder();
            var spaces = 40 - name.Length;
            for (int x = 0; x < spaces; x++)
            {
                ret.Append(" ");
            }
            return ret.ToString();
        }
        private string BuildBorder(int total, char c)
        {
            var ret = new StringBuilder();
            for (int x = 0; x < total; x++)
            {
                ret.Append(c);
            }
            return ret.ToString();
        }

        #endregion // Util Methods
    }
}