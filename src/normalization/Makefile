# Copyright 2010 Hannes Baldursson. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc #$(GOARCH)

TARG=idn/normalization
GOFILES=\
	normalization.go\
	tables.go\
	
	
include $(GOROOT)/src/Make.pkg

#maketables: maketables.go
#	$(GC) maketables.go
#	$(LD) -o maketables maketables.$O



tables:	maketables.py
#	$(GC) maketables.go#
#	$(LD) -o maketables maketables.$O
	python maketables.py > tables.go
	gofmt -w tables.go
#	rm -f maketables
