package bullettime

import (
	"errors"
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
	meeting     bool
	created     time.Time
	modified    time.Time
}

type Bullets []bullet

func (b Bullets) String() string {
	s := strings.Builder{}
	s.WriteString(" Bullet-time\n")
	s.WriteString(getLine((74)))
	for i, v := range b {
		if v.meeting {
			v.description = fmt.Sprintf("%s - %s", v.scheduled.Format("15:04"), v.description)
		}
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

// TODO: make load function
func (b *Bullets) Load(path string) error {
	return errors.New("not implemented")
}

// TODO: make save to function
func (b *Bullets) Save() error {
	return errors.New("not implemented")
}

func (b *Bullets) Add(task string, reference string, meeting bool, scheduled time.Time) error {
	id := uint64(len(*b) + 1)
	scheduled = scheduled.Round(time.Duration(15 * time.Minute))
	if !meeting {
		scheduled = time.Date(scheduled.Year(), scheduled.Month(), scheduled.Day(), 0, 0, 0, 0, time.Local)
	}
	//t := time.Date(scheduled.Year(), scheduled.Month(), scheduled.Day(), scheduled.Hour(), scheduled.Round(time.Minute*15).Minute(), 0, 0, time.Local)
	newTask := bullet{
		id:          id,
		status:      Scheduled,
		description: task,
		reference:   reference,
		meeting:     meeting,
		scheduled:   scheduled,
		created:     time.Now(),
		modified:    time.Now(),
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

	bullet := &bl[id]

	bullet.modified = time.Now()
	bullet.status = Completed

	return nil
}

func (b *Bullets) Reschedule(id, days int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.modified = time.Now()
	bullet.status = Rescheduled
	b.Add(bullet.description, bullet.reference, bullet.meeting, bullet.scheduled.Add(time.Duration(days)*(time.Hour*24)))

	return nil
}

func (b *Bullets) Postpone(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.modified = time.Now()
	bullet.status = Postponed

	return nil
}

func (b *Bullets) Cancel(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.modified = time.Now()
	bullet.status = Canceled

	return nil
}

func (b *Bullets) GetSchedule() string {
	ret := Bullets{}

	// Filter only bullets of today
	for _, v := range *b {
		if v.scheduled.Format("2006-01-02") == time.Now().Format("2006-01-02") &&
			!(v.status == Canceled || v.status == Postponed) {
			ret = append(ret, v)
		}
	}

	return ret.String()
}

func (b *Bullets) getRealId(id int) (int, error) {
	if id > len(*b) || id < 1 {
		return -1, fmt.Errorf("id is out of scope")
	}
	return id - 1, nil
}
