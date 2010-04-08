# Copyright 2010 Hannes Baldursson. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=idn/tld
GOFILES=\
	tld.go\
	tables.go\

include $(GOROOT)/src/Make.pkg


maketables: maketables.go
	$(GC) maketables.go
	$(LD) -o maketables maketables.$O

tables:	maketables
	$(GC) maketables.go
	$(LD) -o maketables maketables.$O
	./maketables --tables=all > tables.go
	gofmt -w tables.go
	rm -f maketables

testtables: maketables
	echo '***' Be sure to make tables and make install first
	./maketables -test
	rm -f maketables
