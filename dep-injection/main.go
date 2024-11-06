package main

import "fmt"

type SafetyPlacer interface {
	placeSafeties()
}

type RockClimber struct {
	rocksClimbed int
	kind         int
	sp           SafetyPlacer
}

type IceSafetyPlacer struct {
	// api
	// db
}

type NoSafetyPlacer struct{}

func (isp IceSafetyPlacer) placeSafeties() {
	fmt.Println("Placing ice safeties")
}

func (nsp NoSafetyPlacer) placeSafeties() {
	fmt.Println("No safeties")
}

func NewClimber(sp SafetyPlacer) *RockClimber {
	return &RockClimber{sp: sp}
}

func (rc *RockClimber) Climb() {
	rc.rocksClimbed++

	if rc.rocksClimbed == 10 {
		rc.sp.placeSafeties()
	}
}

// func (rc *RockClimber) placeSafeties() {
// 	fmt.Println("Placing safeties")
// }

func main() {
	climber := NewClimber(NoSafetyPlacer{})
	for i := 0; i < 11; i++ {
		climber.Climb()
	}
}
