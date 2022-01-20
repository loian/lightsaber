package mode

import "sync"

var TerminateRenderChannel = make(chan bool)
var renderIsRunning bool
var renderIsrunningMutex sync.Mutex
