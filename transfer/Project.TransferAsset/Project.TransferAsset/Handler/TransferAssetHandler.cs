using MediatR;

namespace Project.TransferAsset.Handler
{

    public record TransferAssetReq( string customerId, string coinId, decimal amount,TransferDirection side) : IRequest { }
    public class TransferAssetHandler : IRequestHandler<TransferAssetReq>
    {
        private readonly IAssetTransferRecordBusiness _assetTransferRecordBusiness;
        public TransferAssetHandler(IAssetTransferRecordBusiness assetTransferRecordBusiness)
        {
            _assetTransferRecordBusiness = assetTransferRecordBusiness;
        }
        public async Task<Unit> Handle(TransferAssetReq request, CancellationToken cancellationToken)
        {
            await _assetTransferRecordBusiness.TransferAssets(request.customerId, request.coinId, request.amount, request.side);
            return new Unit();
        }
    }
}
