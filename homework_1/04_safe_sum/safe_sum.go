package safe_sum

import "math"

// SafeSum последовательно суммирует транзакции, пропуская те, добавление
// которых привело бы к переполнению int64. skipped — индексы пропущенных
// элементов исходного среза, в порядке появления.
func SafeSum(transactions []int64) (total int64, skipped []int) {
	total = 0

	for i, transaction := range transactions {
		if transaction > 0 {
			if total > math.MaxInt64-transaction {
				skipped = append(skipped, i)
				continue
			}
		} else if transaction < 0 {
			if total < math.MinInt64-transaction {
				skipped = append(skipped, i)
				continue
			}
		}
		total += transaction
	}

	return total, skipped
}
