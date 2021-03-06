package message

// Election message contains an array of previously visited server ids
type Election struct {
	notified []string
}

// NewElection returns an instance of an Election
func NewElection(id string) Election {
	e := Election{}
	e.notified = make([]string, 0)
	e.AddNotified(id)
	return e
}

// AddNotified appends a new value to the notified list
func (e *Election) AddNotified(id string) {
	e.notified = append(e.notified, id)
}

// NotifiedCount gets the number of items in the notified list
func (e *Election) NotifiedCount() int {
	return len(e.notified)
}

// GetHighest will return the highest value in the notified list
func (e *Election) GetHighest() string {
	max := ""
	for _, v := range e.notified {
		if max < v {
			max = v
		}
	}
	return max
}

// Exists returns whether or not the id is in the notified list
func (e *Election) Exists(id string) bool {
	for _, v := range e.notified {
		if v == id {
			return true
		}
	}

	return false
}
