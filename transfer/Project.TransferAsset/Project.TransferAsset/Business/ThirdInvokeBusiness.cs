namespace Project.TransferAsset.Business;

public class ThirdInvokeBusiness        //待实现调用其他的服务
{
    public async Task<bool?> ConfirmTansfer(string orderId,bool? success)
    {
        // http://localhost:8080/v1/transfer/getavailablequantity
        await Task.CompletedTask;
        return true;
    }

    public Task<decimal> GetCoinAvalibaleQuantity(string customerId,string coin)
    {
        // http://localhost:8080/v1/transfer/getavailablequantity
        return Task.FromResult(100.0m);
    }
    
    public async Task<bool> TransferAssets(string token, string orderId, string coin, decimal amount, TransferDirection side){
        // http://localhost:8080/v1/transfer/transferasset
        await Task.CompletedTask;
        return true;
    }

}
