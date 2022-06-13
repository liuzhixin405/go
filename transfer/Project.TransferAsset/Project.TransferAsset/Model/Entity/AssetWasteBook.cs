using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Project.TransferAsset.Model.Entity;

 /// <summary>
    /// 资产流水
    /// </summary>
    [Table("AssetWasteBook")]
    public class AssetWasteBook
    {

        /// <summary>
        /// 主键
        /// </summary>
        [Key, StringLength(50), Column(TypeName = "varchar(50)", Order = 1)]
        public String Id { get; set; }

        /// <summary>
        /// 创建时间
        /// </summary>
        [Column(TypeName = "datetime")]
        public DateTime CreateTime { get; set; }

        /// <summary>
        /// 创建人
        /// </summary>
        [Column(TypeName = "varchar(50)"), StringLength(50)]
        public String CreatorId { get; set; }

        /// <summary>
        /// 是否软删除
        /// </summary>
        public Boolean Deleted { get; set; }

        /// <summary>
        /// 客户Id
        /// </summary>
        [Column(TypeName = "varchar(50)"), StringLength(50)]
        public String CustomerId { get; set; }

        /// <summary>
        /// 商户或联盟Id
        /// </summary>
        [Column(TypeName = "varchar(50)"), StringLength(50)]
        public String BusinessId { get; set; }

        /// <summary>
        /// 币种Id
        /// </summary>
        [Column(TypeName = "varchar(50)"), StringLength(50)]
        public String CoinId { get; set; }

        /// <summary>
        /// 备注
        /// </summary>
        [Column(TypeName = "varchar(500)"), StringLength(500)]
        public String Remarks { get; set; }

        /// <summary>
        /// 关联订单Id
        /// </summary>
        [Column(TypeName = "varchar(200)"), StringLength(200)]
        public String AssociatedOrderId { get; set; }

        /// <summary>
        /// 原始可用数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal OriginalAvailableQuantity { get; set; }

        /// <summary>
        /// 更改数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal ChangeQuantity { get; set; }

        /// <summary>
        /// 可用数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal AvailableQuantity { get; set; }

        /// <summary>
        /// 原始冻结数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal OriginalFrozenQuantity { get; set; }

        /// <summary>
        /// 更改冻结数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal ChangeFrozenQuantity { get; set; }

        /// <summary>
        /// 冻结数量
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal FrozenQuantity { get; set; }
        /// <summary>
        /// 资产流水类型
        /// </summary>
        public AssetsWasteBookType AssetsWasteBookType { set; get; }
    }
