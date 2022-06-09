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
	id          uint64
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

func (b *Bullets) Add(task, reference string, dateTime time.Time) error {
	id := uint64(len(*b) + 1)
	newTask := bullet{
		id:          id,
		status:      Scheduled,
		description: task,
		reference:   reference,
		scheduled:   dateTime,
		created:     time.Now().UTC().Local(),
		modified:    time.Now().UTC().Local(),
	}
	*b = append(*b, newTask)

	return nil
}

func (b *Bullets) Remove(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	*b = append(bl[:id], bl[id+1:]...)

	return nil
}

func (b *Bullets) Complete(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id-1]

	bullet.modified = time.Now()
	bullet.status = Completed

	return nil
}

func (b *Bullets) Reschedule(id, days int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id-1]

	bullet.modified = time.Now()
	bullet.status = Rescheduled
	b.Add(bullet.description, bullet.reference, bullet.scheduled.Add(time.Duration(days)*(time.Hour*24)))

	return nil
}

func (b *Bullets) Postpone(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id-1]

	bullet.modified = time.Now()
	bullet.status = Postponed

	return nil
}

func (b *Bullets) Cancel(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id-1]

	bullet.modified = time.Now()
	bullet.status = Canceled

	return nil
}

func (b *Bullets) getRealId(id int) (int, error) {
	if id > len(*b) || id < 1 {
		return -1, fmt.Errorf("id is out of scope")
	}
	return id - 1, nil
}
