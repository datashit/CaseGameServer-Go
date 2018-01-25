package match

type message struct {
	PlayerID uint64
	Command  string
	Data     string
}

type incomeJobs struct {
	client Client
	msg    message
}
