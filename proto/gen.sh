protoc user.proto --proto_path=./ \
    --go_opt=Muser.proto=../UserSrv/services \
    --go-grpc_opt=Muser.proto=../UserSrv/services \
    --go_out=./ --go-grpc_out=./

protoc user.proto --proto_path=./ \
    --go_opt=Muser.proto=../ApiGateway/grpcClient \
    --go-grpc_opt=Muser.proto=../ApiGateway/grpcClient \
    --go_out=./ --go-grpc_out=./