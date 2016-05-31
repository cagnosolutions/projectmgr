package projectmgr

import (
	"math"
	"time"
)

const (
	BRAINSTORMING = iota
	TESTING
	CODING
)

type Activity struct {
	Type        int
	StartTime   time.Time
	EndTime     time.Time
	ElapsedTime time.Time
}

type Project struct {
	Id   int64
	Name string
	//StartDate            time.Time
	//EndDate              time.Time
	Budget               float64
	Rate                 float64
	CodingPercent        float64
	BrainstormingPercent float64
	TestingPercent       float64
	Developers           float64
	Activities           []Activity
	CurrentActivity      int
}

func NewProject() *Project {
	return &Project{
		Id: time.Now().Unix(),
	}
}

func (p *Project) StartActivity(activity int) {
	a := Activity{
		Type:      activity,
		StartTime: time.Now(),
	}
	p.Activities = append(p.Activities, a)
}

func (p *Project) EndActivity() time.Time {
	a := p.Activities[p.CurrentActivity]
	a.EndTime = time.Now()
	a.ElapsedTime = time.Unix(a.EndTime.Unix()-a.StartTime.Unix(), 0)
	return a.ElapsedTime
}

func (p *Project) BillableTime() float64 {
	return math.Ceil(p.Budget / p.Rate)
}

func (p *Project) CodingHrs() float64 {
	return p.BillableTime() * (p.CodingPercent / 100)
}

func (p *Project) BrainstormingHrs() float64 {
	return p.BillableTime() * (p.BrainstormingPercent / 100)
}

func (p *Project) TestingHrs() float64 {
	return p.BillableTime() * (p.TestingPercent / 100)
}

func (p *Project) Days() float64 {
	return math.Ceil((p.BillableTime() / 8) / p.Developers)
}

func (p *Project) UnaccountedHrs() float64 {
	return math.Ceil(p.BillableTime() - (p.CodingHrs() + p.BrainstormingHrs() + p.TestingHrs()))
}

func (p *Project) UnaccountedPercent() float64 {
	return math.Ceil(100 - (p.CodingPercent + p.BrainstormingPercent + p.TestingPercent))
}
