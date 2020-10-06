ifeq ($(PREFIX),)
    PREFIX := /usr
endif

install: ha-mqtt-iot
    install -d $(DESTDIR)$(PREFIX)/bin/
    install -m 755 ha-mqtt-iot $(DESTDIR)$(PREFIX)/bin/