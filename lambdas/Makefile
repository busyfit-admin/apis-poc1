SUBDIRS := $(wildcard */.)

test: ${SUBDIRS}

build: $(SUBDIRS)

update: ${SUBDIRS}

deploy:

local: ${SUBDIRS}

$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: test local build update deploy $(SUBDIRS)