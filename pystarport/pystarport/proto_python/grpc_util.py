import json

import grpc

import cosmos.bank.v1beta1.tx_pb2
import cosmos.bank.v1beta1.tx_pb2_grpc
import cosmos.crypto.ed25519.keys_pb2
import cosmos.staking.v1beta1.query_pb2
import cosmos.staking.v1beta1.query_pb2_grpc


class GrpcUtil:
    def __init__(self, ip_port):
        self.ip_port = ip_port

    def get_validators(self):
        channel = grpc.insecure_channel(self.ip_port)
        stub = cosmos.staking.v1beta1.query_pb2_grpc.QueryStub(channel)
        response = stub.Validators(
            cosmos.staking.v1beta1.query_pb2.QueryValidatorsRequest()
        )
        return response
