using System.ComponentModel;
namespace Project.TransferAsset.Model;

public enum TransferStatus
{
        /// <summary>
        /// 确认中
        /// </summary>
        [Description("确认中")]
        Confirming=1,
        /// <summary>
        /// 已确认
        /// </summary>
        [Description("已确认")]
        Confirmed=2,
        /// <summary>
        /// 取消
        /// </summary>
        [Description("取消")]
        Undo=3,
        /// <summary>
        /// 错误
        /// </summary>
        [Description("错误")]
        Error=4
}
