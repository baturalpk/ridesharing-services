using System;

var PORT = Environment.GetEnvironmentVariable("PORT");

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddGrpc();

var app = builder.Build();

app.MapGrpcService<mappingservice.Service>().RequireHost($"*:{PORT}");

app.Run();
