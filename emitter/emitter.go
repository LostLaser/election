package emitter

//Emitter is a specialized queue for messaging
type Emitter struct {
	messages chan map[string]string
}

//New creates an instance of an emitter
func New(bufferSize int) *Emitter {
	e := new(Emitter)
	e.messages = make(chan map[string]string, bufferSize)

	return e
}

//Write will add a new message to the emitter
func (e *Emitter) Write(from string, to string, details string) {
	e.messages <- map[string]string{"from": from, "to": to, "details": details}
}

//Read returns the oldest message in the emitter, will block if no message is available
func (e *Emitter) Read() string {
	val := <-e.messages
	return val["from"] + "sent" + val["details"] + "to" + val["to"]
}