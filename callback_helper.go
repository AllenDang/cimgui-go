package cimgui

import "C"
import "sync"

var (
	inputTextCallbackFuncMap = make(map[C.int]ImGuiInputTextCallback)
	inputTextCallbackMutex   sync.Mutex
)

func registerNewInputTextCallback(cb ImGuiInputTextCallback) C.int {
	inputTextCallbackMutex.Lock()
	defer inputTextCallbackMutex.Unlock()

	key := C.int(len(inputTextCallbackFuncMap) + 1)

	for inputTextCallbackFuncMap[key] != nil {
		key++
	}

	inputTextCallbackFuncMap[key] = cb

	return key
}

func releaseInputTextCallback(key C.int) {
	inputTextCallbackMutex.Lock()
	defer inputTextCallbackMutex.Unlock()

	delete(inputTextCallbackFuncMap, key)
}

// export invokeInputTextCallback
func invokeInputTextCallback(key C.int, data *ImGuiInputTextCallbackData) {
	if cb, ok := inputTextCallbackFuncMap[key]; ok {
		cb(data)
	}
}
