package service

// CalculateDamage ダメージ計算式
func CalculateDamage(
	level int64,
	power int64,
	attack int64,
	block int64,
) (int64, int64) {
	max := (float64(level) * float64(2) / float64(5)) + float64(2)
	max = float64(int64(max)) *
		float64(power) *
		float64(attack) /
		float64(block)
	max = (float64(int64(max)) / 50) + 2

	// min, max の順で返す
	return int64(max * 0.85), int64(max)
}
