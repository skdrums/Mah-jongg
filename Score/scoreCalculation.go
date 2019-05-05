package Score

type ScoreData struct {
	Hansuu   int
	Fu       int
	IsParent bool
	IsTsumo  bool
}

func NewScoreData(hansuu int, fu int) *ScoreData {
	scoreData := new(ScoreData)
	scoreData.Hansuu = hansuu
	scoreData.Fu = fu

	return scoreData
}
