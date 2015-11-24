package apidemic

import "strings"

var fieldTags = struct {
	Brand                     string
	Character                 string
	Characters                string
	CharactersN               string
	City                      string
	Color                     string
	Company                   string
	Continent                 string
	Country                   string
	CreditCardNum             string
	Currency                  string
	CurrencyCode              string
	Day                       string
	Digits                    string
	DigitsN                   string
	DomainName                string
	DomainZone                string
	EmailAddress              string
	EmailBody                 string
	FemaleFirstName           string
	FemaleFullName            string
	FemaleFullNameWithPrefix  string
	FemaleFullNameWithSuffix  string
	FemaleLastName            string
	FemaleLastNamePratronymic string
	FirstName                 string
	FullName                  string
	FullNameWithPrefix        string
	FullNameWithSuffix        string
	Gender                    string
	GenderAbrev               string
	HexColor                  string
	HexColorShort             string
	IPv4                      string
	Industry                  string
	JobTitle                  string
	Language                  string
	LastName                  string
	LatitudeDegrees           string
	LatitudeDirection         string
	LatitudeMinutes           string
	LatitudeSeconds           string
	Latitude                  string
	Longitude                 string
	LongitudeDegrees          string
	LongitudeDirection        string
	LongitudeMinutes          string
	LongitudeSeconds          string
	MaleFirstName             string
	MaleFullNameWithPrefix    string
	MaleFullNameWithSuffix    string
	MaleLastName              string
	MalePratronymic           string
	Model                     string
	Month                     string
	MonthNum                  string
	MonthShort                string
	Paragraph                 string
	Patagraphs                string
	PatagraphsN               string
	Password                  string
	Patronymic                string
	Phone                     string
	Product                   string
	ProductName               string
	Sentence                  string
	Sentences                 string
	SentencesN                string
	SimplePassWord            string
	State                     string
	StateAbbrev               string
	Street                    string
	StreetAddress             string
	Title                     string
	TopLevelDomain            string
	UserName                  string
	WeekDay                   string
	WeekDayShort              string
	WeekDayNum                string
	Word                      string
	Words                     string
	WordsN                    string
	Year                      string
	Zip                       string
}{
	"brand", "character", "characters", "characters_n",
	"city", "color", "company", "continent", "country",
	"credit_card_num", "currency", "currency_code", "day",
	"digits", "digits_n", "domain_name", "domain_zone",
	"email_address", "email_body", "female_first_name",
	"female_full_name", "female_full_name_with_prefix",
	"female_full_name_with_suffix", "female_last_name",
	"female_last_name_pratronymic", "first_name", "full_name",
	"full_name_with_prefix", "full_name_with_suffix", "gender",
	"gender_abrev", "hex_color", "hex_color_short", "i_pv_4",
	"industry", "job_title", "language", "last_name",
	"latitude_degrees", "latitude_direction", "latitude_minutes",
	"latitude_seconds", "latitude", "longitude", "longitude_degrees",
	"longitude_direction", "longitude_minutes", "longitude_seconds",
	"male_first_name", "male_full_name_with_prefix", "male_full_name_with_suffix",
	"male_last_name", "male_pratronymic", "model", "month",
	"month_num", "month_short", "paragraph", "patagraphs", "patagraphs_n",
	"password", "patronymic", "phone", "product", "product_name", "sentence",
	"sentences", "sentences_n", "simple_pass_word", "state", "state_abbrev",
	"street", "street_address", "title", "top_level_domain", "user_name", "week_day",
	"week_day_short", "week_day_num", "word", "words", "words_n", "year", "zip",
}

type Tags map[string]string

func (t Tags) Load(src string) {
	ss := strings.Split(src, ",")
	first := strings.TrimSpace(ss[0])

	if len(ss) < 2 {
		t["type"] = first
	}

	for _, v := range ss[0:] {
		ts := strings.SplitN(v, "=", 1)
		if len(ts) < 2 {
			t[v] = ""
			continue
		}
		t[strings.TrimSpace(ts[0])] = strings.TrimSpace(ts[1])
	}

}
