include $(GOROOT)/src/Make.inc

TARG    = hranoprovod
GOFILES = \
  hranoprovod.go\
  options.go\
  parser.go\
  types.go\
  resolver.go\
  processor.go\
  accumulator.go\

include $(GOROOT)/src/Make.cmd
