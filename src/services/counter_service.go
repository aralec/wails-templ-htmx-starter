package services

type GlobalState struct {
	count int
}

type CounterService struct {
	state *GlobalState
}

func NewCounterService() *CounterService {
	return &CounterService{
		state: &GlobalState{
			count: 0,
		},
	}
}

func (s *CounterService) GetCount() int {
	return s.state.count
}

func (cs *CounterService) Increment() {
	cs.state.count++
}

func (cs *CounterService) Decrement() {
	cs.state.count--
}
