package mposClient

import (
	"strconv"
)

type BlockStats struct {
	OneHourhAmout     float64 `json:"1HourAmount"`
	OneHourhDiff      float64 `json:"1HourDifficulty"`
	OneHourhEstShares uint    `json:"1HourEstimatedShares"`
	OneHourhOrphan    string  `json:"1HourOrphan"`
	OneHourhShares    string  `json:"1HourShares"`
	OneHourhTotal     string  `json:"1HourTotal"`
	OneHourhValid     string  `json:"1HourValid"`

	YesterdayAmout     float64 `json:"24HourAmount"`
	YesterdayDiff      float64 `json:"24HourDifficulty"`
	YesterdayEstShares uint    `json:"24HourEstimatedShares"`
	YesterdayOrphan    string  `json:"24HourOrphan"`
	YesterdayShares    string  `json:"24HourShares"`
	YesterdayTotal     string  `json:"24HourTotal"`
	YesterdayValid     string  `json:"24HourValid"`

	WeekAmout     float64 `json:"7DaysAmount"`
	WeekDiff      float64 `json:"7DaysDifficulty"`
	WeekEstShares uint    `json:"7DaysEstimatedShares"`
	WeekOrphan    string  `json:"7DaysOrphan"`
	WeekShares    string  `json:"7DaysShares"`
	WeekTotal     string  `json:"7DaysTotal"`
	WeekValid     string  `json:"7DaysValid"`

	FourWeeksAmout      float64 `json:"4WeeksAmount"`
	FourWeekshDiff      float64 `json:"4WeeksDifficulty"`
	FourWeekshEstShares uint    `json:"4WeeksEstimatedShares"`
	FourWeekshOrphan    string  `json:"4WeeksOrphan"`
	FourWeekshShares    string  `json:"4WeeksShares"`
	FourWeekshTotal     string  `json:"4WeeksTotal"`
	FourWeekshValid     string  `json:"4WeeksValid"`

	YearAmout     float64 `json:"12MonthAmount"`
	YearDiff      float64 `json:"12MonthDifficulty"`
	YearEstShares uint    `json:"12MonthEstimatedShares"`
	YearOrphan    string  `json:"12MonthOrphan"`
	YearShares    string  `json:"12MonthShares"`
	YearTotal     string  `json:"12MonthTotal"`
	YearValid     string  `json:"12MonthValid"`

	Total          uint    `json:"Total"`
	TotalAmount    float64 `json:"TotalAmount"`
	TotalDiff      float64 `json:"TotalDifficulty"`
	TotalEstSahres uint    `json:"TotalEstimatedShares"`
	TotalOrphan    string  `json:"TotalOrphan"`
	TotalShares    string  `json:"TotalShares"`
	TotalValid     string  `json:"TotalValid"`
}

func (b *BlockStats) OneHourEfficency() float32 {
	var oneHourEfficency float32
	if b.OneHourhShares == "0" {
		return 0
	}
	oneHourShares, err := strconv.ParseFloat(b.OneHourhShares, 32)
	if err != nil {
		return 0
	}
	oneHourEfficency = (float32(oneHourShares) / float32(b.OneHourhEstShares)) * 100
	return oneHourEfficency

}

func (b *BlockStats) YestardayEfficency() float32 {
	var yesterdayEfficency float32
	if b.YesterdayShares == "0" {
		return 0
	}

	yesterdayShares, err := strconv.ParseFloat(b.YesterdayShares, 32)
	if err != nil {
		return 0
	}
	yesterdayEfficency = (float32(yesterdayShares) / float32(b.YesterdayEstShares)) * 100
	return yesterdayEfficency

}

func (b *BlockStats) WeekEfficency() float32 {
	var WeekEfficency float32
	if b.WeekShares == "0" {
		return 0
	}

	WeekShares, err := strconv.ParseFloat(b.WeekShares, 32)
	if err != nil {
		return 0
	}
	WeekEfficency = (float32(WeekShares) / float32(b.WeekEstShares)) * 100
	return WeekEfficency

}

func (b *BlockStats) FourWeekEfficency() float32 {
	var fourWeekEfficency float32
	if b.FourWeekshShares == "0" {
		return 0
	}
	fourWeekShares, err := strconv.ParseFloat(b.FourWeekshShares, 32)
	if err != nil {
		return 0
	}
	fourWeekEfficency = (float32(fourWeekShares) / float32(b.FourWeekshEstShares)) * 100
	return fourWeekEfficency

}

func (b *BlockStats) YearEfficency() float32 {
	var yearEfficency float32
	if b.YearShares == "0" {
		return 0
	}
	yearShares, err := strconv.ParseFloat(b.YearShares, 32)
	if err != nil {
		return 0
	}
	yearEfficency = (float32(yearShares) / float32(b.YearEstShares)) * 100
	return yearEfficency
}

func (b *BlockStats) TotalEfficency() float32 {
	var totalEfficency float32
	if b.TotalShares == "0" {
		return 0
	}

	totalShares, err := strconv.ParseFloat(b.TotalShares, 32)
	if err != nil {
		return 0
	}
	totalEfficency = (float32(totalShares) / float32(b.TotalEstSahres)) * 100
	return totalEfficency

}
