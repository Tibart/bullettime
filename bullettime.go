package bullettime

import (
	"fmt"
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

func (b bullet) String() string {
	return fmt.Sprintf(" %2s | %10s | %-20s | %10s |",
		b.status.String(),
		b.reference,
		b.description,
		b.scheduled.Format("2006-01-02"))
}

type Bullets []bullet

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
