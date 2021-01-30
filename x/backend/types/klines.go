package types

import (
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/pkg/errors"
)

const (
	KlineTypeM1      = "kline_m1"
	KlineTypeM3      = "kline_m3"
	KlineTypeM5      = "kline_m5"
	KlineTypeM15     = "kline_m15"
	KlineTypeM30     = "kline_m30"
	KlineTypeM60     = "kline_m60"
	KlineTypeM120    = "kline_m120"
	KlineTypeM240    = "kline_m240"
	KlineTypeM360    = "kline_m360"
	KlineTypeM720    = "kline_m720"
	KlineTypeM1440   = "kline_m1440"
	KlineTypeM4320   = "kline_m4320"
	KlineTypeM10080  = "kline_m10080"
	KlineTypeM44640  = "kline_m44640"
	KlineTypeM525600 = "kline_m525600"
)

// nolint
type IKline interface {
	GetFreqInSecond() int
	GetAnchorTimeTS(ts int64) int64
	GetTableName() string
	GetProduct() string
	GetTimestamp() int64
	GetOpen() float64
	GetClose() float64
	GetHigh() float64
	GetLow() float64
	GetVolume() float64
	PrettyTimeString() string
	GetBriefInfo() []string
}

var (
	kline2channel = map[string]string{
		KlineTypeM1:      "dex_spot/candle60s",
		KlineTypeM3:      "dex_spot/candle180s",
		KlineTypeM5:      "dex_spot/candle300s",
		KlineTypeM15:     "dex_spot/candle900s",
		KlineTypeM30:     "dex_spot/candle1800s",
		KlineTypeM60:     "dex_spot/candle3600s",
		KlineTypeM120:    "dex_spot/candle7200s",
		KlineTypeM240:    "dex_spot/candle14400s",
		KlineTypeM360:    "dex_spot/candle21600s",
		KlineTypeM720:    "dex_spot/candle43200s",
		KlineTypeM1440:   "dex_spot/candle86400s",
		KlineTypeM4320:   "dex_spot/candle259200s",
		KlineTypeM10080:  "dex_spot/candle604800s",
		KlineTypeM44640:  "dex_spot/candle2678400s",
		KlineTypeM525600: "dex_spot/candle31536000s",
	}

	klineType2Freq = map[string]int{
		KlineTypeM1:      60,
		KlineTypeM3:      180,
		KlineTypeM5:      300,
		KlineTypeM15:     900,
		KlineTypeM30:     1800,
		KlineTypeM60:     3600,
		KlineTypeM120:    7200,
		KlineTypeM240:    14400,
		KlineTypeM360:    21600,
		KlineTypeM720:    43200,
		KlineTypeM1440:   86400,
		KlineTypeM4320:   259200,
		KlineTypeM10080:  604800,
		KlineTypeM44640:  2678400,
		KlineTypeM525600: 31536000,
	}
)

func GetChannelByKlineType(klineType string) string {
	return kline2channel[klineType]
}

func GetFreqByKlineType(klineType string) int {
	return klineType2Freq[klineType]
}

// nolint
type IKlines []IKline

// nolint
func (klines IKlines) Len() int {
	return len(klines)
}

// nolint
func (klines IKlines) Swap(i, j int) {
	klines[i], klines[j] = klines[j], klines[i]
}

// nolint
func (klines IKlines) Less(i, j int) bool {
	return klines[i].GetTimestamp() < klines[j].GetTimestamp()
}

// nolint
type IKlinesDsc []IKline

// nolint
func (klines IKlinesDsc) Len() int {
	return len(klines)
}

// nolint
func (c IKlinesDsc) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// nolint
func (klines IKlinesDsc) Less(i, j int) bool {
	return klines[i].GetTimestamp() > klines[j].GetTimestamp()
}

// BaseKline define the basic data of Kine
type BaseKline struct {
	Product   string  `gorm:"PRIMARY_KEY;type:varchar(20)" json:"product"`
	Timestamp int64   `gorm:"PRIMARY_KEY;type:bigint;" json:"timestamp"`
	Open      float64 `gorm:"type:DOUBLE" json:"open"`
	Close     float64 `gorm:"type:DOUBLE" json:"close"`
	High      float64 `gorm:"type:DOUBLE" json:"high"`
	Low       float64 `gorm:"type:DOUBLE" json:"low"`
	Volume    float64 `gorm:"type:DOUBLE" json:"volume"`
	impl      IKline
}

