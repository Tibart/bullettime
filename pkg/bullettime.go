package bullettime

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
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

type Bullet struct {
	Id          uint
	Status      Status
	Reference   string
	Description string
	DateTime    time.Time
	Meeting     bool
	Created     time.Time
	Modified    time.Time
}

type Bullets []Bullet

func (b Bullets) String() string {
	s := strings.Builder{}
	s.WriteString(" Bullet-time\n")
	s.WriteString(getLine((74)))
	for _, v := range b {
		if isMeeting(v.DateTime) {
			v.Description = fmt.Sprintf("%s - %s", v.DateTime.Format("15:04"), v.Description)
		}
		s.WriteString(
			fmt.Sprintf(" %03d | %s | %-12s | %-34s | %10s |\n",
				v.Id,
				v.Status.String(),
				v.Reference,
				v.Description,
				v.DateTime.Format("2006-01-02")))
	}

	return s.String()
}

func isMeeting(time time.Time) bool {
	if time.Hour() == 0 && time.Minute() == 0 {
		return false
	}

	return true
}

func getLine(width int) string {
	ln := strings.Builder{}
	for i := 0; i < width; i++ {
		ln.WriteRune('-')
	}
	ln.WriteRune('\n')

	return ln.String()
}

func (b *Bullets) Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("path '%s' does not exist", path)
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	if len(f) == 0 {
		return fmt.Errorf("config file is empty")
	}

	err = json.Unmarshal(f, b)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bullets) Save(path string) error {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("path '%s' does not exist", path)
	}

	d, err := json.Marshal(*b)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, d, 0644); err != nil {
		return err
	}

	return nil
}

func (b *Bullets) Add(bullet Bullet) error {
	// Set id
	bullet.Id = b.getNextId()

	// Set date and time
	bullet.DateTime = bullet.DateTime.Round(time.Duration(15 * time.Minute))
	if bullet.DateTime.Hour() == 0 && bullet.DateTime.Minute() == 0 {
		bullet.DateTime = time.Date(bullet.DateTime.Year(), bullet.DateTime.Month(), bullet.DateTime.Day(), 0, 0, 0, 0, time.Local)
	}

	// Set create date time
	bullet.Created = time.Now().Local()
	bullet.Modified = time.Now().Local()

	*b = append(*b, bullet)

	return nil
}

func (b *Bullets) Remove(id int) error {
	i, err := b.getIndex(id)
	if err != nil {
		return err
	}

	bl := *b
	l := len(bl)
	if l == 1 {
		*b = Bullets{}
	} else if i == 0 {
		*b = bl[1:]
	} else if i+1 == l {
		*b = bl[:i]
	} else {
		*b = append(bl[:i], bl[i+1:]...)
	}

	return nil
}

func (b *Bullets) Complete(id int) error {
	var err error
	id, err = b.getIndex(id)
	if err != nil {
		return err
	}

	bl := *b
	bullet := &bl[id]

	bullet.Status = Completed
	bullet.Modified = time.Now()

	return nil
}

func (b *Bullets) Reschedule(id int) error {
	var err error
	id, err = b.getIndex(id)
	if err != nil {
		return err
	}

	bl := *b
	bullet := &bl[id]
	bullet.Status = Rescheduled
	bullet.Modified = time.Now().Local()

	nb := Bullet{}
	nb.Description = bullet.Description
	dayOfTheWeek := time.Now().Weekday()
	skipDays := 1
	if dayOfTheWeek == time.Thursday || dayOfTheWeek == time.Friday || dayOfTheWeek == time.Saturday {
		skipDays = 8 - int(dayOfTheWeek)
	}
	dur, _ := time.ParseDuration(fmt.Sprintf("%dh00m", skipDays*24))
	nb.DateTime = bullet.DateTime.Add(dur)

	b.Add(nb)

	return nil
}

func (b *Bullets) Postpone(id int) error {
	var err error
	id, err = b.getIndex(id)
	if err != nil {
		return err
	}

	bl := *b
	bullet := &bl[id]

	bullet.Modified = time.Now()
	bullet.Status = Postponed

	return nil
}

func (b *Bullets) Cancel(id int) error {
	var err error
	id, err = b.getIndex(id)
	if err != nil {
		return err
	}

	bl := *b
	bullet := &bl[id]

	bullet.Modified = time.Now()
	bullet.Status = Canceled

	return nil
}

func (b *Bullets) TodaysSchedule() Bullets {
	// Filter only bullets of today
	ret := Bullets{}
	for _, v := range *b {
		if (v.DateTime.Format("2006-01-02") >= time.Now().Format("2006-01-02") || v.Status == Scheduled) &&
			!(v.Status == Canceled || v.Status == Postponed) {
			ret = append(ret, v)
		}
	}

	return ret
}

func (b *Bullets) getIndex(id int) (int, error) {
	if id < 1 {
		return -1, fmt.Errorf("id is out of scope")
	}
	bl := *b
	for i := 0; i < len(bl); i++ {
		if bl[i].Id == uint(id) {
			return i, nil
		}
	}

	return -1, errors.New("id not found")
}

func (b *Bullets) getNextId() uint {
	var h uint = 0
	for _, v := range *b {
		if v.Id > h {
			h = v.Id
		}
	}

	return h + 1
}
