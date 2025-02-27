package network

import (
	"errors"
	"fmt"
	"os"

	"github.com/fxsjy/gonn/gonn"

	"lr3/internal/app/network/data"
	"lr3/internal/constants"
)

type Network struct {
	Data *data.NNData
}

func New() *Network {
	return &Network{
		Data: data.New(),
	}
}

func (n *Network) create() {
	nn := gonn.DefaultNetwork(4, 4, 1, false)
	nn.Train(n.Data.Inputs, n.Data.Outputs, 100000)
	gonn.DumpNN(constants.NnDumpName.String(), nn)
}

func (n *Network) GetNN(forceCreate bool, onComplete func()) *gonn.NeuralNetwork {
	_, err := os.Stat(constants.NnDumpName.String())
	if errors.Is(err, os.ErrNotExist) || forceCreate {
		n.create()
	}
	dump := gonn.LoadNN(constants.NnDumpName.String())
	onComplete()
	return dump
}

func (n *Network) GetResult(output []float64) string {
	return fmt.Sprintf("%v", output[0])
}
