from concurrent import futures
import os

from pricer import calc_cost
from dotenv import load_dotenv
import base_pb2
import base_pb2_grpc
import grpc


load_dotenv()
PORT = os.environ['PORT']


class PricingServicer(base_pb2_grpc.PricingServiceServicer):

    def CalculateCost(self, request, context):
        type = request.type
        dist = request.distance

        cost = calc_cost(ride_type=type, distance=dist)
        if cost is None:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details(
                f'Received invalid ride type: "<type: {type}>"')
            return base_pb2.CalculateCostResponse()
        else:
            return base_pb2.CalculateCostResponse(cost=cost)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor())
    base_pb2_grpc.add_PricingServiceServicer_to_server(
        PricingServicer(),
        server
    )
    server.add_insecure_port(f'[::]:{PORT}')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print(f'starting pricingservice at: {PORT}')
    serve()
