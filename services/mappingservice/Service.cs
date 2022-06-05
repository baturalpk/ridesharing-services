using Grpc.Core;

namespace mappingservice
{
    public class Service : MappingService.MappingServiceBase
    {
        private const float AverageDrivingSpeed = 18.6f;

        private readonly Random _random;
        public Service(ILogger<Service> logger)
        {
            _random = new Random();
        }

        public override Task<CalculateDistanceResponse> CalculateDistance(CalculateDistanceRequest request, ServerCallContext context)
        {
            var route = request.Route;

            return Task.FromResult(new CalculateDistanceResponse
            {
                Distance = CalculateDistanceBetween(destination: route.Destination, start: route.Start)
            });
        }

        public override Task<DoGeocodingResponse> DoGeocoding(DoGeocodingRequest request, ServerCallContext context)
        {
            var address = request.Address;

            // In reality do some geocoding stuff here...

            var location = new Location();
            location.Latitude = _random.Next(-90, 89) + (float)_random.NextDouble();
            location.Longitude = _random.Next(-180, 179) + (float)_random.NextDouble();

            return Task.FromResult(new DoGeocodingResponse
            {
                Location = location
            });
        }

        public override Task<MakeTimeEstimationResponse> MakeTimeEstimation(MakeTimeEstimationRequest request, ServerCallContext context)
        {
            var route = request.Route;
            float dist = CalculateDistanceBetween(destination: route.Destination, start: route.Start);

            return Task.FromResult(new MakeTimeEstimationResponse
            {
                Eta = dist / AverageDrivingSpeed
            });
        }

        private static float CalculateDistanceBetween(Location destination, Location start)
        {
            return (float)Math.Sqrt(
                Math.Pow(destination.Longitude - start.Longitude, 2) + Math.Pow(destination.Latitude - start.Latitude, 2)
            );
        }
    }
}
