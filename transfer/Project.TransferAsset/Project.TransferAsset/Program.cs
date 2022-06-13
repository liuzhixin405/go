
using MediatR;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddDbContext<TransferContext>();
builder.Services.AddTransient<IAssetBusiness,AssetBusiness>();
builder.Services.AddTransient<IAssetTransferRecordBusiness,AssetTransferRecordBusiness>();
builder.Services.AddTransient<IAssetWasteBookBusiness,AssetWasteBookBusiness>();
builder.Services.AddTransient(typeof(ThirdInvokeBusiness));
builder.Services.AddHostedService<HostService>();
builder.Services.AddMediatR(typeof(Program).Assembly);

builder.Services.AddControllers();

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

using(var scope = builder.Services.BuildServiceProvider().CreateScope())
{
        var services = scope.ServiceProvider;
        try
        {
            var context = services.GetRequiredService<TransferContext>();
            DbdataInitializer.Initialize(context);
        }
        catch (Exception ex)
        {
            var logger = services.GetRequiredService<ILogger<Program>>(); 
            logger.LogError(ex, "An error occurred while seeding the database.");
        }
}

var app = builder.Build();

if(GlobalConfigure.ServiceLocatorInstance==null)
GlobalConfigure.ServiceLocatorInstance = app.Services;
// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();

/*
dotnet ef migrations add InitialCreate
dotnet ef database update
*/