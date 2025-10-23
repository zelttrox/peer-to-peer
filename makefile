i: 
	chmod +x bin/install.sh
	cd bin && ./install.sh

u:
	chmod +x bin/install.sh
	cd bin && ./uninstall.sh

r:
	cd bin && ./uninstall.sh
	cd bin && ./install.sh