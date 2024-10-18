package utils

func Map[T, I any](ts []T, f func(T) I) []I {
	mapped := make([]I, len(ts))

	for i := range ts {
		mapped[i] = f(ts[i])
	}

	return mapped
}
