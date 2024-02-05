protoc user.proto --proto_path=./ \
    --go_opt=Muser.proto=../UserSrv/services \
    --go-grpc_opt=Muser.proto=../UserSrv/services \
    --go_out=./ --go-grpc_out=./

protoc user.proto --proto_path=./ \
    --go_opt=Muser.proto=../ApiGateway/grpcClient \
    --go-grpc_opt=Muser.proto=../ApiGateway/grpcClient \
    --go_out=./ --go-grpc_out=./


protoc product.proto --proto_path=./ \
    --go_opt=Mproduct.proto=../ProductSrv/services \
    --go-grpc_opt=Mproduct.proto=../ProductSrv/services \
    --go_out=./ --go-grpc_out=./


protoc product.proto --proto_path=./ \
    --go_opt=Mproduct.proto=../ApiGateway/grpcClient \
    --go-grpc_opt=Mproduct.proto=../ApiGateway/grpcClient \
    --go_out=./ --go-grpc_out=./



protoc order.proto --proto_path=./ \
    --go_opt=Morder.proto=../OrderSrv/services \
    --go-grpc_opt=Morder.proto=../OrderSrv/services \
    --go_out=./ --go-grpc_out=./

protoc order.proto --proto_path=./ \
    --go_opt=Morder.proto=../ApiGateway/grpcClient \
    --go-grpc_opt=Morder.proto=../ApiGateway/grpcClient \
    --go_out=./ --go-grpc_out=./
