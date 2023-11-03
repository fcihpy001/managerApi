package model

import (
	"gorm.io/gorm"
	"time"
)

// 合约信息
type Contract struct {
	gorm.Model
	Address string
	Total   uint64
	Minted  uint64
	Name    string
	Desc    string
}

// 链上交易数据
type Transaction struct {
	gorm.Model
	Address  string
	ChainId  uint
	TxHash   string
	From     string
	To       string
	Success  bool
	Event    string
	Function string
	Value    string
	Nonce    uint64
	Symbol   string
}

// 钱包用户数据
type Wallet struct {
	gorm.Model
	Address            string         `json:"address" gorm:"uniqueIndex;type:char(42)"`
	USDTAmount         uint           `json:"usdt_amount"`
	TPLAmount          string         `json:"tpl_amount"`
	Referrer           string         `json:"referrer"`
	InviteAmount       uint           `json:"invite_amount"`
	BindTime           string         `json:"bind_time"`
	BindSource         BindSourceType `json:"bind_source"`
	CommunityNFTAmount uint           `json:"community_nft_amount"`
	MedalNFTAmount     uint           `json:"medal_nft_amount"`
	WithdrawAmount     uint           `json:"withdraw_amount"`
}

type BindSourceType string

// 邀请绑定来源:链接，7人团购合约
const (
	BindSourceTypeLink     BindSourceType = "link"
	BindSourceTypeContract BindSourceType = "contract"
)

// 链
type Chain struct {
	gorm.Model
	ChainId   uint
	ChainName string
	IsTest    bool
	RPC       string
	WebSocket string
	BlockScan string
}

// 代币
type Coin struct {
	gorm.Model
	CoinType string `form:"coin_type" binding:"required"`
	Symbol   string `form:"symbol" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Index    uint   `form:"index" binding:"required"`
	Address  string `form:"address" binding:"required"`
	Remark   string `form:"remark"`
	Did      string
}

// 排名
type Rank struct {
	Type    RankType
	Address string `json:"address"`
	Amount  string `json:"amount"`
}

type RankType string

const (
	RankTypeInvite    RankType = "invite"
	RankTypeUSDT      RankType = "usdt"
	RankTypeCommunity RankType = "community"
	RankTypeMedal     RankType = "medal"
)

type RankRequest struct {
	PageNum  int      `form:"page_num"`
	PageSize int      `form:"page_size"`
	RankType RankType `form:"rank_type"`
}

type ListRequest struct {
	Wallet      string      `form:"wallet"`
	PageRequest PageRequest `gorm:"embedded"`
	StartDate   string      `form:"start_date"`
	EndDate     string      `form:"end_date"`
}

// 奖励
type Reward struct {
	gorm.Model
	RewardTime string     `json:"reward_time"`
	TxHash     string     `gorm:"tx_hash;uniqueIndex; type:char(50);"`
	RewardType RewardType `json:"reward_type" form:"reward_type" binding:"required"`
	Wallet     string     `json:"wallet" form:"wallet" binding:"required"`
	NFTType    NFTType    `json:"nft_type" form:"nft_type" `
	NFTID      uint       `json:"nft_id" form:"nft_id"`
	USDTAmount string     `json:"usdt_amount" form:"usdt_amount"`
	Status     int        `json:"status"  form:"status" `
	Remark     string     `json:"remark"  form:"remark" gorm:"default:1"`
	Source     string     `json:"source"  form:"source" `
}

type RewardType string

const (
	RewardTypeUSDT         RewardType = "USDT"
	RewardTypeCommunityNFT RewardType = "CommunityNFT"
	RewardTypeMedalNFT     RewardType = "MedaNFT"
)

type USDT struct {
	gorm.Model
	Source string `form:"source" binding:"required"`
	From   string `form:"from" binding:"required"`
	To     string `form:"to" binding:"required"`
	Amount uint   `form:"amount" binding:"required"`
	Status int    `form:"status"`
	Remark string `form:"remark"`
}

// nft
type NFT struct {
	gorm.Model
	Type   NFTType `gorm:"uniqueIndex:type_id_idx; type:char(10)"`
	NFTID  uint    `gorm:"uniqueIndex:type_id_idx"`
	From   string  `gorm:"from"`
	To     string  `gorm:"to"`
	Source string  `gorm:"source"`
	Status int     `gorm:"default:1"`
	Remark string  `gorm:"remark"`
}
type NFTType string

const (
	//会员，伙伴、创世、
	NFTTypeMember    NFTType = "member"
	NFTTypeFellow    NFTType = "fellow"
	NFTTypeCommunity NFTType = "community"
	NFTTypeMedal     NFTType = "medal"
)

type Deposit struct {
	gorm.Model
	From   string      `json:"from" form:"from" binding:"required"`
	To     string      `json:"to" form:"to" binding:"required"`
	Type   DepositType `json:"type" form:"type" binding:"required"`
	Amount uint        `json:"amount" form:"amount" binding:"required"`
	Profit uint        `json:"profit" form:"profit" binding:"required"`
}

type DepositType string

const (
	DepositTypeTPL          DepositType = "tpl"
	DepositTypeCommunityNFT DepositType = "community"
	DepositTypeMedalNFT     DepositType = "medal"
)

type DepositRequest struct {
	Wallet      string      `form:"wallet" binding:"required"`
	Type        DepositType `form:"type" binding:"required"`
	PageRequest PageRequest `gorm:"embedded"`
}

type Token struct {
	gorm.Model
	Type   TokenType
	TxHash string `gorm:"tx_hash;uniqueIndex; type:char(50);"`
	Source string `form:"source" binding:"required"`
	From   string `form:"from" binding:"required"`
	To     string `form:"to" binding:"required"`
	Amount uint   `form:"amount" binding:"required"`
	Status int    `form:"status" gorm:"default:1"`
	Remark string `form:"remark"`
}

type TokenType string

const (
	TokenTypeUSDT TokenType = "usdt"
	TokenTypeTPL  TokenType = "tpl"
)

type WithdrawResp struct {
	CreatedAt time.Time `json:"created_at"`
	To        string    `json:"wallet"`
	Amount    string    `json:"amount"`
	BlockTime string    `json:"block_time"`
}
