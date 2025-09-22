package duty_system

type Dude struct {
	Name   string
	Avatar interface{}
}

type DutySystem interface {
	GetDude() (Dude, error)
}
