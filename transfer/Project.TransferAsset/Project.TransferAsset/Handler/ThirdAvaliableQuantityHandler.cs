using MediatR;

namespace Project.TransferAsset.Handler
{
    public record ThirdAvaliableQuantityReq(string customerId,string coinId) : IRequest<decimal> { }
    public class ThirdAvaliableQuantityHandler : IRequestHandler<ThirdAvaliableQuantityReq, decimal>
    {
        private readonly IAssetBusiness _assetBusiness;
        public ThirdAvaliableQuantityHandler(IAssetBusiness assetBusiness)
        {
            _assetBusiness = assetBusiness;
        }
        public async Task<decimal> Handle(ThirdAvaliableQuantityReq request, CancellationToken cancellationToken)
        {
            return await _assetBusiness.GetThirdPartyAvailableQuantityAsync(request.customerId, request.coinId);
        }
    }
}
