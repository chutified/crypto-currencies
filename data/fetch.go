package data

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	models "github.com/chutified/crypto-currencies/models"
	"github.com/pkg/errors"
)

// record holds all values that are stored in the source.
type record struct {
	name              string
	symbol            string
	marketCap         string
	price             string
	circulatingSupply string
	volume            string
	changeHour        string
	changeDay         string
	changeWeek        string
}

func FetchRecords(url string) ([]record, error) {

	// fetch url
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, errors.Wrap(err, "fetching source url")
	}

	// prepare list
	var rs []record
	// handle html
	doc.Find(".cmc-table-row").Each(func(i int, s *goquery.Selection) {
		ch := s.Children()

		// select
		n := ch.First().Next()
		sm := n.Next()
		m := sm.Next()
		p := m.Next()
		cs := p.Next()
		v := cs.Next()
		cr := v.Next()
		cd := cr.Next()
		cw := cd.Next()

		// trim spaces
		name := strings.TrimSpace(n.Text())
		symbol := strings.TrimSpace(sm.Text())
		marketCap := strings.TrimSpace(m.Text())
		price := strings.TrimSpace(p.Text())
		circulatingSupply := strings.TrimSpace(cs.Text())
		volume := strings.TrimSpace(v.Text())
		changeHour := strings.TrimSpace(cr.Text())
		changeDay := strings.TrimSpace(cd.Text())
		changeWeek := strings.TrimSpace(cw.Text())

		// construct
		r := record{
			name:              name,
			symbol:            symbol,
			marketCap:         marketCap,
			price:             price,
			circulatingSupply: circulatingSupply,
			volume:            volume,
			changeHour:        changeHour,
			changeDay:         changeDay,
			changeWeek:        changeWeek,
		}

		// add
		rs = append(rs, r)
	})

	// success
	return rs, nil
}

func ParseRecords(rs []record) (map[string]*models.Currency, error) {

	// prepare holder
	ccs := make(map[string]*models.Currency)

	for _, r := range rs {

		// name
		name := strings.ToUpper(r.name)

		// symbol
		symbol := strings.ToUpper(r.symbol)

		// market cap
		mcStr := r.marketCap[1:]
		mcStr = strings.ReplaceAll(mcStr, ",", "")
		marketCap, err := strconv.ParseFloat(mcStr, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("%s - parse marketCap", name))
		}

		// price
		pStr := r.price[1:]
		pStr = strings.ReplaceAll(pStr, ",", "")
		price, err := strconv.ParseFloat(pStr, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("%s - parse price", name))
		}

		supply := strings.Split(r.circulatingSupply, " ")
		// circulating supply
		csStr := strings.ReplaceAll(supply[0], ",", "")
		circulatingSupply, err := strconv.ParseFloat(csStr, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("%s - parse circulatingSupply", name))
		}
		// mineable
		var mineable bool
		if len(supply) == 3 {
			mineable = true
		}

		// volume
		vStr := r.volume[1:]
		vStr = strings.ReplaceAll(vStr, ",", "")
		volume, err := strconv.ParseFloat(vStr, 64)
		if err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("%s - parse volume", name))
		}

		// changes
		changeHour := r.changeHour
		changeDay := r.changeDay
		changeWeek := r.changeWeek

		// construct
		c := &models.Currency{
			Name:              name,
			Symbol:            symbol,
			MarketCapUSD:      marketCap,
			Price:             price,
			CirculatingSupply: circulatingSupply,
			Mineable:          mineable,
			Volume:            volume,
			ChangeHour:        changeHour,
			ChangeDay:         changeDay,
			ChangeWeek:        changeWeek,
		}

		// add to map
		ccs[name] = c
	}

	return ccs, nil
}
