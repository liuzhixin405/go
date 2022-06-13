 namespace Project.TransferAsset.Model;
 
 public class AssetsChangeItemDTO
    { 
        /// <summary>
        /// 商户Id
        /// </summary>
        public string BusinessId { set; get; }
        /// <summary>
        /// 客户ID
        /// </summary>
        public string CustomerId { get; set; }
        /// <summary>
        /// 币种Id
        /// </summary>
        public string CoinId { get; set; } 
        /// <summary>
        /// 冻结资产变动数，+表示增加，-表示扣除
        /// </summary>
        public decimal ChangeFrozenQuantity { get; set; }
        /// <summary>
        /// 可用资产变动数，+表示增加，-表示扣除
        /// </summary>
        public decimal ChangeAvailableQuantity { get; set; }
        /// <summary>
        /// 关联单号ID
        /// </summary>
        public string AssociatedOrderId { get; set; }
        /// <summary>
        /// 资产流水类型
        /// </summary>
        public AssetsWasteBookType AssetsWasteBookType { get; set; } 
        /// <summary>
        /// 备注
        /// </summary>
        public string Remarks { get; set; }
    }