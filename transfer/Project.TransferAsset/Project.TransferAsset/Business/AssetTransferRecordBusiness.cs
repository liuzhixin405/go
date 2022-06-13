
namespace Project.TransferAsset.Business;

public class AssetTransferRecordBusiness : IAssetTransferRecordBusiness
{

    private readonly static ActionBlock<AssetTransferRecord> _assetTransferRecordActionBlock;
    private readonly static BufferBlock<AssetTransferRecord> _assetTransferRecordBufferBlock;
    private readonly IConfiguration _configuration;
    private readonly TransferContext _context;
    private readonly ThirdInvokeBusiness _thirdInvokeBusiness;
    private readonly IAssetBusiness _assetBusiness;
    static AssetTransferRecordBusiness()
    {
        _assetTransferRecordBufferBlock = new BufferBlock<AssetTransferRecord>(new DataflowBlockOptions { BoundedCapacity = DataflowBlockOptions.Unbounded });
        _assetTransferRecordActionBlock = new ActionBlock<AssetTransferRecord>(record =>
        {
            JobHelper.SetDelayJob(() =>
            {
                using var scope = GlobalConfigure.ServiceLocatorInstance.CreateScope();
                var thisBus = scope.ServiceProvider.GetRequiredService<IAssetTransferRecordBusiness>();
                thisBus.ConfirmTransfer(record).ConfigureAwait(false).GetAwaiter().GetResult();

            }, TimeSpan.FromSeconds(1));
        });
    }
    public AssetTransferRecordBusiness(IConfiguration configuration, TransferContext context, ThirdInvokeBusiness thirdInvokeBusiness, IAssetBusiness assetBusiness)
    {
        _configuration = configuration;
        _context = context;
        _thirdInvokeBusiness = thirdInvokeBusiness;
        _assetBusiness = assetBusiness;


    }
    public async Task TransferAssets(string customerId, string coinId, decimal amount, TransferDirection side)
    {
        //省略参数校验
        var isOpenAssetTransfer = _configuration.GetSection("IsOpenAssetTransfer").Get<bool>();
        if (!isOpenAssetTransfer)
            throw new Exception("资产划转功能未开启");
        if (side == TransferDirection.Out)
        {
            var availableQuantity = _context.Asset.Where(x => x.CustomerId == customerId && x.CoinId == coinId).Select(x => x.AvailableQuantity).FirstOrDefault();
            if (availableQuantity < amount)
                throw new Exception($"可用资产为{availableQuantity},超过划转金额{amount}");
        }
        else
        {
            var availableQuantity = await _thirdInvokeBusiness.GetCoinAvalibaleQuantity(customerId, coinId);
            if (availableQuantity < amount)
                throw new Exception($"第三方可用资产为{availableQuantity},超过划转金额{amount}");
        }

        var order = new AssetTransferRecord
        {
            Id = Guid.NewGuid().ToString(),
            Status = TransferStatus.Confirming,
            Amount = amount,
            BusinessId = "1",
            CoinID = coinId,
            CreateTime = DateTime.Now,
            CreatorId = "1",
            Direction = side,
            CustomerId = customerId
        };

        await _context.AssetTransferRecord.AddAsync(order);
        await _context.SaveChangesAsync();

        var isSuccess = await _thirdInvokeBusiness.TransferAssets("token", order.Id, coinId, amount, side);
        if (!isSuccess)
        {
            order.Status = TransferStatus.Error;
            _context.AssetTransferRecord.Update(order);
            _context.SaveChanges();
            throw new Exception("asset trnasfer failed. ");
        }

        var assetsChangeItem = new AssetsChangeItemDTO()
        {
            AssetsWasteBookType = side == TransferDirection.Out ? AssetsWasteBookType.TransferOut : AssetsWasteBookType.TransferIn,
            BusinessId = "1",
            CoinId = coinId,
            CustomerId = customerId,
            AssociatedOrderId = order.Id,
            ChangeAvailableQuantity = (side == TransferDirection.Out) ? -1 * amount : 0 * amount,
            ChangeFrozenQuantity = amount,
            Remarks = "资产划转--冻结"
        };
        await _assetBusiness.UpdateAssetsAsync(new List<AssetsChangeItemDTO>() { assetsChangeItem });
        await _assetTransferRecordBufferBlock.SendAsync(order);
    }

    public Task BatchConfirmTransfer()
    {
        var records = _context.AssetTransferRecord.Where(x => x.Status == TransferStatus.Confirming).ToList();
        records.ForEach(async record =>
        {
            await ConfirmTransfer(record);
        });
        return Task.CompletedTask;
    }
    /// <summary>
    /// 资产划转状态
    /// </summary>
    /// <returns></returns>
    public async Task<bool> GetTransferAssetsStatus()
    {
        await Task.CompletedTask;
        return _configuration.GetSection("IsOpenAssetTransfer").Get<bool>();
    }
    public async Task ConfirmTransfer(AssetTransferRecord record)
    {
        var assetTransferRecord = _context.AssetTransferRecord.FirstOrDefault(x => x.Id == record.Id);
        if (record.Status != assetTransferRecord?.Status) return;
        var isSuccess = await _thirdInvokeBusiness.ConfirmTansfer(record.Id, true);
        if (isSuccess == null)
        {
            await _assetTransferRecordBufferBlock.SendAsync(record);
            return;
        }
        if (!(isSuccess ?? false))
        {
            record.Status = TransferStatus.Error;
            _context.AssetTransferRecord.Update(record);
            throw new Exception("Asset confirm transfer failed.");
        }

        var assetsChangeItem = new AssetsChangeItemDTO()
        {
            AssetsWasteBookType = ((record.Direction == TransferDirection.Out) ? AssetsWasteBookType.TransferOut : AssetsWasteBookType.TransferIn),
            BusinessId = record.BusinessId,
            CoinId = record.CoinID,
            CustomerId = record.CustomerId,
            AssociatedOrderId = record.Id,
            ChangeAvailableQuantity = (record.Direction == TransferDirection.In) ? record.Amount : 0,
            ChangeFrozenQuantity = -record.Amount,
            Remarks = "资产划转--确认"
        };

        await _assetBusiness.UpdateAssetsAsync(new List<AssetsChangeItemDTO> { assetsChangeItem });
        record.Status = TransferStatus.Confirmed;
        _context.AssetTransferRecord.Update(record);
    }
}
