using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using GinScoreKeeper.Models;
using Newtonsoft.Json;

namespace GinScoreKeeper.Controllers
{
    public class GameController
    {
        public GameController()
        {
            _rand = new Random((int)DateTime.Now.Ticks);
        }

        private readonly Random _rand;

        public Game CreateNewGame(List<Player> players)
        {
            var ret = new Game();
            ret.Players = players;
            ret.Round = 0;
            ret.Id = _rand.Next(0, 100000);
            return ret;
        }

        public Game CloseRound(Game game)
        {
            if (game.Round == 13)
            {
                game.GameOver = true;
            }
            else
            {
                game.Round += 1;
            }
            return game;
        }

        public Game SetWinner(Game game)
        {
            game.Winner = game.Players
                .OrderBy(p => p.Score)
                .FirstOrDefault();
            return game;
        }

        public Game SaveGame(Game game)
        {
            var json = JsonConvert.SerializeObject(game);
            File.WriteAllText($"/Users/mbeesley/Documents/{game.Id}-{DateTime.Now.ToString("yyyyMMddHHMMSS")}.json", json);
            return game;
        }
    }
}