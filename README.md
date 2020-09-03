# WebRTCDemo
Based on golang server, implement Webrtc to build Peer-to-peer video conference.

## 1. remote chat
Demo：`https://www.haoailan.online:8081`.


# Run
```
go run webserver/main.go
```

default port number is 8081, which can be changed `config/web_server.go`


# Demo flow
1.  ![webrtc_process](webrtc_process.png)
2.  When setLocalDescription() set/write local description, RTCPeerConnection would trigger the icecandidate and send the description to remote RTCPeerConnection to update remote candidate.


# WEBRTC study
RTCPeerConnection
    
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

RTCDataChannel

Servers and Protocols

Security
https for signalling
DTLS for data

