using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;

namespace Project.TransferAsset.Model.Entity;

    /// <summary>
    /// 转账记录
    /// </summary>
    [Table("AssetTransferRecord")]
    public class AssetTransferRecord
    {

        /// <summary>
        /// 主键
        /// </summary>
        [Key, Column(Order = 1)]
        public String Id { get; set; }

        /// <summary>
        /// ????
        /// </summary>
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
        /// CoinID
        /// </summary>
        public String CoinID { get; set; }

        /// <summary>
        /// Direction
        /// </summary>
        public TransferDirection Direction { get; set; }

        /// <summary>
        /// Amount
        /// </summary>
        [Column(TypeName = "decimal(28,16)")]
        public Decimal Amount { get; set; }

        /// <summary>
        /// Status
        /// </summary>
        public TransferStatus Status { get; set; }

        /// <summary>
        /// BusinessId
        /// </summary>
        public String BusinessId { get; set; }

    
    }
