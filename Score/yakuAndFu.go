package Score

type ScoreData struct {
	Hansuu               int
	Fu                   int
}

func NewScoreData(hansuu int, fu int) *ScoreData {
	scoreData := new(ScoreData)
	scoreData.Hansuu = hansuu
	scoreData.Fu = fu

	return scoreData
}

type Yaku struct {
	List []string
	Hansuu int
}

func NewYaku(list []string) *Yaku{
	yaku := new(Yaku)
	yaku.List=list
	return yaku
}

