package cmd

import "github.com/gohouse/converter/majiang/state_machine"

type BaseCmd interface {
	Execute()
	Undo()
}

type Invoker struct {
	cmdStack storage
}

func (i *Invoker) Invoke(cmd BaseCmd) {
	i.cmdStack.push(cmd)
	cmd.Execute()
}

func (i *Invoker) UndoLastCmd() {
	cmd := i.cmdStack.pop()
	cmd.Undo()
}

// 不含倍率（点成 点晃）
type Single2singleCmd struct {
	winner       int64
	loser        int64
	stateMachine state_machine.StateMachine
}

type Single2restCmd struct {
	winner       int64
	rate         int64
	ignoreOdds   bool
	stateMachine state_machine.StateMachine
}

func (s *Single2singleCmd) Execute() {
	s.stateMachine.TransferScope(s.winner, s.loser, 1)
}
func (s *Single2singleCmd) Undo() {
	s.stateMachine.TransferScope(s.loser, s.winner, 1)
}

func (s *Single2restCmd) Execute() {
	for _, playerID := range s.stateMachine.ListPlayers() {
		if playerID == s.winner {
			continue
		}
		if s.ignoreOdds {
			s.stateMachine.TransferScope(s.winner, playerID, s.rate)
		} else {
			s.stateMachine.TransferScopeWithOdds(s.winner, playerID, s.rate)
		}
	}
}
func (s *Single2restCmd) Undo() {
	for _, playerID := range s.stateMachine.ListPlayers() {
		if playerID == s.winner {
			continue
		}
		if s.ignoreOdds {
			s.stateMachine.TransferScope(playerID, s.winner, s.rate)
		} else {
			s.stateMachine.TransferScopeWithOdds(playerID, s.winner, s.rate)
		}
	}
}

type MultiplyOddsCmd struct {
	player       int64
	stateMachine state_machine.StateMachine
}

func (s *MultiplyOddsCmd) Execute() {
	s.stateMachine.MultiplyOddsBy(s.player, 2)
}
func (s *MultiplyOddsCmd) Undo() {
	s.stateMachine.MultiplyOddsBy(s.player, 0.5)
}

func NewMultiplyOddsCmd(s state_machine.StateMachine, player int64) BaseCmd {
	return &MultiplyOddsCmd{
		stateMachine: s,
		player:       player,
	}
}

type ResetOddsCmd struct {
	stateMachine state_machine.StateMachine
	oldOdds      map[int64]int64
}

func (s *ResetOddsCmd) Execute() {
	newOdds := make(map[int64]int64)
	oldOdds := make(map[int64]int64)
	for _, state := range s.stateMachine.States() {
		newOdds[state.ID] = 1
		oldOdds[state.ID] = state.Odds
	}
	s.stateMachine.SetOdds(newOdds)
	s.oldOdds = oldOdds
}
func (s *ResetOddsCmd) Undo() {
	s.stateMachine.SetOdds(s.oldOdds)
}

func NewResetOddsCmd(s state_machine.StateMachine) BaseCmd {
	return &ResetOddsCmd{
		stateMachine: s,
	}
}
