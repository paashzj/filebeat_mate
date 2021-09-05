package fb

import (
	"filebeat_mate/pkg/path"
	"filebeat_mate/pkg/util"
	"go.uber.org/zap"
)

func Start() {
	startFb()
}

func startFb() {
	err := generateFbFile()
	if err != nil {
		util.Logger().Error("generate filebeat config file failed ", zap.Error(err))
		return
	}
	err = util.CallScript(path.FbStartScript)
	util.Logger().Error("run start filebeat scripts failed ", zap.Error(err))
}

func generateFbFile() (err error) {
	// todo
	return nil
}
