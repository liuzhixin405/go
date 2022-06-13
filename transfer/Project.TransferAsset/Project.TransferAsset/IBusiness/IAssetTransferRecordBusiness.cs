namespace Project.TransferAsset.IBusiness;

public interface IAssetTransferRecordBusiness
{
        Task TransferAssets(string customerId, string coinId, decimal amount, TransferDirection side);

        Task BatchConfirmTransfer();
        /// <summary>
        /// 资产划转状态
        /// </summary>
        /// <returns></returns>
        Task<bool> GetTransferAssetsStatus();
        Task ConfirmTransfer(AssetTransferRecord record);
}