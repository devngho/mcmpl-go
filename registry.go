package mcmplgo

var tasks = map[string]TaskType{
	"OnEnable":  OnEnable,
	"EventTask": EventTask,
}
var requests = map[string]RequestType{
	"Log":           Log,
	"RegisterEvent": RegisterEvent,
	"RegisterTask":  RegisterTask,
}
var events = map[string]EventType{
	"BlockBreakEvent":    BlockBreakEvent,
	"BlockBurnEvent":     BlockBurnEvent,
	"BlockCanBuildEvent": BlockCanBuildEvent,
	"PlayerJoinEvent":    PlayerJoinEvent,
	"PlayerQuitEvent":    PlayerQuitEvent,
}
var listeners = make(map[EventType][]EventListener, 0)
var taskListeners = make(map[TaskType][]TaskListener, 0)
