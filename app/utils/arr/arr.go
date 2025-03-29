package arr

func Reduce[T, S any](arr []T, fn func(acc S, item T) S, initial S) S {
	for _, item := range arr {
		initial = fn(initial, item)
	}
	return initial
}
