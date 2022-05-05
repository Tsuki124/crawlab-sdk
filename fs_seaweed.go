package crawlab_sdk

import (
	"github.com/Tsuki124/crawlab-sdk/internal/engines"
	"github.com/Tsuki124/crawlab-sdk/internal/interfaces"
	"sync"
)

var FS = fsService{_FileSyses: make(map[string]interfaces.SeaweedFS)}

type fsService struct {
	_Mtx sync.RWMutex
	_FileSyses map[string]interfaces.SeaweedFS
}

func (my *fsService) Path(prePath string) interfaces.SeaweedFS {
	my._Mtx.RLock()
	fileSys, ok := my._FileSyses[prePath]
	my._Mtx.RUnlock()
	if ok {
		return fileSys
	}

	my._Mtx.Lock()
	fileSys = engines.NewSeaweedFS(prePath)
	my._FileSyses[prePath] = fileSys
	my._Mtx.Unlock()

	return fileSys
}

