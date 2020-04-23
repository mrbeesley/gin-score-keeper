using System;
using GinScoreKeeper.Models;

namespace GinScoreKeeper.Controllers
{
    public class PlayerController
    {
        public PlayerController()
        {
            _rand = new Random((int)DateTime.Now.Ticks);
        }

        private readonly Random _rand;

        public Player CreateNewPlayer(string name, int seedScore = 0)
        {
            var ret = new Player();
            ret.Id = _rand.Next(0, 200);
            ret.Name = name;
            ret.Score = seedScore;
            return ret;
        }

        public Player UpdateScore(Player player, int score, int round)
        {
            player.Score += score;
            player.ScoreHistory.Add(round, score);
            return player;
        }
    }
}