package game

import (
	"github.com/gohouse/converter/majiang/cmd"
	"github.com/gohouse/converter/majiang/state_machine"
)

type Controller struct {
	stateMachine state_machine.StateMachine
	id2player    map[int64]string
	cmdInvoker   cmd.Invoker
}

func NewGame(magnification int64, players []string) Controller {
	id2player := make(map[int64]string)
	for idx, player := range players {
		id2player[int64(idx)] = player
	}
	return Controller{
		stateMachine: state_machine.New(len(players), magnification),
		id2player:    id2player,
	}
}

type SingleScopeStat struct {
	Player string
	state_machine.StatesStat
}

func (c Controller) PrintScope() []SingleScopeStat {
	res := make([]SingleScopeStat, 0)
	for _, state := range c.stateMachine.States() {
		res = append(res, SingleScopeStat{
			Player:     c.id2player[state.ID],
			StatesStat: state,
		})
	}
	return res
}

func (c *Controller) Step点晃(winnerID, loserID int64) {
	c.cmdInvoker.Invoke(cmd.New点晃(c.stateMachine, winnerID, loserID))
}
func (c *Controller) Step点成(winnerID, loserID int64) {
	c.cmdInvoker.Invoke(cmd.New点成(c.stateMachine, winnerID, loserID))
	c.cmdInvoker.Invoke(cmd.NewResetOddsCmd(c.stateMachine))
}
func (c *Controller) Step贴晃(winnerID int64) {
	c.cmdInvoker.Invoke(cmd.New贴晃(c.stateMachine, winnerID))
}
func (c *Controller) Step硬晃(winnerID int64) {
	c.cmdInvoker.Invoke(cmd.New硬晃(c.stateMachine, winnerID))
}
func (c *Controller) Step摸成(winnerID int64, existLaizi bool) {
	c.cmdInvoker.Invoke(cmd.New摸成(c.stateMachine, winnerID, existLaizi))
	c.cmdInvoker.Invoke(cmd.NewResetOddsCmd(c.stateMachine))
}

func (c *Controller) Step刷赖子(playerID int64) {
	c.cmdInvoker.Invoke(cmd.NewMultiplyOddsCmd(c.stateMachine, playerID))
}

func (c *Controller) Step流局() {
	c.cmdInvoker.Invoke(cmd.NewResetOddsCmd(c.stateMachine))
}

func (c *Controller) RevokeLastOp() {
	c.cmdInvoker.UndoLastCmd()
}
