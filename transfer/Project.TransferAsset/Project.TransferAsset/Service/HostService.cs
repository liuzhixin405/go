namespace Project.TransferAsset.Service;

public class HostService : BackgroundService
{
    //private readonly IAssetTransferRecordBusiness _assetTransferRecordBusiness;
    private readonly ILogger<HostService> _logger;
    public HostService(ILogger<HostService> logger)
    {
        //_assetTransferRecordBusiness = assetTransferRecordBusiness;
        _logger = logger;
    }
    protected override Task ExecuteAsync(CancellationToken stoppingToken)
    {
        _logger.LogDebug("hostservice starting...");
        JobHelper.SetIntervalJob(async () =>
        {
            try
            {
                using var scope = GlobalConfigure.ServiceLocatorInstance.CreateAsyncScope();
                   var bus =  scope.ServiceProvider.GetRequiredService<IAssetTransferRecordBusiness>();
                  await  bus.BatchConfirmTransfer();
                //await _assetTransferRecordBusiness.BatchConfirmTransfer();
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, "批量确认资产划转失败");
            }
        }, TimeSpan.FromSeconds(15));
        return Task.CompletedTask;
    }
}
