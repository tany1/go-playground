package main

import "fmt"

type Position struct {
	x, y float64
}

type SpecialPosition struct {
	Position
}

func (sp *SpecialPosition) MoveSpecial(x, y float64) {
	sp.x += x * 2
	sp.y += y * 2
}

func (p *Position) Move(x, y float64) {
	p.x += x
	p.y += y
}

func (p *Position) Teleport(x, y float64) {
	p.x = x
	p.y = y
}

type Player struct {
	*Position
}

func NewPlayer() *Player {
	return &Player{Position: &Position{}}
}

type Enemy struct {
	*SpecialPosition
}

func NewEnemy() *Enemy {
	return &Enemy{SpecialPosition: &SpecialPosition{}}
}

func main() {
	raidBoss := NewEnemy()

	raidBoss.Move(1.0, 2.0)
	raidBoss.MoveSpecial(1.0, 2.0)
	fmt.Println(raidBoss.Position)

	player := NewPlayer()

	player.Move(1.34, 2.45)
	fmt.Println(player.Position)

	player.Move(-1.0, 12)
	fmt.Println(player.Position)

	player.Teleport(3.45, 4.56)
	fmt.Println(player.Position)
}
