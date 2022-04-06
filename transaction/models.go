package transaction

import (
	"fmt"
	"time"

	"jobapp.com/m/common"
)

type TransactionModel struct {
	Id         int                  `json:"id"`
	MerchantId int                  `json:"merchant_id"`
	OutletId   int                  `json:"outlet_id"`
	BillTotal  float64              `json:"bill_total"`
	Modified   common.ModifiedModel `json:"modified"`
}

type OmzetPerDayModel struct {
	MerchantName string `json:"merchant_name"`
	TotalOmzet   string `json:"total_omzet"`
	Date         string `json:"date"`
}

func OmzetPerDay(userId string, yearMonth string) ([]OmzetPerDayModel, error) {

	var models []OmzetPerDayModel

	rows, err := common.Db.Query(`SELECT MIN(M.merchant_name) AS merchant_name, SUM(T.bill_total) AS total_omzet, DATE_FORMAT(T.updated_at, '%Y-%m-%d') AS date
	FROM transactions AS T 
	JOIN merchants AS M ON T.merchant_id=M.id
	JOIN users AS U ON M.user_id=U.id
	WHERE U.id = ? AND DATE_FORMAT(T.updated_at, '%Y-%m') = ?
	GROUP BY date `, userId, yearMonth)

	if err != nil {
		return models, err
	}

	for rows.Next() {
		var model OmzetPerDayModel
		err := rows.Scan(&model.MerchantName, &model.TotalOmzet, &model.Date)
		if err != nil {
			return models, err
		}
		models = append(models, model)
	}

	t, _ := time.Parse("2006-01", yearMonth)

	year, month, _ := t.Date()
	_, _, lastDay := time.Date(year, month+1, 0, 0, 0, 0, 0, time.Local).Date()

	var tempModels []OmzetPerDayModel

	merchantName := models[0].MerchantName

	fmt.Println("lastDay : ", lastDay)

	for index := 1; index <= lastDay; index++ {
		date := fmt.Sprint(year, "-", int(month), "-", fmt.Sprintf("%02d", index))
		fmt.Println("date : ", date)
		isSameDay := false
		for _, model := range models {
			if date == model.Date {
				tempModels = append(tempModels, model)
				isSameDay = true
			}
		}
		if !isSameDay {
			tempModels = append(tempModels, OmzetPerDayModel{
				MerchantName: merchantName,
				TotalOmzet:   "0",
				Date:         date,
			})
		}
	}

	return tempModels, err
}
