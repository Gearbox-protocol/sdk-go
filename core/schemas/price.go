package schemas

import (
	"fmt"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
)

type PriceOracleT string
type TokenOracle struct {
	PriceOracle PriceOracleT     `gorm:"primaryKey;column:price_oracle" json:"priceOracle"`
	BlockNumber int64            `gorm:"primaryKey;column:block_num" json:"blockNum"`
	Token       string           `gorm:"primaryKey;column:token" json:"token"`
	Oracle      string           `gorm:"column:oracle" json:"oracle"`
	Feed        string           `gorm:"column:feed" json:"feed"`
	DisabledAt  int64            `gorm:"column:disabled_at"`
	Dirty bool `gorm:"-"`
	Version     core.VersionType `gorm:"column:version" json:"version"`
	Reserve     bool             `gorm:"primaryKey;column:reserve" json:"reserve"`
	FeedType    string           `gorm:"feed_type" json:"-"`
	Underlyings []string         `gorm:"-"`
}

func (t TokenOracle) String() string {
	return fmt.Sprintf("token %s of %s with feed %s and version %d, %v added at %d",
		t.Token, t.FeedType,
		t.Feed, t.Version, t.Reserve, t.BlockNumber)
}

func (TokenOracle) TableName() string {
	return "token_oracle"
}

type PriceFeed struct {
	ID int64 `gorm:"column:id;autoIncrement:true" json:"-"`
	//
	BlockNumber int64 `gorm:"primaryKey;column:block_num" json:"blockNum"`
	// Token           string          `gorm:"primaryKey;column:token" json:"token"`
	// MergedPFVersion MergedPFVersion `gorm:"primaryKey;column:merged_pf_version" json:"mergedPFVersion"`
	Feed string `gorm:"column:feed" json:"feed"`
	//
	RoundId int64        `gorm:"column:round_id" json:"roundId"`
	PriceBI *core.BigInt `gorm:"column:price_bi" json:"priceBI"`
	Price   float64      `gorm:"column:price" json:"price"`
}

func (pf PriceFeed) String() string {
	return fmt.Sprintf("Feed(%s) at block %d with Price %f",
		pf.Feed,
		pf.BlockNumber,
		pf.Price,
	)
}

func (p *PriceFeed) Clone() *PriceFeed {
	return &PriceFeed{
		BlockNumber: p.BlockNumber,
		Feed:        p.Feed,
		RoundId:     p.RoundId,
		// PFVersion:   p.PFVersion,
		PriceBI: p.PriceBI,
		Price:   p.Price,
	}
}

func (PriceFeed) TableName() string {
	return "price_feeds"
}

type SortedPriceFeed []*PriceFeed

func (ts SortedPriceFeed) Len() int {
	return len(ts)
}
func (ts SortedPriceFeed) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

// sort in increasing order by blockNumber,index
func (ts SortedPriceFeed) Less(i, j int) bool {
	return ts[i].BlockNumber < ts[j].BlockNumber
}

type TokenCurrentPrice struct {
	PriceBI  *core.BigInt `gorm:"column:price_bi"`
	Price    float64      `gorm:"column:price"`
	BlockNum int64        `gorm:"column:block_num"`
	Token    string       `gorm:"column:token;primaryKey" json:"token"`
	PriceSrc string       `gorm:"column:price_source;primaryKey" json:"-"`
	Save     bool         `json:"-" gorm:"-"`
}

func (TokenCurrentPrice) TableName() string {
	return "token_current_price"
}

type PFVersion int8

type MergedPFVersion int8

const (
	_            PFVersion = iota
	V1PF         PFVersion = 1 // 1
	V2PF         PFVersion = 2 // 2
	V3PF_MAIN    PFVersion = 4 // 3
	V3PF_REVERSE PFVersion = 8 // 3
)

func (v PFVersion) IsPriceInUSD() bool {
	return v != V1PF
}

func VersionToPFVersion(version core.VersionType, reserve bool) PFVersion {
	if version.MoreThanEq(core.NewVersion(300)) {
		if reserve {
			return V3PF_REVERSE
		}
		return V3PF_MAIN
	} else if version == core.NewVersion(2) {
		return V2PF
	} else if version == core.NewVersion(1) {
		return V1PF
	} else {
		log.Fatal("unknown version type", version)
		panic("")
	}
}

func (v PFVersion) ToVersion() core.VersionType {
	if v == V1PF {
		return core.NewVersion(1)

	} else if v == V2PF {
		return core.NewVersion(2)
	} else if v == V3PF_MAIN || v == V3PF_REVERSE {
		return core.NewVersion(300)
	} else {
		log.Fatal("pf version not supported", v)
		panic("")
	}
}

func (v PFVersion) Decimals() int8 {
	if v&V1PF != 0 {
		return 18 // eth decimals

	} else if (v&V2PF != 0) || (v&V3PF_MAIN != 0) || (v&V3PF_REVERSE != 0) {
		return 8 // USD decimals
	} else {
		log.Fatal("pf version not supported", v)
		panic("")
	}
}
func (v PFVersion) IsReverse() bool {
	return v == V3PF_REVERSE
}

type TokenAndMergedPFVersion struct {
	Token           string
	Feed            string
	MergedPFVersion MergedPFVersion
}

func GetReservefromDetails(details core.Json) bool {
	if details["reserve"] == nil {
		return false
	}
	return details["reserve"].(bool)
}

func allPFVersions() []PFVersion {
	return []PFVersion{V1PF, V2PF, V3PF_MAIN, V3PF_REVERSE}
}
func (v MergedPFVersion) MergedPFVersionToList() (ans []PFVersion) {
	for _, pfVersion := range allPFVersions() {
		if int8(v)&int8(pfVersion) != 0 {
			ans = append(ans, pfVersion)
		}
	}
	return
}

func (v MergedPFVersion) Decimals() int8 {
	for _, version := range v.MergedPFVersionToList() {
		return version.Decimals()
	}
	log.Fatal("Can't get the decimals  for mergedpfversion ", v)
	panic("")
}
