namespace Project.TransferAsset.Model;

public class DbdataInitializer
{
    public static void Initialize(TransferContext context)
    {
        context.Database.EnsureCreated();

        if (context.Asset.Count() > 0)
        {
            return;
        }

        Asset asset = new Asset
        {
            Id = Guid.NewGuid().ToString(),
            CreateTime = DateTime.Now,
            Deleted = false,
            CreatorId = "test",
            CustomerId = "test001",
            BusinessId = "1",
            CoinId = "1",
            AvailableQuantity = 100m,
            FrozenQuantity = 0m,
            Drawout = 0m,
            Include = 0m
        };
        context.Asset.Add(asset);
        if (context.SaveChanges() == 0)
        {
            throw new Exception("写入默认数据失败。");
        }
    }
}