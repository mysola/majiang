package cmd

import "github.com/gohouse/converter/majiang/state_machine"

type 点晃Cmd struct {
	Single2singleCmd
}

func New点晃(s state_machine.StateMachine, winner, loser int64) BaseCmd {
	return &点晃Cmd{Single2singleCmd{
		winner:       winner,
		loser:        loser,
		stateMachine: s,
	}}
}

type 点成Cmd struct {
	Single2singleCmd
}

func New点成(s state_machine.StateMachine, winner, loser int64) BaseCmd {
	return &点成Cmd{Single2singleCmd{
		winner:       winner,
		loser:        loser,
		stateMachine: s,
	}}
}

type 贴晃Cmd struct {
	Single2restCmd
}

func New贴晃(s state_machine.StateMachine, winner int64) BaseCmd {
	return &贴晃Cmd{Single2restCmd{
		winner:       winner,
		rate:         1,
		ignoreOdds:   true,
		stateMachine: s,
	}}
}

type 硬晃Cmd struct {
	Single2restCmd
}

func New硬晃(s state_machine.StateMachine, winner int64) BaseCmd {
	return &硬晃Cmd{Single2restCmd{
		winner:       winner,
		rate:         2,
		ignoreOdds:   true,
		stateMachine: s,
	}}
}

type 摸成Cmd struct {
	Single2restCmd
	existLaizi bool //是否有赖子
}

func New摸成(s state_machine.StateMachine, winner int64, existLaizi bool) BaseCmd {
	rate := int64(2)
	if existLaizi {
		rate = 1
	}
	return &摸成Cmd{Single2restCmd{
		winner:       winner,
		rate:         rate,
		ignoreOdds:   false,
		stateMachine: s,
	}, existLaizi}
}
