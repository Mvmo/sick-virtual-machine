ifeq ($(PREFIX),)
    PREFIX := /usr/local
endif

install: sick
	install -d $(DESTDIR)$(PREFIX)/bin/
	install sick $(DESTDIR)$(PREFIX)/bin/
	rm sick

sick:
	go build -o sick cmd/virtual-machine/main.go
