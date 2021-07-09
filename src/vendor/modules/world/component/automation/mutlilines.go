package automation

import (
	"container/ring"
	"sync"
)

const MaxMultiLines = 200

type MultiLines struct {
	Locker sync.Mutex
	Ring   *ring.Ring
}

func NewMultiLines() *MultiLines {
	return &MultiLines{
		Ring: ring.New(MaxMultiLines),
	}
}
func (l *MultiLines) Append(message string) {
	l.Locker.Lock()
	defer l.Locker.Unlock()
	l.Ring = l.Ring.Next()
	l.Ring.Value = message
}
func (l *MultiLines) Flush() {
	l.Locker.Lock()
	defer l.Locker.Unlock()
	l.Ring = ring.New(MaxMultiLines)
}

func (l *MultiLines) Last(count int) []string {
	l.Locker.Lock()
	defer l.Locker.Unlock()
	if count <= 0 {
		return []string{}
	}
	if count > MaxMultiLines {
		count = MaxMultiLines
	}
	result := make([]string, 0, count)

	r := l.Ring.Move(1 - count)
	var i = 0
	r.Do(func(v interface{}) {
		i = i + 1

		if i <= count && v != nil {
			result = append(result, v.(string))
		}
	})
	return result
}
