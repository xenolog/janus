package main

import (
	"fmt"
	"gopkg.in/xenolog/go-tiny-logger.v1"
)

type WriterBuffType *[]byte

//////////
// Emulate library func
type LibType struct {
	log         *logger.Logger
	localValue  string
	GlogalValue string
	buff        WriterBuffType
}

// io.Writer compatible writer for
func (s *LibType) Write(p []byte) (int, error) {
	if s.buff == nil {
		s.buff = &p
	} else {
		*s.buff = append(*s.buff, p...)
	}
	return len(p), nil
}

func (s *LibType) localMethod() WriterBuffType {
	s.log.Debug("localMethod called")
	return s.buff
}

func (s *LibType) GlobalMethod() WriterBuffType {
	s.log.Debug("GlobalMethod called")
	return s.buff
}

func (s *LibType) AddToBuff(data string) {
	if _, err := s.Write([]byte(data)); err != nil {
		s.log.Debug("Can't collect string '%s':%v", data, err)
	} else {
		s.log.Debug("String '%s' successfully collected in buffer", data)
	}
}

func (s *LibType) CallApp() { //ee interface{}) {
	s.log.Debug("Before start AppFunction")
	func(s *LibType) {
		defer s.log.Debug("1-st Defer for AppFunction")
		s.log.Debug("AppFunction running")
		s.AddToBuff("zzz")
		defer s.log.Debug("2-nd Defer for AppFunction")
	}(s)
	s.log.Debug("After start AppFunction")
}

func NewLibType(log *logger.Logger) *LibType {
	rv := new(LibType)
	rv.log = log
	return rv
}

//////////
// External module, that use Lib

//////////
var Log *logger.Logger

func init() {
	Log = logger.New()
	Log.SetMinimalFacility(logger.LOG_D)
}

func main() {
	fmt.Printf("Testing Lib directly...\n")
	ll := NewLibType(Log)
	ll.AddToBuff("xxx")
	ll.AddToBuff("yyy")
	ll.CallApp()
	Log.Info("%v", *ll.GlobalMethod())
	Log.Info("=> %s", *ll.GlobalMethod())
	fmt.Printf("Testing Lib throught module usage...\n")
	fmt.Printf("End.\n")
}
