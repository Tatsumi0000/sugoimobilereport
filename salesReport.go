package sugoimobilereport

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

// SalesReport SalesReportををパースする時に使う構造体
type SalesReport struct {
	Provider              string                `csv:"Provider"`
	ProviderCountry       string                `csv:"Provider Country"`
	SKU                   string                `csv:"SKU"`
	Developer             string                `csv:"Developer"`
	Title                 string                `csv:"Title"`
	Version               string                `csv:"Version"`
	ProductTypeIdentifier ProductTypeIdentifier `csv:"Product Type Identifier"`
	Units                 int                   `csv:"Units"`
	DeveloperProceeds     string                `csv:"Developer Proceeds"`
	BeginDate             string                `csv:"Begin Date"`
	EndDate               string                `csv:"End Date"`
	CustomerCurrency      string                `csv:"Customer Currency"`
	CountryCode           string                `csv:"Country Code"`
	CurrencyOfProceeds    string                `csv:"Currency of Proceeds"`
	AppleIdentifier       int                   `csv:"Apple Identifier"`
	CustomerPrice         string                `csv:"Customer Price"`
	PromoCode             string                `csv:"Promo Code"`
	ParentIdentifier      string                `csv:"Parent Identifier"`
	Subscription          string                `csv:"Subscription"`
	Period                string                `csv:"Period"`
	Category              string                `csv:"Category"`
	CMB                   string                `csv:"CMB"`
	Device                string                `csv:"Device"`
	SupportedPlatforms    string                `csv:"Supported Platforms"`
	ProceedsReason        string                `csv:"Proceeds Reason"`
	PreservedPricing      string                `csv:"Preserved Pricing"`
	Client                string                `csv:"Client"`
	OrderType             string                `csv:"Order Type"`
}

// ParseTsvFile TsvFileをパースしてSalesReportのSlicesポインタを返す
func ParseTsvFile(filePath string) ([]*SalesReport, error) {
	salesReports := []*SalesReport{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '\t'
		return r
	})
	// (ファイルパス, 読み込み専用, パーミッション)
	tsvFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer tsvFile.Close()
	if err := gocsv.UnmarshalFile(tsvFile, &salesReports); err != nil {
		return nil, err
	}
	return salesReports, nil
}