using System.Collections.Generic;

namespace GinScoreKeeper.Models
{
    public class Player
    {
        public Player()
        {
            ScoreHistory = new Dictionary<int, int>();
        }

        public int Id { get; set; }
        public string Name { get; set; }
        public int Score { get; set; }
        public Dictionary<int, int> ScoreHistory { get; set; }
    }
}