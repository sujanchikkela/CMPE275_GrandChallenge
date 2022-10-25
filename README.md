
![forthebadge made-with-go](https://forthebadge.com/images/badges/made-with-go.svg)


## CMPE275_GrandChallenge

## Project Members:
| Project Members | GitHub-Profile-Link | 
| ----- | ----- |
| Sujan Rao Chikkela | https://github.com/sujanchikkela |
| Satvik | https://github.com/ArunSatvik |
| Sakruthi Avirineni |https://github.com/sakruthiavirineni |
| Rajashekar Reddy Kommula | https://github.com/Rajashekarredde |
| Akshay Madiwalar | https://github.com/akshaymadiwalar |



![grpc](https://user-images.githubusercontent.com/27505090/197429185-d0383a19-c4bc-48c0-89a0-2ab874ed58a1.svg)

## Collaboration Plan:


## Objective:
The objective of this grand challenge is to construct a software tool that connects processes together to manage the three Vs. (Volume, Velocity, and variety) and the volatility of the process participation lifecycle.

## Approach:
The target of this grand challenge is to build a software tool having multiple clients connected to each other through a computer network and communicate their actions by passing messages to each other clients. In our case, we have distributed processes (involving a variety of processes) connected, talking to each other with varying messages using the gRPC protocol.

We have used protobuf to store the data that is being sent between clients. Protobuf is the most commonly used IDL for gRPC. In our case, all clients have the same proto format and are used by the client to call any available functions from the gRPC server.

We have implemented Bi-directional streaming RPC. Streaming is a core concept in gRPC where multiple actions take place in a single request. In Bi-directional streaming, clients and servers send messages to each other without waiting for a response. 

We have attempted to connect similar clients (Go lang clients) and stream the message service along with one python client in the message streaming process. In our architecture, 3 Go lang clients are connected to a grpc server, and a python grpc client is also connected to a grpc server. We have implemented a queueing system in the grpc server; when one client tries to send a message, a blocking queue queues the message in the grpc server and sends it to the rest of the active clients.

## Technologies and Frameworks:
  * gRpc
  * ProtoBuf
  * Blocking Queue
  * goLang
  * Python
  
## Objective
The objective of this grand challenge is to construct a software tool that connects processes together to manage the three Vs. (Volume, Velocity, and    variety) and the volatility of the process participation lifecycle.
 
 
 
 
 
 
 
 
 
 
 
 
