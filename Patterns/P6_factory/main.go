/*
Паттерн фабричный метод - это порождающий паттерн проектирования, позволяющий
создавать объекты различных типов без явного указания типа в коде-породителе.
Применимость:
	- Заранее неизвестны типы объектов, с которыми будет работать определённый
	участок кода.
Плюсы:
	- Так как код не зависит от конкретного типа, он становится более гибким.
	Этот паттерн помогает реализовать принцип открытости/закрытости:
	при необходимости добавить поддержку объектов новых типов, мы не меняем
	код бизнес логики. В коде бизнес-логики нас волнует только поведение(методы),
	а всё "ветвление" остаётся в коде конструктора.
	- Нет необходимости прописывать общие методы для каждого отдельного типа
Минусы:
	- Повышение уровня абстракции. Если типов объектов немного,
		логика в виде if-else или switch-case может выглядеть понятнее.
	- Если реализовывать "как в ООП-языках" (имитируя наследование), повышается
		уровень абстракции, увеличивается количество кода.
*/

/*
Пример реализиции основан на следующем кейсе.
Есть некий сервис аналитики, задачей которого является сбор и агрегация информации
о согласовании задач. Сервис получает информацию от брокера сообщений и записывает
данные в БД. Данные представляют собой события типа "задача создана",
"задача отправлена на согласование согласующему А", "задача отклонена согласующим Б",
"задача полностью согласована" и т.д.
Для получения данных сервис реализует REST API.
На данном этапе от сервиса требуется выдавать только информацию о количестве
полностью согласованных/не согласованных задач и общем времени согласования
конкретной задачи.
Запрашиваемые данные можно объединить под общим названием "агрегат".
По мере использования сервиса планируется увеличение количества агрегатов.
В зависимости от запроса создаётся тот или иной тип агрегата.
Коду, который работает с агрегатом, не важен его тип.
*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// AggregateBehaviour - общий для всех типов интерфейс
type AggregateBehaviour interface {
	Get() int
	Save() error
}

// Aggregate - "тип-родитель" для других типов.
// Предполагается, что все агрегаты будут "наследовать" этот тип.
type Aggregate struct {
	Name  string
	Value int
}

// Signed "наследует" Aggregate через копирование
type Signed = Aggregate

// Unsigned "наследует" Aggregate через встраивание
type Unsigned struct {
	Aggregate
}

type TotalTime struct {
	Aggregate
	ID string
}

func (a Aggregate) Get() int {
	val := getFromDb(a)

	return val
}

func (a Aggregate) Save() error {
	return saveToDB(a)
}

// NewSigned - конструктор объектов типа Signed.
func NewSigned() Signed {
	var s Signed

	s.Name = "signed"

	return s
}

// NewUnsigned - конструктор объектов типа Unsigned
func NewUnsigned() Unsigned {
	var u Unsigned

	u.Name = "unsigned"

	return u
}

// NewSignitionTime - конструктор объектов типа TotalTime
func NewSignitionTime(r *http.Request) TotalTime {
	var t TotalTime

	t.Name = "total_time"
	t.ID = r.URL.Query().Get("id")

	return t
}

func NewAggregate(r *http.Request) AggregateBehaviour {
	aggrType := r.FormValue("type")

	switch {
	case aggrType == "signed":
		return NewSigned()
	case aggrType == "unsigned":
		return NewUnsigned()
	case aggrType == "total_time":
		return NewSignitionTime(r)
	default:
		log.Println("unknown type")
		return nil
	}
}

func main() {
	url1 := `http://localhost/aggregate?type=signed`
	url2 := `https://localhost:80/aggregate?type=unsigned`
	url3 := `http://localhost:8081/aggregate?type=total_time&id=some_id"`

	U1, _ := url.Parse(url1)
	U2, _ := url.Parse(url2)
	U3, _ := url.Parse(url3)

	r1 := http.Request{URL: U1}
	r2 := http.Request{URL: U2}
	r3 := http.Request{URL: U3}

	aggregate1 := NewAggregate(&r1)
	if err := aggregate1.Save(); err != nil {
		log.Println("unable save aggregate")
	} else {
		log.Println(aggregate1, "saved successfully")
		fmt.Println(aggregate1.Get())
	}

	aggregate2 := NewAggregate(&r2)
	if err := aggregate2.Save(); err != nil {
		log.Println("unable save aggregate")
	} else {
		log.Println(aggregate2, "saved successfully")
		fmt.Println(aggregate2.Get())
	}

	aggregate3 := NewAggregate(&r3)
	if err := aggregate3.Save(); err != nil {
		log.Println("unable save aggregate")
	} else {
		log.Println(aggregate3, "saved successfully")
		fmt.Println(aggregate3.Get())
	}
}

func getFromDb(a Aggregate) int {
	return a.Value * 10
}

func saveToDB(a Aggregate) error {
	if a.Value < 0 {
		return fmt.Errorf("error. No negative values allowed")
	}

	return nil
}
