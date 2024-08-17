package main

import "iter"

func Map[A any, B any](s []A, f func(A) B) iter.Seq2[int, B] {
	return func(y func(idx int, v B) bool) {
		for i := 0; i < len(s); i++ {
			if !y(i, f(s[i])) {
				return
			}
		}
	}

}
