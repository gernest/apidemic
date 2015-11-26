package apidemic

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/icrowley/fake"
)

var helperTags = map[string]string{
	"age": "digitsN, max=2",
}

type Value struct {
	Tags Tags
	Data interface{}
}

func (v Value) Update() Value {
	switch v.Data.(type) {
	case string:
		return fakeString(&v)
	case float64:
		return fakeFloats(&v)
	case []interface{}:
		return fakeArray(&v)
	case map[string]interface{}:
		return fakeObject(&v)
	}
	return v
}

func (v Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Update().Data)
}

func NewValue(val interface{}) Value {
	return Value{Tags: make(Tags), Data: val}
}

type Object struct {
	Data map[string]Value
}

func NewObject() *Object {
	return &Object{Data: make(map[string]Value)}
}

func (o *Object) Load(src map[string]interface{}) error {
	for key, val := range src {
		value := NewValue(val)
		sections := strings.Split(key, ":")
		if len(sections) == 2 {
			key = sections[0]
			value.Tags.Load(sections[1])
		}
		o.Set(key, value)
	}
	return nil
}

func (o *Object) Set(key string, val Value) {
	o.Data[key] = val
}


func (v *Object) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.Data)
}

func parseJSONData(src io.Reader) (*Object, error) {
	var in map[string]interface{}
	err := json.NewDecoder(src).Decode(&in)
	if err != nil {
		return nil, err
	}
	o := NewObject()
	err = o.Load(in)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func fakeString(v *Value) Value {
	return Value{Data: genFakeData(v, "string")}
}

func fakeArray(v *Value) Value {
	arrV, ok := v.Data.([]interface{})
	if !ok {
		return *v
	}
	nv := *v
	var rst []interface{}
	if len(arrV) > 0 {
		origin := arrV[0]
		n := len(arrV)

		if max, ok := v.Tags["max"]; ok {
			maxV, err := strconv.Atoi(max)
			if err != nil {
				return *v
			}
			n = maxV
		}
		for i := 0; i < n; i++ {
			newVal := NewValue(origin)
			rst = append(rst, newVal)
		}
	}
	nv.Data = rst
	return nv
}

func fakeFloats(v *Value) Value {
	return Value{Data: genFakeData(v, "float")}
}

func fakeObject(v *Value) Value {
	obj := NewObject()
	obj.Load(v.Data.(map[string]interface{}))
	return NewValue(obj.Data)
}

func genFakeData(v *Value, kind string) interface{} {
	if len(v.Tags) == 0 {
		return v.Data
	}

	typ, ok := v.Tags["type"]
	if !ok {
		return v.Data
	}
	switch typ {
	case fieldTags.Brand:
		return fake.Brand()
	case fieldTags.Character:
		return fake.Character()
	case fieldTags.Characters:
		return fieldTags.Characters
	case fieldTags.CharactersN:
		return fake.CharactersN(5)
	case fieldTags.City:
		return fake.City()
	case fieldTags.Color:
		return fake.Color()
	case fieldTags.Company:
		return fake.Company()
	case fieldTags.Continent:
		return fake.Continent()
	case fieldTags.Country:
		return fake.Country()
	case fieldTags.CreditCardNum:
		fake.CreditCardNum("vendor")
	case fieldTags.Currency:
		fake.Currency()
	case fieldTags.CurrencyCode:
		fake.CurrencyCode()
	case fieldTags.Day:
		return fake.Day()
	case fieldTags.Digits:
		return fake.Digits()
	case fieldTags.DigitsN:
		return fake.DigitsN(5)
	case fieldTags.DomainName:
		return fake.DomainName()
	case fieldTags.DomainZone:
		return fake.DomainZone()
	case fieldTags.EmailAddress:
		return fake.EmailAddress()
	case fieldTags.EmailBody:
		return fake.EmailBody()
	case fieldTags.FemaleFirstName:
		return fake.FemaleFirstName()
	case fieldTags.FemaleFullName:
		return fake.FemaleFullName()
	case fieldTags.FemaleFullNameWithPrefix:
		return fake.FemaleFullNameWithPrefix()
	case fieldTags.FemaleFullNameWithSuffix:
		return fake.FemaleFullNameWithSuffix()
	case fieldTags.FemaleLastName:
		return fake.FemaleLastName()
	case fieldTags.FemaleLastNamePratronymic:
		return fake.FemalePatronymic()
	case fieldTags.FirstName:
		return fake.FirstName()
	case fieldTags.FullName:
		return fake.FullName()
	case fieldTags.FullNameWithPrefix:
		return fake.FullNameWithPrefix()
	case fieldTags.FullNameWithSuffix:
		return fake.FullNameWithSuffix()
	case fieldTags.Gender:
		return fake.Gender()
	case fieldTags.GenderAbrev:
		return fake.GenderAbbrev()
	case fieldTags.HexColor:
		return fake.HexColor()
	case fieldTags.HexColorShort:
		return fake.HexColorShort()
	case fieldTags.IPv4:
		return fake.IPv4()
	case fieldTags.Industry:
		return fake.Industry()
	case fieldTags.JobTitle:
		return fake.JobTitle()
	case fieldTags.Language:
		return fake.Language()
	case fieldTags.LastName:
		return fake.LastName()
	case fieldTags.LatitudeDegrees:
		return fake.LatitudeDegress()
	case fieldTags.LatitudeDirection:
		return fake.LatitudeDirection()
	case fieldTags.LatitudeMinutes:
		return fake.LatitudeMinutes()
	case fieldTags.LatitudeSeconds:
		return fake.LatitudeSeconds()
	case fieldTags.Latitude:
		return fake.Latitute()
	case fieldTags.LongitudeDegrees:
		return fake.LongitudeDegrees()
	case fieldTags.LongitudeDirection:
		return fake.LongitudeDirection()
	case fieldTags.LongitudeMinutes:
		return fake.LongitudeMinutes()
	case fieldTags.LongitudeSeconds:
		return fake.LongitudeSeconds()
	case fieldTags.MaleFirstName:
		return fake.MaleFirstName()
	case fieldTags.MaleFullNameWithPrefix:
		return fake.MaleFullNameWithPrefix()
	case fieldTags.MaleFullNameWithSuffix:
		return fake.MaleFullNameWithSuffix()
	case fieldTags.MaleLastName:
		return fake.MaleLastName()
	case fieldTags.MalePratronymic:
		return fake.MalePatronymic()
	case fieldTags.Model:
		return fake.Model()
	case fieldTags.Month:
		return fake.Month()
	case fieldTags.MonthNum:
		return fake.MonthNum()
	case fieldTags.MonthShort:
		return fake.MonthShort()
	case fieldTags.Paragraph:
		return fake.Paragraph()
	case fieldTags.Patagraphs:
		return fake.Paragraphs()
	case fieldTags.PatagraphsN:
		return fake.ParagraphsN(4)
	case fieldTags.Password:
	//return fake.Password()
	case fieldTags.Patronymic:
		return fake.Patronymic()
	case fieldTags.Phone:
		return fake.Phone()
	case fieldTags.Product:
		return fake.Product()
	case fieldTags.ProductName:
		return fake.ProductName()
	case fieldTags.Sentence:
		return fake.Sentence()
	case fieldTags.Sentences:
		return fake.Sentence()
	case fieldTags.SentencesN:
		return fake.SentencesN(4)
	case fieldTags.SimplePassWord:
		return fake.SimplePassword()
	case fieldTags.State:
		return fake.State()
	case fieldTags.StateAbbrev:
		return fake.StateAbbrev()
	case fieldTags.Street:
		return fake.Street()
	case fieldTags.StreetAddress:
		return fake.StreetAddress()
	case fieldTags.Title:
		return fake.Title()
	case fieldTags.TopLevelDomain:
		return fake.TopLevelDomain()
	case fieldTags.UserName:
		return fake.UserName()
	case fieldTags.WeekDay:
		return fake.WeekDay()
	case fieldTags.WeekDayNum:
		return fake.WeekdayNum()
	case fieldTags.WeekDayShort:
		return fake.WeekDayShort()
	case fieldTags.Word:
		return fake.Word()
	case fieldTags.Words:
		return fake.Words()
	case fieldTags.WordsN:
		return fake.WordsN(4)
	case fieldTags.Year:
	//return fake.Year()
	case fieldTags.Zip:
		return fake.Zip()
	}

	return v.Data
}
