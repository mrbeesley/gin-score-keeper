
using System;
using GinScoreKeeper.Manage;

namespace GinScoreKeeper
{
    class Program
    {
        static void Main(string[] args)
        {
            new GameManager()
                .StartGame()
                .SetupPlayers()
                .KeepScore()
                .FinishGame()
                .SaveGame();
        }
    }
}
