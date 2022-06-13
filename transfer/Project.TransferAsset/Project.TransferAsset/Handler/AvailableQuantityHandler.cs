using MediatR;

namespace Project.TransferAsset.Handler
{
    public record AvailableQuantityReq(string businessId, string customerId, string coinId) : IRequest<decimal> { }

    public class AvailableQuantityHandler : IRequestHandler<AvailableQuantityReq, decimal>
    {
        private readonly IAssetBusiness _assetBusiness;
        public AvailableQuantityHandler(IAssetBusiness assetBusiness)
        {
            _assetBusiness = assetBusiness;
        }
    
        public async Task<decimal> Handle(AvailableQuantityReq request, CancellationToken cancellationToken)
        {
            var result = await _assetBusiness.GetAvailableQuantity(request.businessId,request.customerId,request.coinId);
            return result;
        }
    }


}
