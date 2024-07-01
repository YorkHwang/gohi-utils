package slices

func Transform[E1 any, T2 ~[]E2, E2 any](source []E1, apply func(target *T2)) T2 {
	if source == nil {
		return nil
	}
	if len(source) == 0 {
		return T2{}
	}
	target := make(T2, 0, len(source))
	apply(&target)
	return target
}

// TransformTo
// deprecated
func TransformTo[T1 any, T2 any](target *[]T2, source []T1, apply func(target *[]T2)) {
	if source == nil {
		return
	}
	if len(source) == 0 {
		return
	}
	apply(target)
}

func TransformToMap[E any, M ~map[K]V, K comparable, V any](source []E, apply func(target M)) M {
	if source == nil {
		return nil
	}
	if len(source) == 0 {
		return M{}
	}
	target := make(M, len(source))
	apply(target)
	return target
}

func Associate[E any, K comparable, V any](source []E, transform func(it E) (K, V)) map[K]V {
	return TransformToMap(source, func(target map[K]V) {
		for _, it := range source {
			k, v := transform(it)
			target[k] = v
		}
	})
}

func AssociateBy[E any, K comparable](source []E, transform func(it E) K) map[K]E {
	return TransformToMap(source, func(target map[K]E) {
		for _, it := range source {
			target[transform(it)] = it
		}
	})
}

func AssociateWith[E comparable, V any](source []E, transform func(it E) V) map[E]V {
	return TransformToMap(source, func(target map[E]V) {
		for _, it := range source {
			target[it] = transform(it)
		}
	})
}
