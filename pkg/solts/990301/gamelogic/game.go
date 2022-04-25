package gamelogic

import (
	"fmt"
	wlzb "go-game-sdk/example/game_LaBa/990301/msg"
	"go-game-sdk/example/game_LaBa/labacom/config"
	"go-game-sdk/inter"

	"github.com/kubegames/kubegames-games/internal/pkg/score"

	"github.com/kubegames/kubegames-sdk/pkg/log"
	"github.com/kubegames/kubegames-sdk/pkg/player"

	"github.com/kubegames/kubegames-sdk/pkg/table"
)

//用户回来的消息
type UserRebackInfo struct {
	FreeGameTimes int
	FreeGameGold  int64
	LastBetGold   int64
	FreeGameIndex int
	EnterFreeGame bool
}

type LaBaRoom struct {
}

//初始化桌子
func (lbr *LaBaRoom) InitTable(table table.TableInterface) {
	//log.Tracef("init table num %d", table.GetID())
	g := new(Game)
	g.InitTable(table)
	g.Init(&config.LBConfig)
	table.BindGame(g)
}

func (lbr *LaBaRoom) UserExit(user player.PlayerInterface) {
}

func (lbr *LaBaRoom) AIUserLogin(user inter.AIUserInter, game table.TableHandler) {
}

func (g *Game) InitTable(table table.TableInterface) {
	g.table = table
}

func (g *Game) CloseTable() {
	if g.user != nil {
		if g.FreeGameTimes != 0 || g.EnterFreeGame {
			str := fmt.Sprintf("%v,%v,%v,%v,%v", g.FreeGameTimes, g.FreeGameGold, g.EnterFreeGame, g.FreeGameIndex, g.LastBet)
			g.user.SetTableData(str)
		}
		g.user.SendRecord(g.table.GetGameNum(), g.user.GetScore()-g.curr, g.AllBet*int64(g.Line), 0, g.UserTotalWin, "")
		g.table.WriteLogs(g.user.GetID(), fmt.Sprintln("游戏结束金币:", score.GetScoreStr(g.user.GetScore())))

		g.curr = g.user.GetScore()

		g.table.KickOut(g.user)
		g.table.EndGame()
	}
}

//用户坐下
func (g *Game) OnActionUserSitDown(user player.PlayerInterface, chairId int, config string) int {
	g.FreeGameTimes = 0
	g.EnterFreeGame = false
	g.LastBet = 0
	g.AllBet = 0
	return 1
}

func (g *Game) UserExit(user player.PlayerInterface) bool {
	if g.FreeGameTimes != 0 || g.EnterFreeGame {
		str := fmt.Sprintf("%v,%v,%v,%v,%v", g.FreeGameTimes, g.FreeGameGold, g.EnterFreeGame, g.FreeGameIndex, g.LastBet)
		user.SetTableData(str)
	}
	user.SendRecord(g.table.GetGameNum(), user.GetScore()-g.curr, g.AllBet*int64(g.Line), 0, g.UserTotalWin, "")
	g.table.WriteLogs(g.user.GetID(), fmt.Sprintln("游戏结束金币:", score.GetScoreStr(user.GetScore())))
	g.table.EndGame()
	g.curr = user.GetScore()
	return true
}

func (g *Game) LeaveGame(user player.PlayerInterface) bool {
	if g.FreeGameTimes != 0 || g.EnterFreeGame {
		str := fmt.Sprintf("%v,%v,%v,%v,%v", g.FreeGameTimes, g.FreeGameGold, g.EnterFreeGame, g.FreeGameIndex, g.LastBet)
		user.SetTableData(str)
	}
	user.SendRecord(g.table.GetGameNum(), user.GetScore()-g.curr, g.AllBet*int64(g.Line), 0, g.UserTotalWin, "")
	g.table.WriteLogs(g.user.GetID(), fmt.Sprintln("游戏结束金币:", score.GetScoreStr(user.GetScore())))
	g.table.EndGame()
	g.curr = user.GetScore()
	return true
}

//游戏消息
func (g *Game) OnGameMessage(subCmd int32, buffer []byte, user player.PlayerInterface) {
	switch subCmd {
	case int32(wlzb.MsgIDC2S_Bet):
		g.OnUserBet(buffer)
		break
	case int32(wlzb.MsgIDC2S_AskSence):
		g.SendSence()
		break
	case int32(wlzb.MsgIDC2S_ChoseFreeGameTimes):
		g.ChoseFreeGameTimes(buffer)
		break

	// TODO: 线上环境注释
	case int32(wlzb.MsgIDC2S_Test):
		//g.handleTestMsg(buffer)
	}
}

func (g *Game) UserReady(user player.PlayerInterface) bool {
	return true
}

//场景消息
func (g *Game) SendScene(user player.PlayerInterface) bool {
	g.user = user
	g.UserTotalWin = 0
	g.GetRoomconfig()
	g.GetRebackInfo()
	g.curr = user.GetScore()
	senddata := new(wlzb.Sence)
	senddata.BetValue = append(senddata.BetValue, g.BetArr...)
	senddata.Gold = user.GetScore()
	if g.FreeGameTimes != 0 || g.EnterFreeGame {
		senddata.FreeGameTimes = int32(g.FreeGameTimes)
		senddata.FreeGameGold = g.FreeGameGold
		senddata.FreeGameIndex = g.FreeGameIndex + 1
		senddata.EnterFreeGame = g.EnterFreeGame
		for i := 0; i < len(g.BetArr); i++ {
			if g.BetArr[i] == int32(g.LastBet) {
				senddata.LastBetIndex = int32(i)
				break
			}
		}
	} else {
		senddata.LastBetIndex = 0
	}

	log.Tracef("场景消息：%v", senddata)
	user.SendMsg(int32(wlzb.ReMsgIDS2C_SenceID), senddata)
	g.table.StartGame()
	g.table.WriteLogs(g.user.GetID(), fmt.Sprintln("游戏开始金币:", score.GetScoreStr(user.GetScore())))
	return true
}

func (g *Game) GameStart(user player.PlayerInterface) bool {
	return true
}

func (g *Game) ResetTable() {

}
