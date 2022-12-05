package game

import (
	"bet/model"
	"bet/utils"
	"fmt"
	"time"
)

type GameIssueRespone struct {
	Issue          int64     `json:"issue"`
	OpenTime       time.Time `json:"open_time"`
	CloseTime      time.Time `json:"close_time"`
	StartTime      time.Time `json:"start_time"`
	OpenTimeStr    string    `json:"open_time_str"`
	CloseTimeStr   string    `json:"close_time_str"`
	StartTimeStr   string    `json:"start_time_str"`
	CloseCountdown int64     `json:"bet_countdown"`
	TimeNow        time.Time `json:"time_now"`
	Date           string    `json:"date"`
}

func GetIssue(gameCode int32) *utils.Result[GameIssueRespone] {

	game := &model.SysGame{}
	game.GetByCodeCache(gameCode)

	rIssue := &utils.Result[GameIssueRespone]{
		Code:      500,
		IsSuccess: false,
	}
	if game.Status != 1 {
		rIssue.Message = "Status Off"
		return rIssue
	}
	//2006-01-02 15:04:05
	issueList := make([]GameIssueRespone, 0)

	format := time.Now().Format("2006-01-02")
	timeStart, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", format, game.StartTime), time.Local)
	timeEnd, _ := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s %s", format, game.EndTime), time.Local)
	var issueIndex int64 = 1
	for timeStart.Before(timeEnd) {
		duration, _ := time.ParseDuration(fmt.Sprintf("%ds", game.IntervalSeconds))
		closeDur, _ := time.ParseDuration(fmt.Sprintf("-%ds", 10))
		m := GameIssueRespone{}
		m.StartTime = timeStart
		timeStart = timeStart.Add(duration)
		m.Issue = issueIndex
		issueIndex++
		m.CloseTime = timeStart.Add(closeDur)
		m.Date = timeStart.Format("2006-01-02")
		m.OpenTime = timeStart
		issueList = append(issueList, m)

	}
	var currentIssue GameIssueRespone
	timeNow := time.Now()
	for _, v := range issueList {

		if timeNow.After(v.StartTime) && timeNow.Before(v.OpenTime) {

			v.CloseCountdown = v.CloseTime.Unix() - timeNow.Unix()
			v.TimeNow = timeNow
			currentIssue = v
			continue
		}
	}
	rIssue.Message = ""
	rIssue.Code = 200
	currentIssue.CloseTimeStr = currentIssue.CloseTime.Format("2006-01-02 15:04:05")
	currentIssue.OpenTimeStr = currentIssue.OpenTime.Format("2006-01-02 15:04:05")
	currentIssue.StartTimeStr = currentIssue.StartTime.Format("2006-01-02 15:04:05")

	rIssue.Data = currentIssue

	fmt.Print(currentIssue)
	return rIssue
}
