build:
	go install

install: build
	cp $$GOPATH/bin/mplayer-daemon /usr/bin
	cp systemd/mplayer-daemon.service /lib/systemd/system/
	systemctl daemon-reload
