package arctic

var _ Logger = (*NilLogger)(nil)

type NilLogger struct{}

func (n NilLogger) Info(string, ...interface{}) {}

func (n NilLogger) Success(string, ...interface{}) {}

func (n NilLogger) Warn(string, ...interface{}) {}

func (n NilLogger) Error(string, ...interface{}) {}

func (n NilLogger) Fatal(string, ...interface{}) {}
