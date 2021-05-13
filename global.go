package logger

var logger = NewLogger()

var GetLevel = logger.GetLevel
var SetLevel = logger.SetLevel

// var WithContext = logger.WithContext

var Trace = logger.Trace
var Tracef = logger.Tracef
var Tracew = logger.Tracew

var Debug = logger.Debug
var Debugf = logger.Debugf
var Debugw = logger.Debugw

var Info = logger.Info
var Infof = logger.Infof
var Infow = logger.Infow

var Warn = logger.Warn
var Warnf = logger.Warnf
var Warnw = logger.Warnw

var Error = logger.Error
var Errorf = logger.Errorf
var Errorw = logger.Errorw

var Panic = logger.Panic
var Panicf = logger.Panicf
var Panicw = logger.Panicw
