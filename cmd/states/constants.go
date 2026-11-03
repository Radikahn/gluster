package states

type DaemonMode int

const (
	Orchestrator DaemonMode = iota
	Worker
)

var DaemonModeName = map[DaemonMode]string{
	Orchestrator: "orchestrator",
	Worker:       "worker",
}
