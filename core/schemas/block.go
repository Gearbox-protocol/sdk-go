package schemas

type (
	Block struct {
		BlockNumber             int64                     `gorm:"primaryKey;column:id" json:"blockNum"` // Block Number
		Timestamp               uint64                    `gorm:"column:timestamp" json:"timestamp"`
		AccountOperations       []*AccountOperation       `gorm:"foreignKey:block_num" json:"accountOperations"`
		TokenOracles            []*TokenOracle            `gorm:"foreignKey:block_num" json:"tokenOracles"`
		PriceFeeds              []*PriceFeed              `gorm:"foreignKey:block_num" json:"priceFeeds"`
		Protocols               []*Protocol               `gorm:"foreignKey:block_num" json:"protocols"`
		CSS                     []*CreditSessionSnapshot  `gorm:"foreignKey:block_num" json:"css"`
		PoolStats               []*PoolStat               `gorm:"foreignKey:block_num" json:"poolStats"`
		PoolLedgers             []*PoolLedger             `gorm:"foreignKey:block_num" json:"poolLedgers"`
		CMStats                 []*CreditManagerStat      `gorm:"foreignKey:block_num" json:"cmStats"`
		AllowedTokens           []*AllowedToken           `gorm:"foreignKey:block_num" json:"allowedTokens"`
		DAOOperations           []*DAOOperation           `gorm:"foreignKey:block_num" json:"daoOperations"`
		Params                  []*Parameters             `gorm:"foreignKey:block_num" json:"params"`
		FastCheckParams         []*FastCheckParams        `gorm:"foreignKey:block_num" json:"fastCheckParams"`
		TreasuryTransfers       []*TreasuryTransfer       `gorm:"foreignKey:block_num" json:"treasuryTransfers"`
		TreasurySnapshots       []*TreasurySnapshotModel2 `gorm:"foreignKey:block_num" json:"treasurySnapshots"`
		NoSessionTokenTransfers []*TokenTransfer          `gorm:"foreignKey:block_num" json:"noSessionTokenTransfers"`
		TAA                     []*TransferAccountAllowed `gorm:"foreignKey:block_num" json:"transferAccountAllowed"`
		DieselTransfers         []*DieselTransfer         `gorm:"foreignKey:block_num" json:"dieselTransfers"`
	}
)

func (Block) TableName() string {
	return "blocks"
}

func (b *Block) AddAccountOperation(accountOperation *AccountOperation) {
	b.AccountOperations = append(b.AccountOperations, accountOperation)
}
func (b *Block) AddDieselTransfer(transfer *DieselTransfer) {
	b.DieselTransfers = append(b.DieselTransfers, transfer)
}
func (b *Block) AddTokenOracle(tokenOracle *TokenOracle) {
	b.TokenOracles = append(b.TokenOracles, tokenOracle)
}
func (b *Block) AddPriceFeed(pf *PriceFeed) {
	b.PriceFeeds = append(b.PriceFeeds, pf)
}
func (b *Block) AddAllowedProtocol(p *Protocol) {
	b.Protocols = append(b.Protocols, p)
}

func (b *Block) AddAllowedToken(atoken *AllowedToken) {
	b.AllowedTokens = append(b.AllowedTokens, atoken)
}

func (b *Block) AddDAOOperation(operation *DAOOperation) {
	b.DAOOperations = append(b.DAOOperations, operation)
}

func (b *Block) AddCreditSessionSnapshot(css *CreditSessionSnapshot) {
	b.CSS = append(b.CSS, css)
}

func (b *Block) AddPoolStat(ps *PoolStat) {
	b.PoolStats = append(b.PoolStats, ps)
}
func (b *Block) AddPoolLedger(pl *PoolLedger) {
	b.PoolLedgers = append(b.PoolLedgers, pl)
}

func (b *Block) AddCreditManagerStats(cms *CreditManagerStat) {
	b.CMStats = append(b.CMStats, cms)
}

func (b *Block) AddTreasuryTransfer(tt *TreasuryTransfer) {
	b.TreasuryTransfers = append(b.TreasuryTransfers, tt)
}

func (b *Block) AddTreasurySnapshot(tss *TreasurySnapshot) {
	b.TreasurySnapshots = append(b.TreasurySnapshots, &TreasurySnapshotModel2{
		BlockNum:              tss.BlockNum,
		Date:                  tss.Date,
		PricesInUSD:           tss.PricesInUSD,
		ValueInUSD:            tss.ValueInUSD,
		OperationalValueInUSD: tss.OperationalValueInUSD,
		Balances:              tss.Balances,
		OperationalBalances:   tss.OperationalBalances,
	})
}

func (b *Block) GetAllowedTokens() []*AllowedToken {
	return b.AllowedTokens
}

func (b *Block) GetPriceFeeds() []*PriceFeed {
	return b.PriceFeeds
}

func (b *Block) AddParameters(params *Parameters) {
	lenOfParams := len(b.Params)
	if lenOfParams > 0 {
		// if there are two params for same creditmanager if same block use the last one.
		if b.Params[lenOfParams-1].CreditManager == params.CreditManager {
			b.Params[lenOfParams-1] = params
			return
		}
	}
	b.Params = append(b.Params, params)
}

func (b *Block) AddFastCheckParams(params *FastCheckParams) {
	b.FastCheckParams = append(b.FastCheckParams, params)
}

func (b *Block) AddNoSessionTx(tt *TokenTransfer) {
	b.NoSessionTokenTransfers = append(b.NoSessionTokenTransfers, tt)
}

func (b *Block) GetCSS() []*CreditSessionSnapshot {
	return b.CSS
}

func (b *Block) GetPoolStats() []*PoolStat {
	return b.PoolStats
}

func (b *Block) GetParams() []*Parameters {
	return b.Params
}

func (b *Block) AddTransferAccountAllowed(obj *TransferAccountAllowed) {
	b.TAA = append(b.TAA, obj)
}

type DieselTransfer struct {
	LogId       int64   `gorm:"primaryKey;column:log_id"`
	BlockNum    int64   `gorm:"primaryKey;column:block_num"`
	TokenSymbol string  `gorm:"column:token_sym"`
	From        string  `gorm:"column:from_user"`
	To          string  `gorm:"column:to_user"`
	Amount      float64 `gorm:"column:amount"`
}

func (DieselTransfer) TableName() string {
	return "diesel_transfers"
}
