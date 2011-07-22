include $(GOROOT)/src/Make.inc

TARG    = hranoprovod
GOFILES = \
  hranoprovod.go\
  parser.go\
  types.go\
  resolver.go\

include $(GOROOT)/src/Make.cmd
