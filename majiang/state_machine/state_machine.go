package state_machine

type StateMachine struct {
	magnification int64 //基础倍率

	id2scope map[int64]int64
	id2odds  map[int64]int64 //赔率
}

type StatesStat struct {
	ID    int64
	Scope int64
	Odds  int64
}

func (s StateMachine) States() []StatesStat {
	res := make([]StatesStat, 0, len(s.id2scope))
	for id, scope := range s.id2scope {
		res = append(res, StatesStat{
			ID:    id,
			Scope: scope,
			Odds:  s.id2odds[id],
		})
	}
	return res
}

func (s StateMachine) ListPlayers() []int64 {
	res := make([]int64, 0, len(s.id2scope))
	for id := range s.id2scope {
		res = append(res, id)
	}
	return res
}

func (s *StateMachine) TransferScope(winner, loser int64, rate int64) {
	delta := rate * s.magnification
	s.id2scope[winner] += delta
	s.id2scope[loser] -= delta

}
func (s *StateMachine) TransferScopeWithOdds(winner, loser int64, rate int64) {
	delta := rate * s.id2odds[winner] * s.magnification * s.id2odds[loser]
	s.id2scope[winner] += delta
	s.id2scope[loser] -= delta
}

func (s *StateMachine) SetOdds(id2odds map[int64]int64) {
	for id := range id2odds {
		s.id2odds[id] = id2odds[id]
	}
}

func (s *StateMachine) MultiplyOddsBy(id int64, times float64) {
	s.id2odds[id] = int64(float64(s.id2odds[id]) * times)
}

func New(playerSum int, magnification int64) StateMachine {
	id2scope := make(map[int64]int64, playerSum)
	id2odds := make(map[int64]int64, playerSum)
	for i := 0; i < playerSum; i++ {
		id2scope[int64(i)] = 0
		id2odds[int64(i)] = 1
	}
	return StateMachine{magnification, id2scope, id2odds}
}
