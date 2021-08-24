package fb

import (
	"filebeat_mate/pkg/path"
	"filebeat_mate/pkg/util"
	"go.uber.org/zap"
)

func Start() {
	startHb()
}

func startHb() {
	err := generateHbFile()
	if err != nil {
		util.Logger().Error("generate prom config file failed ", zap.Error(err))
		return
	}
	err = util.CallScript(path.FbStartScript)
	util.Logger().Error("run start prom scripts failed ", zap.Error(err))
}

func generateHbFile() (err error) {
	// todo
	return nil
}
