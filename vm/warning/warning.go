package warning

type warning struct {
	msg string
}

func New(msg string) warning {
	return warning{msg}
}

func (w warning) Error() string {
	return w.msg
}
