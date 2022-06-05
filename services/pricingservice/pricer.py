from common_pb2 import RideType, DEFAULT, COMFORT, PREMIUM


_PRICE_PER_DIST = {
    DEFAULT: 2.0,
    COMFORT: 3.5,
    PREMIUM: 7.0,
}


def calc_cost(*, ride_type, distance):
    for type in RideType.items():
        if ride_type == type[1]:
            return _PRICE_PER_DIST[ride_type] * distance

    return None
