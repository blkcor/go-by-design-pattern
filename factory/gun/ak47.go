package gun

type AK47 struct {
	Gun
}

func newAk47() IGun {
	return &AK47{
		Gun{
			name:  "AK47",
			power: 100,
		},
	}
}
