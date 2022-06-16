package bullettime

import (
	"encoding/json"
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
	Id          uint64
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
	for i, v := range b {
		if isMeeting(v.DateTime) {
			v.Description = fmt.Sprintf("%s - %s", v.DateTime.Format("15:04"), v.Description)
		}
		s.WriteString(
			fmt.Sprintf(" %02d | %s | %-34s | %-12s | %10s |\n",
				i+1,
				v.Status.String(),
				v.Description,
				v.Reference,
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
	id := uint64(len(*b) + 1)
	bullet.Id = id

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
	bl := *b
	id, _ = b.getRealId(id)

	*b = append(bl[:id], bl[id+1:]...)

	return nil
}

func (b *Bullets) Complete(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.Modified = time.Now()
	bullet.Status = Completed

	return nil
}

func (b *Bullets) Reschedule(id, days int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.Modified = time.Now().Local()
	bullet.Status = Rescheduled

	nb := Bullet{}
	nb.Description = bullet.Description
	nb.DateTime = bullet.DateTime.Add(time.Hour * 24)

	b.Add(nb)

	return nil
}

func (b *Bullets) Postpone(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.Modified = time.Now()
	bullet.Status = Postponed

	return nil
}

func (b *Bullets) Cancel(id int) error {
	bl := *b
	id, _ = b.getRealId(id)

	bullet := &bl[id]

	bullet.Modified = time.Now()
	bullet.Status = Canceled

	return nil
}

func (b *Bullets) GetSchedule() string {
	ret := Bullets{}

	// Filter only bullets of today
	for _, v := range *b {
		if v.DateTime.Format("2006-01-02") == time.Now().Format("2006-01-02") &&
			!(v.Status == Canceled || v.Status == Postponed) {
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
