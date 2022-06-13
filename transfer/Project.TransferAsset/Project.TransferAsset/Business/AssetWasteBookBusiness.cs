
namespace Project.TransferAsset.Business;

public class AssetWasteBookBusiness:IAssetWasteBookBusiness
{

    private readonly TransferContext _context; 
    public AssetWasteBookBusiness(TransferContext context)
    {
        _context = context;
    }
}
