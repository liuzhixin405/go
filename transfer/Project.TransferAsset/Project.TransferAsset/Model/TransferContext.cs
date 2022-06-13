using Microsoft.EntityFrameworkCore;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.AspNetCore.Builder;
using Project.TransferAsset.Model.Entity;

namespace Project.TransferAsset.Model;

public class TransferContext:DbContext
{
    public TransferContext(DbContextOptions<TransferContext> options):base(options)
    {
        
    }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {
        if(!optionsBuilder.IsConfigured)
        //optionsBuilder.UseMySql("server=8.142.71.127:3306;uid=root;pwd=123456@;database=Account", ServerVersion.Parse("8.0.27-mysql"),builder=>{
            optionsBuilder.UseMySql("server=localhost;user=root;password=123456@;database=account", new MySqlServerVersion(new Version(8,0,13)),builder=>{  //本地localhost  或者127.0.0.1 默认自带端口不能再加3306
             builder.EnableRetryOnFailure(
                maxRetryCount: 5,
                maxRetryDelay: TimeSpan.FromSeconds(30),
                errorNumbersToAdd: null);
        });
        
    }

    public DbSet<Asset> Asset{get;set;}
    public DbSet<AssetWasteBook> AssetWasteBook{get;set;}
    public DbSet<AssetTransferRecord> AssetTransferRecord{get;set;}

}