func (b *BaseKline) GetChannelInfo() (channel, filter string, err error) {
	if b.impl == nil {
		return "", "", errors.New("failed to find channel because of no specific kline type found")
	}

	channel, ok := kline2channel[b.impl.GetTableName()]
	if !ok {
		return "", "", errors.Errorf("failed to find channel for %s", b.GetTableName())
	}
	return channel, b.Product, nil
}

func (b *BaseKline) GetFullChannel() string {
	channel, filter, e := b.GetChannelInfo()
	if e == nil {
		return channel + ":" + filter
	} else {
		return "InvalidChannelName"
	}
}

// GetFreqInSecond return interval time
func (b *BaseKline) GetFreqInSecond() int {
	if b.impl != nil {
		return b.impl.GetFreqInSecond()
	} else {
		return -1
	}
}

// GetTableName rerurn database table name
func (b *BaseKline) GetTableName() string {
	if b.impl != nil {
		return b.impl.GetTableName()
	} else {
		return "base_kline"
	}
}

// GetAnchorTimeTS return time interval
func (b *BaseKline) GetAnchorTimeTS(ts int64) int64 {
	m := (ts / int64(b.GetFreqInSecond())) * int64(b.GetFreqInSecond())
	return m
}

// GetProduct return product
func (b *BaseKline) GetProduct() string {
	return b.Product
}

// GetTimestamp return timestamp
func (b *BaseKline) GetTimestamp() int64 {
	return b.Timestamp
}

// GetOpen return open price
func (b *BaseKline) GetOpen() float64 {
	return b.Open
}

// GetClose return close price
func (b *BaseKline) GetClose() float64 {
	return b.Close
}

// GetHigh return high price
func (b *BaseKline) GetHigh() float64 {
	return b.High
}

// GetLow return low price
func (b *BaseKline) GetLow() float64 {
	return b.Low
}

// GetVolume return volume of trade quantity
func (b *BaseKline) GetVolume() float64 {
	return b.Volume
}

// GetBriefInfo return array of kline data
func (b *BaseKline) GetBriefInfo() []string {
	m := []string{
		time.Unix(b.GetTimestamp(), 0).UTC().Format("2006-01-02T15:04:05.000Z"),
		fmt.Sprintf("%.4f", b.GetOpen()),
		fmt.Sprintf("%.4f", b.GetHigh()),
		fmt.Sprintf("%.4f", b.GetLow()),
		fmt.Sprintf("%.4f", b.GetClose()),
		fmt.Sprintf("%.8f", b.GetVolume()),
	}
	return m
}

func (b *BaseKline) FormatResult() interface{} {
	result := map[string]interface{}{}
	result["instrument_id"] = b.Product
	result["candle"] = b.GetBriefInfo()
	return result
}

// TimeString  format time
func TimeString(ts int64) string {
	return time.Unix(ts, 0).Local().Format("20060102_150405")
}

// PrettyTimeString  convert kline data to string
func (b *BaseKline) PrettyTimeString() string {
	return fmt.Sprintf("Product: %s, Freq: %d, Time: %s, OCHLV(%.4f, %.4f, %.4f, %.4f, %.4f)",
		b.Product, b.GetFreqInSecond(), TimeString(b.Timestamp), b.Open, b.Close, b.High, b.Low, b.Volume)
}

// KlineM1 define kline data in 1 minute
type KlineM1 struct {
	*BaseKline
}

// NewKlineM1 create a instance of KlineM1
func NewKlineM1(b *BaseKline) *KlineM1 {
	k := KlineM1{b}
	k.impl = &k
	return &k
}

// GetFreqInSecond return 60
func (k *KlineM1) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// GetTableName return kline_m1
func (k *KlineM1) GetTableName() string {
	return KlineTypeM1
}

// KlineM3 define kline data in 3 minutes
type KlineM3 struct {
	*BaseKline
}

