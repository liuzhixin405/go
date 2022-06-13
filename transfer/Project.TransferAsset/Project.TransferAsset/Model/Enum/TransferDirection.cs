using System.ComponentModel;

namespace Project.TransferAsset.Model;

public enum TransferDirection
{
    [Description("转入")]
    In=1 ,
    [Description("转出")]
    Out=2
}
