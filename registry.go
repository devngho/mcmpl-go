package mcmplgo

var tasks = make(map[string]TaskType, 0)
var requests = make(map[string]RequestType, 0)
var events = make(map[string]EventType, 0)
var listeners = make(map[EventType][]EventListener, 0)
var taskListeners = make(map[TaskType][]TaskListener, 0)
