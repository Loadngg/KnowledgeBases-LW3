package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatMatrix(matrix [][]float64, precision int) string {
	if len(matrix) == 0 {
		return ""
	}

	strMatrix := make([][]string, len(matrix))
	for i, row := range matrix {
		strMatrix[i] = make([]string, len(row))
		for j, val := range row {
			strMatrix[i][j] = fmt.Sprintf("%.*f", precision, val)
		}
	}

	widths := make([]int, len(matrix[0]))
	for j := range matrix[0] {
		max := 0
		for i := range matrix {
			if len(strMatrix[i][j]) > max {
				max = len(strMatrix[i][j])
			}
		}
		widths[j] = max
	}

	// Собираем результат
	var sb strings.Builder
	for i, row := range strMatrix {
		for j, val := range row {
			fmt.Fprintf(&sb, "%*s", widths[j], val)
			if j < len(row)-1 {
				sb.WriteString("  ")
			}
		}
		if i < len(matrix)-1 {
			sb.WriteByte('\n')
		}
	}

	return sb.String()
}

func ParseFloat64Array(text string) ([]float64, error) {
	splitText := strings.Split(text, " ")
	result := make([]float64, len(splitText))
	for i, str := range splitText {
		val, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, err
		}
		result[i] = val
	}
	return result, nil
}
