package Score

type ScoreData struct {
	YakuList     []string
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

type Yaku struct {
	List   []string
	Hansuu int
}

func NewYaku(list []string) *Yaku {
	yaku := new(Yaku)
	yaku.List = list
	return yaku
}
