// eventsHandler.go is responsible for executing appropriate functions based on events that come into the events bus
// The main purpose of this system is to create greater abstratction between calling and event during an action and
// completing the approrpriate task for that action
package events

//TODO: wrap the commandManager functions for executing a command on a machine via SSH

func ExecuteEvent() {
	// TODO: determine how events will be executed and when!
	//
	// Does there need to be a runtime tick system? or is there a way to create a hash watcher?
	// Something realtime
}
