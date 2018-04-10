package apidemic

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/icrowley/fake"
	"github.com/satori/go.uuid"
)

type Value struct {
	Tags Tags
	Data interface{}
}

func (v Value) Update() Value {
	switch v.Data.(type) {
	case bool:
		return fakeString(&v)
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
	Data      map[string]Value
	IsArray   bool
	MaxLength int32
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
	return Value{Data: genFakeData(v)}
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

		if max, ok := v.Tags.Get("max"); ok {
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
	return Value{Data: genFakeData(v)}
}

func fakeObject(v *Value) Value {
	obj := NewObject()
	obj.Load(v.Data.(map[string]interface{}))
	return NewValue(obj.Data)
}

func genFakeData(v *Value) interface{} {
	if len(v.Tags) == 0 {
		return v.Data
	}

	typ, ok := v.Tags.Get("type")
	if !ok {
		return v.Data
	}
	switch typ {
	case fieldTags.Boolean:
		return genBool()
	case fieldTags.Brand:
		return fake.Brand()
	case fieldTags.Character:
		return fake.Character()
	case fieldTags.Characters:
		return fieldTags.Characters
	case fieldTags.CharactersN:
		max := 5
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return fake.CharactersN(max)
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
		vendor, _ := v.Tags.Get("vendor")
		fake.CreditCardNum(vendor)
	case fieldTags.Currency:
		fake.Currency()
	case fieldTags.CurrencyCode:
		fake.CurrencyCode()
	case fieldTags.Day:
		return fake.Day()
	case fieldTags.Decimal:
		min := 0
		max := 100
		prec := 2
		if m, err := v.Tags.Int("min"); err == nil {
			min = m
		}
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		if p, err := v.Tags.Int("prec"); err == nil {
			prec = p
		}
		return genDecimal(int32(min), int32(max)-1, prec)
	case fieldTags.Digits:
		return fake.Digits()
	case fieldTags.DigitsN:
		max := 5
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return fake.DigitsN(max)
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
	case fieldTags.Float:
		min := -100
		max := 100
		if m, err := v.Tags.Int("min"); err == nil {
			min = m
		}
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return genFloat(int32(min), int32(max))
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
	case fieldTags.Integer:
		min := -50
		max := 50
		if m, err := v.Tags.Int("min"); err == nil {
			min = m
		}
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return genInt(int32(min), int32(max))
	case fieldTags.JobTitle:
		return fake.JobTitle()
	case fieldTags.Language:
		return fake.Language()
	case fieldTags.LastName:
		return fake.LastName()
	case fieldTags.LatitudeDegrees:
		return fake.LatitudeDegrees()
	case fieldTags.LatitudeDirection:
		return fake.LatitudeDirection()
	case fieldTags.LatitudeMinutes:
		return fake.LatitudeMinutes()
	case fieldTags.LatitudeSeconds:
		return fake.LatitudeSeconds()
	case fieldTags.Latitude:
		return fake.Latitude()
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
		max := 5
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return fake.ParagraphsN(max)
	case fieldTags.Password:
		var (
			atLeast                                = 5
			atMost                                 = 8
			allowUpper, allowNumeric, allowSpecial = true, true, true
		)
		if least, err := v.Tags.Int("at_least"); err == nil {
			atLeast = least
		}
		if most, err := v.Tags.Int("at_most"); err == nil {
			atMost = most
		}
		if upper, err := v.Tags.Bool("upper"); err == nil {
			allowUpper = upper
		}
		if numeric, err := v.Tags.Bool("numeric"); err == nil {
			allowNumeric = numeric
		}
		if special, err := v.Tags.Bool("special"); err == nil {
			allowSpecial = special
		}
		return fake.Password(atLeast, atMost, allowUpper, allowNumeric, allowSpecial)
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
		max := 5
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return fake.SentencesN(max)
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
	case fieldTags.Uuid:
		return uuid.Must(uuid.NewV4())
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
		max := 5
		if m, err := v.Tags.Int("max"); err == nil {
			max = m
		}
		return fake.WordsN(max)
	case fieldTags.Year:
		from := 1930
		if f, err := v.Tags.Int("from"); err == nil {
			from = f
		}
		to := 2050
		if t, err := v.Tags.Int("to"); err == nil {
			to = t
		}
		return fake.Year(from, to)
	case fieldTags.Zip:
		return fake.Zip()
	}

	return v.Data
}

var seededRand = func() *rand.Rand {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r
}()

func genBool() bool {
	return seededRand.Float64() > 0.5
}

func genDecimal(min, max int32, prec int) string {
	b := genInt(min, max)
	t := fake.DigitsN(prec)

	res := fmt.Sprintf("%d.%s", b, t)

	return res
}

func genInt(min, max int32) int32 {
	var res int32
	diff := max - min

	if diff == 0 {
		return max
	} else if diff > 0 {
		res = seededRand.Int31n(diff) + min
	} else {
		res = (seededRand.Int31n(diff*-1) - min) * -1
	}

	return res
}

func genFloat(min, max int32) float64 {
	diff := max - min
	res := seededRand.Float64()*float64(diff) + float64(min)

	return res
}
