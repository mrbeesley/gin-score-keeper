using System.Collections.Generic;

namespace GinScoreKeeper.Models
{
    public class Game
    {

        public Game()
        {
            Players = new List<Player>();
            Round = 1;
            _wildCards = CreateWilds();
            GameOver = false;
        }

        private readonly Dictionary<int, string> _wildCards;

        public int Id { get; set; }
        public int Round { get; set; }

        public bool GameOver { get; set; }
        public Player Winner { get; set; }
        public List<Player> Players { get; set; }
        public string Wild => _wildCards[Round];

        private Dictionary<int, string> CreateWilds()
        {
            var ret = new Dictionary<int, string>();
            ret.Add(1, "Kings");
            ret.Add(2, "Queens");
            ret.Add(3, "Jacks");
            ret.Add(4, "Tens");
            ret.Add(5, "Nines");
            ret.Add(6, "Eights");
            ret.Add(7, "Sevens");
            ret.Add(8, "Sixes");
            ret.Add(9, "Fives");
            ret.Add(10, "Fours");
            ret.Add(11, "Threes");
            ret.Add(12, "Twos");
            ret.Add(13, "Aces");
            return ret;
        }

    }
}