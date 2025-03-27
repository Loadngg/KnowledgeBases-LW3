package data

var inputs = [][]float64{
	{0, 0, 0, 0},
	{0, 1, 0, 0},
	{1, 0, 1, 0},
	{1, 1, 1, 0},
	{0, 0, 0, 0},
	{0, 1, 0, 0},
	{1, 0, 1, 0},
	{1, 1, 1, 1},
	{0, 0, 0, 1},
	{0, 1, 0, 1},
	{1, 0, 1, 1},
	{1, 1, 1, 1},
	{0, 0, 0, 1},
	{0, 1, 0, 1},
	{1, 0, 1, 1},
	{1, 1, 1, 1},
}

var outputs = [][]float64{
	{0},
	{0},
	{1},
	{1},
	{0},
	{0},
	{1},
	{1},
	{0},
	{0},
	{1},
	{1},
	{0},
	{0},
	{1},
	{1},
}

type NNData struct {
	Inputs  [][]float64
	Outputs [][]float64
}

func New() *NNData {
	return &NNData{
		Inputs:  inputs,
		Outputs: outputs,
	}
}

func (d *NNData) GetMergedData() [][]float64 {
	var result [][]float64
	for index, item := range d.Inputs {
		result = append(result, append(item, d.Outputs[index][0]))
	}
	return result
}
