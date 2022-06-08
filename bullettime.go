package bullettime

import (
	"fmt"
	"strings"
	"time"
)

type Status int8

const (
	Scheduled = iota
	Completed
	Rescheduled
	Postponed
	Canceled
)

func (s Status) String() string {
	switch s {
	case Scheduled:
		return "o"
	case Completed:
		return "x"
	case Rescheduled:
		return ">"
	case Postponed:
		return "<"
	case Canceled:
		return "*"
	}

	return ""
}

type bullet struct {
	status      Status
	reference   string
	description string
	scheduled   time.Time
	created     time.Time
	modified    time.Time
}

type Bullets []bullet

func (b Bullets) String() string {
	s := strings.Builder{}
	s.WriteString(" Bullet-time\n")
	s.WriteString(getLine((74)))
	for i, v := range b {
		s.WriteString(
			fmt.Sprintf(" %02d | %s | %-34s | %-12s | %10s |\n",
				i+1,
				v.status.String(),
				v.description,
				v.reference,
				v.scheduled.Format("2006-01-02")))
	}

	return s.String()
}

func getLine(width int) string {
	ln := strings.Builder{}
	for i := 0; i < width; i++ {
		ln.WriteRune('-')
	}
	ln.WriteRune('\n')

	return ln.String()
}

func (b *Bullets) Add(task string) error {
	newTask := bullet{
		status:      Scheduled,
		description: task,
		reference:   "",
		scheduled:   time.Now().UTC().Truncate(time.Duration(time.Now().Day())),
		created:     time.Now().UTC().Local(),
		modified:    time.Now().UTC().Local(),
	}
	*b = append(*b, newTask)

	return nil
}

func (b *Bullets) Remove(id int) error {
	bl := *b
	id -= 1
	if id > len(bl) || id < 0 {
		return fmt.Errorf("id is out of scope")
	}

	*b = append(bl[:id], bl[id+1:]...)

	return nil
}

func (b *Bullets) SetCompleted(id int) error {
	ls := *b
	id -= 1
	if id > len(ls) || id < 0 {
		return fmt.Errorf("id is out of scope")
	}

	ls[id].modified = time.Now()
	ls[id].status = Completed

	return nil
}
