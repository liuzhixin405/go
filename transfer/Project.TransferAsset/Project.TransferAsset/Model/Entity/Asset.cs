using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;

namespace Project.TransferAsset.Model.Entity;

/// <summary>
    /// 资产
    /// </summary>
    [Table("Asset")]
    public class Asset
    {

        /// <summary>
        /// ??
        /// </summary>
        [Key, Column(Order = 1)]
        public String Id { get; set; }

        /// <summary>
        /// ????
        /// </summary>
        [Column(TypeName = "datetime")]
        public DateTime CreateTime { get; set; }

        /// <summary>
        /// ???Id
        /// </summary>
        public String CreatorId { get; set; }

        /// <summary>
        /// ????
        /// </summary>
        public Boolean Deleted { get; set; }

        /// <summary>
        /// CustomerId
        /// </summary>
        public String CustomerId { get; set; }

        /// <summary>
        /// BusinessId
        /// </summary>
        public String BusinessId { get; set; }

        /// <summary>
        /// CoinId
        /// </summary>
        public String CoinId { get; set; }

        /// <summary>
        /// 可用
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal AvailableQuantity { get; set; }

        /// <summary>
        /// 占用
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal FrozenQuantity { get; set; }
        /// <summary>
        /// 划入
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public decimal Include { get; set; }
        /// <summary>
        /// 划出
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public decimal Drawout { get; set; }
    }