# go-log

##  Usage
```go
import "github.com/Shitovdm/go-log/logger"


func main() {

	loggerInstance := logger.NewInitializedLogger()

	loggerInstance.Trace("Trace message", "category")
    loggerInstance.Tracef("category", "Trace message %s", arg)
	loggerInstance.Info("Info message", "category")
    loggerInstance.Infof("category", "Info message %s", arg)
	loggerInstance.Debug("Debug message", "category")
    loggerInstance.Debugf("category", "Debug message %s", arg)
	loggerInstance.Warning("Warning message", "category")
    loggerInstance.Warningf("category", "Warning message %s", arg)
	loggerInstance.Error("Error message", "category")
    loggerInstance.Errorf("category", "Error message %s", arg)
}
```