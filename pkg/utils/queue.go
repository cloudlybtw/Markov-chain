package utils

type Queue struct {
	queue []string
}

func NewQueue() *Queue {
	return &Queue{
		queue: make([]string, 0),
	}
}

func (q *Queue) Push(v string) {
	q.queue = append(q.queue, v)
}

func (q *Queue) Pop() string {
	if len(q.queue) == 0 {
		return "Empty queue"
	}
	element := q.queue[0]
	q.queue = q.queue[1:]

	return element
}

func (q *Queue) GetString() string {
	str := ""
	for i, a := range q.queue {
		str += a
		if i+1 != len(q.queue) {
			str += " "
		}
	}
	return str
}

func (q *Queue) Len() int {
	return len(q.queue)
}
