package match

type message struct {
	PlayerID uint64
	Command  string
	Data     string
}

type incomeJobs struct {
	inClient Client
	inMsg    message
}
