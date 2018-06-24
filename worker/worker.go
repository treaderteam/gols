package worker

import (
	"log"
	"os"
	"time"
)

// IntervalWorker is daemon, which
// trigger given function once in
// specified interval
type IntervalWorker intervalWorker

// New creates new IntervalWorker
func New(name string, triggerer func(), interval time.Duration) IntervalWorker {
	return IntervalWorker(intervalWorker{
		name:      name,
		triggerer: triggerer,
		interval:  interval,
		l:         log.New(os.Stdout, name+": ", log.LstdFlags),
	})
}

// AppendTriggerDate delays first trigger till given date
func (i *IntervalWorker) AppendTriggerDate(td time.Time) {
	(*i).triggerdate = td
}

type intervalWorker struct {
	// name serve for logging
	name string

	// Triggerer is function to be activated
	triggerer func()

	// Triggerdate is date, when trigger function
	triggerdate time.Time

	// Interval is how much add to triggerdate if it
	// is periodic
	interval time.Duration

	// l for interval worker
	l *log.Logger
}

// Start invokes interval worker
func (i IntervalWorker) Start() {
	i.l.Println("Interval worker starting")

	triggerdate := i.triggerdate

	_defaulttime := time.Time{}

	if triggerdate == _defaulttime {
		triggerdate = time.Now().Add(i.interval)
	}

_triggerloop:
	for {
		after := time.After(time.Until(triggerdate))

		select {
		case <-after:
			i.triggerer()
			triggerdate = triggerdate.Add(i.interval)
			// kill <- true
			continue _triggerloop
		}
	}
}
