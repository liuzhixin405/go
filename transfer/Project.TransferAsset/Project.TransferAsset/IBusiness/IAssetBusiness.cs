
public interface IAssetBusiness
{
     Task<decimal> GetAvailableQuantity(string businessId, string customerId, string coinId);
      Task<Asset> GetCustomerAssets(string businessId, string customerId, string coinId);
      Task<decimal> GetThirdPartyAvailableQuantityAsync(string customerId, string coinId);
      Task<bool> UpdateAssetsAsync(IEnumerable<AssetsChangeItemDTO> assetChangeItems);

}
