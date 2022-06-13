using Project.TransferAsset.IBusiness;
using Project.TransferAsset.Model.Entity;
namespace Project.TransferAsset.Business;

public class AssetBusiness : IAssetBusiness
{
    private readonly TransferContext _context;
    private readonly ThirdInvokeBusiness _thirdInvokeBusiness;
    public AssetBusiness( TransferContext context,ThirdInvokeBusiness thirdInvokeBusiness)
    {
        _context = context;
        _thirdInvokeBusiness =thirdInvokeBusiness;
    }
    public Task<decimal> GetAvailableQuantity(string businessId, string customerId, string coinId)
    {
       var query = _context.Asset.Where(x=>x.BusinessId==businessId&&x.CustomerId==customerId&&x.CoinId==coinId).Select(x=>x.AvailableQuantity).FirstOrDefault();
       return Task.FromResult(query);
    }
    public Task<Asset> GetCustomerAssets(string businessId, string customerId, string coinId)
    {
        var query = _context.Asset.Where(x=>x.BusinessId==businessId&&x.CustomerId==customerId&&x.CoinId==coinId).FirstOrDefault();
        return Task.FromResult(query?? new Asset());
    }
    public async Task<decimal> GetThirdPartyAvailableQuantityAsync(string customerId, string coinId)
    {
        var request = await _thirdInvokeBusiness.GetCoinAvalibaleQuantity(customerId,coinId);
        return request;
    }

    public async Task<bool> UpdateAssetsAsync(IEnumerable<AssetsChangeItemDTO> assetChangeItems)
    {
        var assets = _context.Asset.ToList();
        var newAssets = new List<Asset>();
        var assetWasteBooks = new List<AssetWasteBook>();
        foreach(var assetChangeItem in assetChangeItems)
        {
            var asset = assets.Where(x=>x.BusinessId == assetChangeItem.BusinessId&&x.CustomerId==assetChangeItem.CustomerId&&x.CoinId==assetChangeItem.CoinId).FirstOrDefault();
            if(asset==null)
            {
                throw new Exception($"用户{assetChangeItem.BusinessId}_{assetChangeItem.CustomerId}_{assetChangeItem.CoinId} 没有可用资产");
            }
            assetWasteBooks.Add(CreateWasteBook(asset,assetChangeItem));
        }
       await _context.AssetWasteBook.AddRangeAsync(assetWasteBooks);
       _context.SaveChanges();
       return true;
    }

    private AssetWasteBook CreateWasteBook(Asset assetsItem,AssetsChangeItemDTO assetsChangeItem)
    {
          var assetWasteBook = new AssetWasteBook()
            {
                Id = Guid.NewGuid().ToString(),
                CreateTime = DateTime.Now,
                BusinessId = assetsItem.BusinessId,
                CoinId = assetsItem.CoinId,
                CustomerId = assetsItem.CustomerId,
                CreatorId = assetsItem.CreatorId,
                Deleted = false,
                Remarks = assetsChangeItem.Remarks,
                OriginalAvailableQuantity = assetsItem.AvailableQuantity,
                OriginalFrozenQuantity = assetsItem.FrozenQuantity,
                ChangeFrozenQuantity = assetsChangeItem.ChangeFrozenQuantity,
                ChangeQuantity = assetsChangeItem.ChangeAvailableQuantity,
                AssetsWasteBookType = assetsChangeItem.AssetsWasteBookType,
                AssociatedOrderId = assetsChangeItem.AssociatedOrderId,
                AvailableQuantity = assetsItem.AvailableQuantity + assetsChangeItem.ChangeAvailableQuantity,
                FrozenQuantity = assetsItem.FrozenQuantity + assetsChangeItem.ChangeFrozenQuantity
            };
            assetsItem.AvailableQuantity += assetsChangeItem.ChangeAvailableQuantity;
            assetsItem.FrozenQuantity += assetsChangeItem.ChangeFrozenQuantity;
            if (assetsChangeItem.AssetsWasteBookType == AssetsWasteBookType.TransferIn)
                assetsItem.Include += Math.Abs(assetsChangeItem.ChangeFrozenQuantity);
            else if (assetsChangeItem.AssetsWasteBookType == AssetsWasteBookType.TransferOut)
                assetsItem.Drawout += Math.Abs(assetsChangeItem.ChangeFrozenQuantity);
            return assetWasteBook;
    }
}
