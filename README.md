# Augmented Empathy Project Description

The purpose of this project is to study how different types of feedback based on emotion can augment people's empathy level in remote video chat. 

The project is now detecting and using facial information as the only one factor determining the participants' emotion. 

7 simple type of emotions is included in this project: happy, angry, surprised, neutral, sad, disgusted and fearful

# Implemented features
* Peer-to-peer video chat, implementing [WEBRTC](https://webrtc.org/)
* Face detection and emotion analysis
* Two type of feedback, background color and emoji
* A pop-up questionnaire including 2 questions in every 3 minutes
* REST API for future external device use.

# Dependencies
* Golang : 1.11.* - latest
* webframework : gin-gonic [repo](https://github.com/gin-gonic/gin)
* Face-api.js, refer to [this](https://github.com/justadudewhohacks/face-api.js/)


# Quick Start
* use 'cd' command to go to your repository
```console
$ cd Augmented-Empathy
```

* run backend go file to serve
```console
$ go run webserver/main.go
```

* Configure : default port number is 8081, which can be changed in file ,
```console
$ config/web_server.go
```

# Demo
`https://www.haoailan.online:8081`


# API Use
Namespaces : '/newEmotion'
1. GET

2. POST
Request body:
```json
{
    "userID": "string",
    "roomID": "string",
    "emotions": "[]emotion_detail",
}
```
```go
emotion_detail:
{
    "type": "string"
    "value": []float32
}
```
**Futher Explain:**

'value' in the emotion_detail have 8 float32 value, the first 7 values indicating the probability of emotion(angry,disgusted,fearful,happy,neutral,sad,surprised), and the 8th(last) value indicating the weight of this type of feedback to be calculated with the existing types.

Example:

Javascript in console
``` javascript
await fetch('/newEmotion',{ 
    method: 'POST',
    headers:{'content-type': 'application/json'},
    body:JSON.stringify({
        userID: 'hh',
        roomID: '1',
        emotion_detail: [
        {
            type: 'EDA',
            value: [0.5,0,1,2,0,1,0,1]
        },
        {
            type: 'Heart-Rate',
            value: [0.5,0,1,2,0,1,0,1]
        }]
    })
})
```
POSTMAN
```JSON
{
    "userID": "hh1",
    "roomID": "1",
    "emotion_detail": [
    {
        "type": "EDA",
        "value": [0.5,0,1,2,0,1,0,1]
    },
    {
        "type": "Heart-Rate",
        "value": [0.5,0,1,2,0,1,0,1]
    }]
    
}
```

# Plus: WEBRTC connection flow
1.  ![webrtc_process](webrtc_process.png)
2.  When setLocalDescription() set/write local description, RTCPeerConnection would trigger the icecandidate and send the description to remote RTCPeerConnection to update remote candidate.


