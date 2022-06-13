
using MediatR;
using Project.TransferAsset.Handler;

namespace Project.TransferAsset.Controllers;

[ApiController]
[Route("/api/[controller]/[action]")]
public class TransferController:ControllerBase
{
    private readonly IMediator _mediator;
    public TransferController(IMediator mediator)
    {
        _mediator = mediator;
    }

        /// <summary>
        /// 资产划转
        /// </summary>
        /// <param name="coinId">币种ID</param>
        /// <param name="amount">数量</param>
        /// <param name="side">方向 1、转入  2、转出</param>
        /// <returns></returns>
        [HttpPost]
        public async Task TransferAssetsAsync([Required] string customerId,[Required] string coinId, [Required] decimal amount, [Required]TransferDirection side)
        {
        await _mediator.Send(new TransferAssetReq(customerId,coinId,amount,side));
        }
        /// <summary>
        /// 获取用户在合约平台上币种的可用数量
        /// </summary>
        /// <param name="coinId">币种Id</param>
        /// <returns></returns>
        [HttpPost]
        public async Task<decimal> GetAvailableQuantityAsync([Required] string businessId, [Required] string customerId,[Required] string coinId)
        {
        return await _mediator.Send(new AvailableQuantityReq(businessId,customerId,coinId));
        }
        /// <summary>
        /// 获取第三方平台上的币种可用数量
        /// </summary> 
        /// <param name="coinId">币种Id</param>
        /// <returns></returns>
        [HttpPost]
        public async Task<decimal> GetThirdPartyAvailableQuantityAsync([Required]string customerId,[Required] string coinId)
        {
        return await _mediator.Send(new ThirdAvaliableQuantityReq(customerId, coinId));
        }
}
