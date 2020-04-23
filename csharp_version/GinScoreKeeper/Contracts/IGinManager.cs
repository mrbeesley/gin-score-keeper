namespace GinScoreKeeper.Contracts
{
    public interface IGinManager
    {
        bool StartGame();
        bool SetupPlayers();
        bool KeepScore();
        bool FinishGame();
    }
}