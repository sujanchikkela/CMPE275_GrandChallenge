import my_python
import my_python_2
import time
import grpc

def get_client_stream_requests():
    while True:
        name = input("")

        if name == "":
            break

        hello_request = my_python_2.HelloRequest(greeting = "Hello", name = name)
        yield hello_request
        time.sleep(1)

def run():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = my_python.GreeterStub(channel)
        rpc_call = input("")

        if rpc_call == "1":
            hello_request = my_python_2.HelloRequest(greeting = "", name = "")
            hello_reply = stub.SayHello(hello_request)
            print(hello_reply)
        elif rpc_call == "2":
            hello_request = my_python_2.HelloRequest(greeting = "", name = "")
            hello_replies = stub.ParrotSaysHello(hello_request)

            for hello_reply in hello_replies:
                print(" Response Received:")
                print(hello_reply)
        elif rpc_call == "3":
            delayed_reply = stub.ChattyClientSaysHello(get_client_stream_requests())

            print("Response Received:")
            print(delayed_reply)
        elif rpc_call == "4":
            responses = stub.InteractingHello(get_client_stream_requests())

            for response in responses:
                print("Response Received: ")
                print(response)

if __name__ == "__main__":
    run()