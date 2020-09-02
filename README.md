# WebRTCDemo
基于golang提供web服务器的WebRTC demo


# Demo列表
## 1. local_media_demo
简单的本地摄像头调用,将摄像头的内容放到video标签中进行播放.

## 2. local_screen_replay
简单的本地屏幕抓取\录制\下载.

## 3. signal_server_demo
一个简单的信令服务器通信的demo

## 4.local_peerconnection_demo
本地端对端通信，除了进行信令服务器交换信息外，已具备webrtc实现一对一视频聊天的基础。

## 5. remote_chat
远程一对一视频聊天。由于webrtc不允许远程使用http调用摄像头等媒体设备，因此这里也将web服务器改成了https
演示地址：`https://www.haoailan.online:8081`.


# 使用
本项目采用go作为web服务器实现,webrtc实现采用纯js,只依赖jquery.

然后运行:
```
go run webserver/main.go
```

web服务器就已经启动,默认端口是8081,你可以在`config/web_server.go`文件中进行修改.

# Demo flow
    1.  ![webrtc_process](webrtc_process.png)
    2.  When setLocalDescription() set/write local description, RTCPeerConnection would trigger the icecandidate and send the description to remote RTCPeerConnection to update remote candidate.


# WEBRTC study
RTCPeerConnection
---

    Users' flow

    1. Register the onicecandidate handler: sends any ICE candidates to the other peer
    2. on addstream handler:handle the displaying of video stream
    3. message handler: 
        
        Your signaling server should also have a handler for messages received from the other peer. 
        If the message contains the RTCSessionDescription object, it should be added to the RTCPeerConnection object using the setRemoteDescription() method.
        If the message contains the RTCIceCandidate object, it should be added to the RTCPeerConnection object using the addIceCandidate() method.
        
    4. Utilize getUserMedia() to set up local media stream and add it to the RTCPeerConnection object using the addStream() method
    5. Start offer/answer negotiation process
        Caller: createOffer()=>
        Callee: createAnswer()

---
RTCDataChannel

Servers and Protocols
需要服务器做signalling
Turn and Stun Server（做穿刺， 在NAT网络中）

Security
https for signalling
DTLS for data

