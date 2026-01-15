# Peer-to-peer #
*Made in Go* <br>
Peer-to-peer is a **file transfer program** that uses direct connection from a machine to another with TCP network protocol. It can be ran using a terminal on Linux. It may also work on Windows or Mac but **has not been tested yet, use at your own risks!** <br>

**The sending and receiving machines must be connected to the same network for the file transfer to work.** <br>

**Note that network protection like proxy rules might sometimes block file transfers even thought both machines are connected to the same network.**

# Installation #
### Dependencies ###
- [Go](https://go.dev/doc/install)
### ⭐ Using AUR ###
```bash
yay -S peer-to-peer
```
### ⭐ Using Git ###
```bash
git clone https://github.com/zelttrox/peer-to-peer.git
cd peer-to-peer/bin && chmod +x install.sh && ./install.sh
```
# Usage #
### Send a file ###
The following command will send a file transfer request to your peer, if it gets accepted, the file transfer will process. Otherwise it will be canceled.
```bash
peer send <ip>:<port> <file>
```
<img width="508" height="103" alt="image" src="https://github.com/user-attachments/assets/c71d7883-6aaf-45c2-8eca-fc9b8e6f1796" />


### Receive a file ###
The following command will display your IP on your terminal, you can send it to your peer so he can use it to send you files.
```bash
peer receive <port>
```
<img width="643" height="133" alt="image" src="https://github.com/user-attachments/assets/707bdefc-e142-4c15-8216-6aaa3cc8dc19" />

### Help menu ###
In your terminal, you can also type the following command to get the help menu directly.
```bash
peer help
```
<img width="610" height="316" alt="image" src="https://github.com/user-attachments/assets/808e019f-2809-41af-9943-368e0a0657b2" />
