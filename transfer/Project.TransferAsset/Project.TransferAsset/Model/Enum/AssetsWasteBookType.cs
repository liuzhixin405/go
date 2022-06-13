using System.ComponentModel;

namespace Project.TransferAsset.Model;

public enum AssetsWasteBookType
{
    /// <summary>
        /// 资产转出
        /// </summary>
        [Description("资产划转出")]
        TransferOut = 1,
        /// <summary>
        /// 资产转入
        /// </summary>
        [Description("资产划转入")]
        TransferIn = 2
}
