package cmd

type storage struct {
	commands []BaseCmd
}

func (s *storage) push(cmd BaseCmd) {
	s.commands = append(s.commands, cmd)
}

func (s *storage) pop() BaseCmd {
	last := s.commands[len(s.commands)-1]
	s.commands = s.commands[:len(s.commands)-1]
	return last
}
