gRPC presentation
Installation steps:  

install golang if you have not already
brew install go

1. 

brew install protobuf this is the protobuff compiler which we can use to generate stubs it is technically optional but, widely used


1. I had to make my go modules bin file exposed to I had to overwrite $HOME which will probably cause other problems but, for now I am going to keep it while it is being used and comment it out when it is not being used

This is the env var I have go modules installed on my ~/go with bin having all the packages that it needs
to install a package there just run

go install google.golang.org/protobuf/cmd/protoc-gen-go@<version>
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@<version>


######## golang grpc ##########
export GO_PATH=~/go
export PATH=$PATH:/$GO_PATH/bin

Note: be careful with “-” and “_” the protoc tool is picky about this. Go might let you have those in package names but, the tool might reject it. For now I’m just going to use camelCase everywhere. In a real implementation I’d be more tempted to mess with this more to get better naming convention through some workaround


1. 

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./cmrGrpc/cmrGrpc.proto

output files:
pb.go → logic for serializing and unserializing messages
grpc → file includes generated client server code which we need to implement in our server and client code


1. the vscode-proto3 extension was a big nice to have it is really useful. So you don’t miss message declarations that are missing or simple things like semicolons etc. Note I am mentioning this because by default editors aren’t going to support this file-extension




Biggest benefits for gRPC:

* leverages HTTP/2.0 for backbone of communication
* binary based protocol and serializing over the wire is alot faster than JSON
* save network bandwidth since the messages are smaller
* define interfaces so anyone who has your .proto file can create a client based on that
* Strongly typed defined messages
* defining streaming APIs relatively simple


Biggest cons:

* gRPC is not really suitable as an API to the internet its really meant to be a tool for internal backend server to server communication. This is because unlike simple HTTP, we would need to implement a client with the .proto file from the gRPC project. Which means that if the arc requires some kind of external API passing through API gateway it will have to be different than some of the backend communication. It can still be part of the arc but, it will be less consistent
* Deubgging is harder because you can’t scope at requests on the wire since they have been encoded
* Obviously less widespread and all the cons that go along with that




lessons learned:

* Don’t name things with gRPC in the name. It makes everything more confusing and less readable. In retrospect that should have been obvious because we wouldn’t name projects myProjectNameRest but, by the time it got difficult I'd gotten pretty far into things
* Originally tried to harvest two CMR microservices and started trying to implement their data flow in gPRC that was quite messy and not fruitful. I figured this was the case but, refactoring existing projects that use REST into gPRC is quite difficult. This also got into language difficulty. There is an existing clojure implementation of the gRPC framework but, without some of the buffer compiler tool support; significantly more time must be spent writing code to defining boilerplate interfaces
* Create new RPC services each time don’t try to change exiting ones because backwards compatibility is hard. Specifically if you try to do this with streams
* This did require more research than I was expecting. Trying to find use-cases in our context to demonstrate useful gRPC features was tricky since alot of it is behind the scenes and only potential optimizations. Given the short-time period the amount of demoable request->response got cut



Useful documentation:

1. https://netflixtechblog.com/practical-api-design-at-netflix-part-1-using-protobuf-fieldmask-35cfdc606518 and https://netflixtechblog.com/practical-api-design-at-netflix-part-2-protobuf-fieldmask-for-mutation-operations-2e75e1d230e4 (part 1 and part 2)
2. gRPC: Up and Running by Kasun Indrasiri, Danesh Kuruppu
3. https://adityagoel123.medium.com/in-conversation-with-googles-grpc-part1-f33521f313f googles implementation and notes on scalability
4. https://practical.li/clojure-web-services/reference/ring/ (Some web specifications for ring the CMR’s middleware service)
5. https://grpc.io/docs/guides/benchmarking/ grpc page and benchmarking
6. https://blog.postman.com/postman-now-supports-grpc/ postman supports invoking grpc methods so you could use it for some testing. Insomnia does as well at least according to https://docs.insomnia.rest/insomnia/grpc


Presentation:

* Slide introducing gRPC what it is
* Slide with the community defined pros and cons
* Slide with the architecture diagram being used by this demo project
* Demo with explanation (what was the goal. The goal was to have a basic application that could demonstrate simple CRUD behavior using the GPRC protocol)
* Closing slide and initial dev experience + some references