// NewKlineM3 create a instance of KlineM3
func NewKlineM3(b *BaseKline) *KlineM3 {
	k := KlineM3{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m3
func (k *KlineM3) GetTableName() string {
	return KlineTypeM3
}

// GetFreqInSecond return 180
func (k *KlineM3) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM5 define kline data in 5 minutes
type KlineM5 struct {
	*BaseKline
}

// NewKlineM5 create a instance of KlineM5
func NewKlineM5(b *BaseKline) *KlineM5 {
	k := KlineM5{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m5
func (k *KlineM5) GetTableName() string {
	return KlineTypeM5
}

// GetFreqInSecond return 300
func (k *KlineM5) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM15 define kline data in 15 minutes
type KlineM15 struct {
	*BaseKline
}

// NewKlineM15 create a instance of KlineM15
func NewKlineM15(b *BaseKline) *KlineM15 {
	k := KlineM15{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m15
func (k *KlineM15) GetTableName() string {
	return KlineTypeM15
}

// GetFreqInSecond return 900
func (k *KlineM15) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM30 define kline data in 30 minutes
type KlineM30 struct {
	*BaseKline
}

// NewKlineM30 create a instance of KlineM30
func NewKlineM30(b *BaseKline) *KlineM30 {
	k := KlineM30{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m30
func (k *KlineM30) GetTableName() string {
	return KlineTypeM30
}

// GetFreqInSecond return 1800
func (k *KlineM30) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM60 define kline data in 1 hour
type KlineM60 struct {
	*BaseKline
}

// NewKlineM60 create a instance of KlineM60
func NewKlineM60(b *BaseKline) *KlineM60 {
	k := KlineM60{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m60
func (k *KlineM60) GetTableName() string {
	return KlineTypeM60
}

// GetFreqInSecond return 3600
func (k *KlineM60) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM120 define kline data in 2 hours
type KlineM120 struct {
	*BaseKline
}

// NewKlineM120 create a instance of KlineM120
func NewKlineM120(b *BaseKline) *KlineM120 {
	k := KlineM120{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m120
func (k *KlineM120) GetTableName() string {
	return KlineTypeM120
}

// GetFreqInSecond return 7200
func (k *KlineM120) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM240 define kline data in 4 hours
type KlineM240 struct {
	*BaseKline
}

// NewKlineM240 create a instance of KlineM240
func NewKlineM240(b *BaseKline) *KlineM240 {
	k := KlineM240{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m240
func (k *KlineM240) GetTableName() string {
	return KlineTypeM240
}

// GetFreqInSecond return 14400
func (k *KlineM240) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM360 define kline data in 6 hours
type KlineM360 struct {
	*BaseKline
}

// NewKlineM360 create a instance of KlineM360
func NewKlineM360(b *BaseKline) *KlineM360 {
	k := KlineM360{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m360
func (k *KlineM360) GetTableName() string {
	return KlineTypeM360
}

// GetFreqInSecond return 21600
func (k *KlineM360) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM720 define kline data in 12 hours
type KlineM720 struct {
	*BaseKline
}

// NewKlineM720 create a instance of KlineM720
func NewKlineM720(b *BaseKline) *KlineM720 {
	k := KlineM720{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m720
func (k *KlineM720) GetTableName() string {
	return KlineTypeM720
}

// GetFreqInSecond return 43200
func (k *KlineM720) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM1440 define kline data in 1 day
type KlineM1440 struct {
	*BaseKline
}

// NewKlineM1440 create a instance of NewKlineM1440
func NewKlineM1440(b *BaseKline) *KlineM1440 {
	k := KlineM1440{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m1440
func (k *KlineM1440) GetTableName() string {
	return KlineTypeM1440
}

// GetFreqInSecond return 86400
func (k *KlineM1440) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM4320 define kline data in 1 day
type KlineM4320 struct {
	*BaseKline
}

// NewKlineM4320 create a instance of KlineM4320
func NewKlineM4320(b *BaseKline) *KlineM4320 {
	k := KlineM4320{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m4320
func (k *KlineM4320) GetTableName() string {
	return KlineTypeM4320
}

// GetFreqInSecond return 259200
func (k *KlineM4320) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM10080 define kline data in 1 week
type KlineM10080 struct {
	*BaseKline
}

// NewKlineM10080 create a instance of NewKlineM10080
func NewKlineM10080(b *BaseKline) *KlineM10080 {
	k := KlineM10080{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m10080
func (k *KlineM10080) GetTableName() string {
	return KlineTypeM10080
}

// GetFreqInSecond return 604800
func (k *KlineM10080) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM44640 define kline data in 1 day
type KlineM44640 struct {
	*BaseKline
}

// NewKlineM44640 create a instance of KlineM44640
func NewKlineM44640(b *BaseKline) *KlineM44640 {
	k := KlineM44640{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m44640
func (k *KlineM44640) GetTableName() string {
	return KlineTypeM44640
}

// GetFreqInSecond return 2678400
func (k *KlineM44640) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// KlineM525600 define kline data in 1 day
type KlineM525600 struct {
	*BaseKline
}

// NewKlineM525600 create a instance of KlineM525600
func NewKlineM525600(b *BaseKline) *KlineM525600 {
	k := KlineM525600{b}
	k.impl = &k
	return &k
}

// GetTableName return kline_m525600
func (k *KlineM525600) GetTableName() string {
	return KlineTypeM525600
}

// GetFreqInSecond return 31536000
func (k *KlineM525600) GetFreqInSecond() int {
	return klineType2Freq[k.GetTableName()]
}

// MustNewKlineFactory will panic when err occurred during  NewKlineFactory
func MustNewKlineFactory(name string, baseK *BaseKline) (r interface{}) {
	r, err := NewKlineFactory(name, baseK)

	if err != nil {
		panic(err)
	}
	return r
}

// NewKlineFactory generate kline type by factory pattern
func NewKlineFactory(name string, baseK *BaseKline) (r interface{}, err error) {
	b := baseK
	if b == nil {
		b = &BaseKline{}
	}

	switch name {
	case KlineTypeM1:
		return NewKlineM1(b), nil
	case KlineTypeM3:
		return NewKlineM3(b), nil
	case KlineTypeM5:
		return NewKlineM5(b), nil
	case KlineTypeM15:
		return NewKlineM15(b), nil
	case KlineTypeM30:
		return NewKlineM30(b), nil
	case KlineTypeM60:
		return NewKlineM60(b), nil
	case KlineTypeM120:
		return NewKlineM120(b), nil
	case KlineTypeM240:
		return NewKlineM240(b), nil
	case KlineTypeM360:
		return NewKlineM360(b), nil
	case KlineTypeM720:
		return NewKlineM720(b), nil
	case KlineTypeM1440:
		return NewKlineM1440(b), nil
	case KlineTypeM4320:
		return NewKlineM4320(b), nil
	case KlineTypeM10080:
		return NewKlineM10080(b), nil
	case KlineTypeM44640:
		return NewKlineM44640(b), nil
	case KlineTypeM525600:
		return NewKlineM525600(b), nil
	}

	return nil, errors.New("No kline constructor function found.")
}

// GetAllKlineMap return map about kline table names
func GetAllKlineMap() map[int]string {
	return map[int]string{
		60:       KlineTypeM1,
		180:      KlineTypeM3,
		300:      KlineTypeM5,
		900:      KlineTypeM15,
		1800:     KlineTypeM30,
		3600:     KlineTypeM60,
		7200:     KlineTypeM120,
		14400:    KlineTypeM240,
		21600:    KlineTypeM360,
		43200:    KlineTypeM720,
		86400:    KlineTypeM1440,
		259200:   KlineTypeM4320,
		604800:   KlineTypeM10080,
		2678400:  KlineTypeM44640,
		31536000: KlineTypeM525600,
	}
}

// GetKlineTableNameByFreq return table name
func GetKlineTableNameByFreq(freq int) string {
	m := GetAllKlineMap()
	name := m[freq]
	return name

}

// NewKlinesFactory generate kline type by type of kline
func NewKlinesFactory(name string) (r interface{}, err error) {
	switch name {
	case KlineTypeM1:
		return &[]KlineM1{}, nil
	case KlineTypeM3:
		return &[]KlineM3{}, nil
	case KlineTypeM5:
		return &[]KlineM5{}, nil
	case KlineTypeM15:
		return &[]KlineM15{}, nil
	case KlineTypeM30:
		return &[]KlineM30{}, nil
	case KlineTypeM60:
		return &[]KlineM60{}, nil
	case KlineTypeM120:
		return &[]KlineM120{}, nil
	case KlineTypeM240:
		return &[]KlineM240{}, nil
	case KlineTypeM360:
		return &[]KlineM360{}, nil
	case KlineTypeM720:
		return &[]KlineM720{}, nil
	case KlineTypeM1440:
		return &[]KlineM1440{}, nil
	case KlineTypeM4320:
		return &[]KlineM4320{}, nil
	case KlineTypeM10080:
		return &[]KlineM10080{}, nil
	case KlineTypeM44640:
		return &[]KlineM44640{}, nil
	case KlineTypeM525600:
		return &[]KlineM525600{}, nil
	}

	return nil, ErrNoKlinesFunctionFound()
}

// ToIKlinesArray Convert kline data to array for restful interface
func ToIKlinesArray(klines interface{}, endTS int64, doPadding bool) []IKline {

	originKlines := []IKline{}

	v := reflect.ValueOf(klines)
	elements := v.Elem()
	if elements.Kind() == reflect.Slice {
		for i := 0; i < elements.Len(); i++ {
			r := elements.Index(i).Interface()
			switch r.(type) {
			case KlineM1:
				r2 := r.(KlineM1)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM3:
				r2 := r.(KlineM3)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM5:
				r2 := r.(KlineM5)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM15:
				r2 := r.(KlineM15)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM30:
				r2 := r.(KlineM30)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM60:
				r2 := r.(KlineM60)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM120:
				r2 := r.(KlineM120)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM240:
				r2 := r.(KlineM240)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM360:
				r2 := r.(KlineM360)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM720:
				r2 := r.(KlineM720)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM1440:
				r2 := r.(KlineM1440)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM4320:
				r2 := r.(KlineM4320)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM10080:
				r2 := r.(KlineM10080)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM44640:
				r2 := r.(KlineM44640)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			case KlineM525600:
				r2 := r.(KlineM525600)
				r2.impl = &r2
				originKlines = append(originKlines, &r2)
			}
		}
	}

	if elements.Kind() != reflect.Slice || len(originKlines) == 0 {
		return originKlines
	}

	// 0. Pad Latest Kline
	lastKline := originKlines[0]
	anchorTS := lastKline.GetAnchorTimeTS(endTS)
	if anchorTS > originKlines[0].GetTimestamp() && doPadding {
		baseKline := BaseKline{
			Product:   lastKline.GetProduct(),
			Timestamp: anchorTS,
			Open:      lastKline.GetClose(),
			Close:     lastKline.GetClose(),
			High:      lastKline.GetClose(),
			Low:       lastKline.GetClose(),
			Volume:    0,
		}
		newKline := MustNewKlineFactory(lastKline.GetTableName(), &baseKline)
		newKlines := []IKline{newKline.(IKline)}
		originKlines = append(newKlines, originKlines...)

	}

	// 1. Padding lost klines
	paddings := IKlines{}
	size := len(originKlines)
	for i := size - 1; i > 0 && doPadding; i-- {
		crrIKline := originKlines[i]
		nextIKline := originKlines[i-1]
		expectNextTime := crrIKline.GetTimestamp() + int64(crrIKline.GetFreqInSecond())
		for expectNextTime < nextIKline.GetTimestamp() {
			baseKline := BaseKline{
				Product:   crrIKline.GetProduct(),
				Timestamp: expectNextTime,
				Open:      crrIKline.GetClose(),
				Close:     crrIKline.GetClose(),
				High:      crrIKline.GetClose(),
				Low:       crrIKline.GetClose(),
				Volume:    0,
			}

			newKline := MustNewKlineFactory(crrIKline.GetTableName(), &baseKline)
			paddings = append(paddings, newKline.(IKline))
			expectNextTime += int64(crrIKline.GetFreqInSecond())
		}
	}

	// 2. Merge origin klines & padding klines
	paddings = append(paddings, originKlines...)
	sort.Sort(paddings)

	return paddings
}

// nolint
func ToRestfulData(klines *[]IKline, limit int) [][]string {

	// Return restful datas
	m := [][]string{}
	to := len(*klines)
	from := to - limit
	if from < 0 {
		from = 0
	}

	if to <= 0 {
		return m
	}

	for _, k := range (*klines)[from:to] {
		m = append(m, k.GetBriefInfo())
	}
	return m
}
