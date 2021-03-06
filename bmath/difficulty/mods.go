package difficulty

type Modifier int64

const (
	None   = Modifier(iota)
	NoFail = Modifier(1 << (iota - uint(1)))
	Easy
	TouchDevice
	Hidden
	HardRock
	SuddenDeath
	DoubleTime
	Relax
	HalfTime
	Nightcore  // Only set along with DoubleTime. i.e: NC only gives 576
	Flashlight
	Autoplay
	SpunOut
	Relax2   // Autopilot
	Perfect  // Only set along with SuddenDeath. i.e: PF only gives 16416
	Key4
	Key5
	Key6
	Key7
	Key8
	FadeIn
	Random
	Cinema
	Target
	Key9
	KeyCoop
	Key1
	Key3
	Key2
	ScoreV2
	LastMod
)

func (modifier Modifier) GetScoreMultiplier() float64 {
	multiplier := 1.0

	if modifier&NoFail > 0 {
		multiplier *= 0.5
	}

	if modifier&Easy > 0 {
		multiplier *= 0.5
	}

	if modifier&HalfTime > 0 {
		multiplier *= 0.3
	}

	if modifier&Hidden > 0 {
		multiplier *= 1.06
	}

	if modifier&HardRock > 0 {
		multiplier *= 1.06
	}

	if modifier&DoubleTime > 0 {
		multiplier *= 1.12
	}

	if modifier&Flashlight > 0 {
		multiplier *= 1.12
	}

	if (modifier&Relax | modifier&Relax2) > 0 {
		multiplier = 0
	}

	if modifier&SpunOut > 0 {
		multiplier *= 0.9
	}

	return multiplier
}
