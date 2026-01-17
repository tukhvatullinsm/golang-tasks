package reflect_spell

import (
	//"fmt"
	"log"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	var result int64
	gameObjValue := reflect.ValueOf(object).Elem()
	if !gameObjValue.CanSet() {
		log.Fatalln("Object is not settable")
	}
	sourceValue := reflect.ValueOf(spell.Value())
	switch spell.Name() {
	case "fire":
		fieldValue := gameObjValue.FieldByName(spell.Char())
		if fieldValue.IsValid() && fieldValue.CanSet() {
			currentValue := fieldValue.Int()
			result = currentValue + sourceValue.Int()
			fieldValue.SetInt(result)
		}
	case "protect":
		fieldValue := gameObjValue.FieldByName(spell.Char())
		if fieldValue.IsValid() && fieldValue.CanSet() {
			currentValue := fieldValue.Int()
			result = currentValue + sourceValue.Int()
			fieldValue.SetInt(result)
		}
	default:
		log.Fatal("Unknown field")
	}

}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	// nolint: unused
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}
