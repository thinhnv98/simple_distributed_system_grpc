# Two methods gRPC on GO

## Run server A
```
make run-server
```

## Test client B
```
make run-client
```

## Result
- You will see something like this:
```
2021/12/05 22:46:45 Run test client for method: Simple RPC
Succeed for input: ThinhNguyenV

2021/12/05 22:46:45 Run test client for method: Server-side streaming RPC
Received name: Name A
Received name: Name B
Received name: Name C
No more stream

```
- You can change this text ```ThinhNguyenV``` or ```{"Name A", "Name B", "Name C"}``` in ```B/main.go``` to anything you want.