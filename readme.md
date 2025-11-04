# Description #
This is a **file transfer program** that uses direct connection from a machine to another with TCP network protocol. It can be ran using a terminal on Linux. It may also work on Windows or Mac but **has not been tested yet, use at your own risks!** <br>

**The sending and receiving machines must connected to the same network for the file transfer to work.** <br>

**Note that network protection like proxy rules might sometimes block file transfers even thought both machines are connected to the same network.**

# Installation #
```sh
git clone https://github.com/zelttrox/peer-to-peer.git
cd peer-to-peer/bin && chmod +x install.sh && ./install.sh
```
# Usage #
### Send a file ###
The following command will send a file transfer request to your peer, if it gets accepted, the file transfer will process. Otherwise it will be canceled.
```sh
peer send <ip>:<port> <file>
```

### Receive a file ###
The following command will display your IP on your terminal, you can send it to your peer so he can use it to send you files.
```sh
peer receive <port>
```
