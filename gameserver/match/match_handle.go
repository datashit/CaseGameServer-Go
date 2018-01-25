package match

var incomeMessageJobs = make(chan incomeJobs, 1000)

func createMatchHandleWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go match_handle(incomeMessageJobs)
	}
}

func match_handle(job <-chan incomeJobs) {
	for j := range job {
		switch j.msg.Command {
		case "MATCH":
		case "LOBBY":
		case "CHAT":
		case "CONTACTS":
		case "PARTY":
		default:
		}
	}
}
