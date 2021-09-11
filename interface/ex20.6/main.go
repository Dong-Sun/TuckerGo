package main

import "fmt"

type Attacker interface {
	Attack()
}

type Player struct {
	Lv int
}
type Monster struct {
	Lv int
}

func (m *Player) Attack() {
	fmt.Println("Player Attack")
}
func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()

		var monster *Monster
		monster = att.(*Monster) // error
		fmt.Println(monster.Lv)
	}
}

func main() {
	DoAttack(&Player{20})
}
