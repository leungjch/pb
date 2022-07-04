# Script for generating gRPC stubs from proto files in /proto


yourfilenames=`ls proto/*.proto`
for filename in $yourfilenames
do
    base=`basename "$filename" .proto`
    echo $base
    # Generate Python gRPC stubs
    python -m grpc_tools.protoc -Iproto \
        --python_out=pb/py/pb_leungjch/${base} \
        --grpc_python_out=pb/py/pb_leungjch/${base} \
        $base.proto
    
    # Generate gRPC Gateway stubs
    protoc -Iproto --grpc-gateway_out pb/go/${base} \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        $base.proto

done

# python -m grpc_tools.protoc -Iproto \
#     --python_out=style_transfer/generated \
#     --grpc_python_out=style_transfer/generated \
#     style_transfer.proto

# https://stackoverflow.com/a/69378666
gsed -i 's/^import .*_pb2 as/from . \0/' pb/py/pb_leungjch/*/*_pb2_grpc.py