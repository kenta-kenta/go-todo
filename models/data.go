package models

import "time"

var (
	Todo1 = Todo{
		TodoID:    1,
		Content:   "朝ご飯を食べる",
		CreatedAt: time.Now(),
	}

	Todo2 = Todo{
		TodoID:    2,
		Content:   "散歩する",
		CreatedAt: time.Now(),
	}
)
