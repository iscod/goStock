package price

import (
	"fmt"
	"github.com/iscod/goStock/base"
	"github.com/iscod/goStock/model"
	"gorm.io/gorm"
	"time"
)

func Run(symbol string, db *gorm.DB) {
	fmt.Printf("名称: %s\n", symbol)

	//获取价格信息
	quote, err := base.GetQuote(symbol)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(" %s, %s, 当前价格: %f, 开盘价: %f, 均价: %f, 当天最低: %f, 当天最高: %f, %d\n", quote.Name, quote.Symbol, quote.Current, quote.Open, quote.AvgPrice, quote.Low, quote.High, quote.Time)

	quote.CreatedAt = quote.Time / 1000
	quote.ExecAt = time.Unix(quote.CreatedAt, 0).Format("2006-01-02")
	err = db.Where(model.Quote{Symbol: quote.Symbol, ExecAt: quote.ExecAt}).FirstOrCreate(quote).Error
	if err == gorm.ErrRecordNotFound {
		err = db.Save(quote).Error
	}
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
}