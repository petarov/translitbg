package translitbg

import (
	"testing"
)

func TestAlphabet(t *testing.T) {
	expected := "abvgdezhziyklmnoprstufhtschshshtayyuyai"
	got, _ := New().Encode("абвгдежзийклмнопрстуфхцчшщъьюяѝ")

	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestSentences(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Всички хора се раждат свободни и равни по достойнство и права. Те са надарени с разум и съвест и следва да се отнасят помежду си в дух на братство.", "Vsichki hora se razhdat svobodni i ravni po dostoynstvo i prava. Te sa nadareni s razum i savest i sledva da se otnasyat pomezhdu si v duh na bratstvo."},
		{"Всички хора \nсе раждат свободни\n и равни по достойнство\n и права.", "Vsichki hora \nse razhdat svobodni\n i ravni po dostoynstvo\n i prava."},
		{"Ѝ може да бъде намерен и в други езици \nкато руския език и украинския език.", "I mozhe da bade nameren i v drugi ezitsi \nkato ruskia ezik i ukrainskia ezik."},
		{"СанЯ е на път за ЦаревеЦ-крепост и има МАНТИЯ от Желязо и Злато, носи ЖЕзъл и корона.", "SanYA e na pat za TsareveTS-krepost i ima MANTIA ot Zhelyazo i Zlato, nosi ZHEzal i korona."},
		// mixed
		{"Seдем с едin udар", "Sedem s edin udar"},
		{"Внезapna-smqна-на-$олаri-i-еüро", "Vnezapna-smqna-na-$olari-i-eüro"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For sentence '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestEdgeCases(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"България", "Bulgaria"},
		{"българия", "bulgaria"},
		{"БЪЛГАРИЯ", "BULGARIA"},
		{"БъЛгАриЯ", "BuLgAriA"},
		{"ЖЕЗЪЛ", "ZHEZAL"},
		{"жЕЗЪЛ", "zhEZAL"},
		{"жезъл", "zhezal"},
		{"ЦАРЕВЕЦ", "TSAREVETS"},
		{"Царевец", "Tsarevets"},
		{"ЦАРевец", "TSARevets"},
		{"цАРЕВЕЦ", "tsAREVETS"},
		{"чОвек", "chOvek"},
		{"ЧОВек", "CHOVek"},
		{"ШИВАЧ", "SHIVACH"},
		{"шИВАч", "shIVAch"},
		{"ЩАСТИЕ", "SHTASTIE"},
		{"щАСТИЕ", "shtASTIE"},
		{"ЩаСТИЕ", "ShtaSTIE"},
		{"ЮНГА", "YUNGA"},
		{"юНГА", "yuNGA"},
		{"ЯБЪЛКИ", "YABALKI"},
		{"яБЪЛКИ", "yaBALKI"},
		{"Всички ЖЕзъл, ЦаревеЦ, ЧОвек, шиваЧ, Щастие, юнга, Ябълки, месИЯ, ловеЧ",
			"Vsichki ZHEzal, TsareveTS, CHOvek, shivaCH, Shtastie, yunga, Yabalki, mesIA, loveCH"},
		{"месиЯта борЧвор маЦка маЦКите злоЩастието боЖе боЖЕ крЮгер кРЮгер крЮГер",
			"mesiAta borCHvor maTSka maTSKite zloSHTastieto boZHe boZHE krYUger kRYUger krYUGer"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For edge case '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestLocations(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Стара планина", "Stara planina"},
		{"Атанасовско езеро", "Atanasovsko ezero"},
		{"Централен Балкан", "Tsentralen Balkan"},
		{"София-юг", "Sofia-yug"},
		{"СофИя-юг", "SofIa-yug"},
		{"СофиЯ-юг", "SofiA-yug"},
		{"СофИЯ-ЮГ", "SofIA-YUG"},
		{"гр. София, ул. Тракия.", "gr. Sofia, ul. Trakia."},
		{"гр. СофИЯ, ул. ТракИя.", "gr. SofIA, ul. TrakIa."},
		{"Перник-север", "Pernik-sever"},
		{"Златни пясъци", "Zlatni pyasatsi"},
		{"Горна Оряховица", "Gorna Oryahovitsa"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For other name '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestCityNames(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"София", "Sofia"},
		{"Пловдив", "Plovdiv"},
		{"Варна", "Varna"},
		{"Бургас", "Burgas"},
		{"Русе", "Ruse"},
		{"Стара Загора", "Stara Zagora"},
		{"Плевен", "Pleven"},
		{"Сливен", "Sliven"},
		{"Добрич", "Dobrich"},
		{"Шумен", "Shumen"},
		{"Хасково", "Haskovo"},
		{"Перник", "Pernik"},
		{"Ямбол", "Yambol"},
		{"Благоевград", "Blagoevgrad"},
		{"Велико Търново", "Veliko Tarnovo"},
		{"Враца", "Vratsa"},
		{"Габрово", "Gabrovo"},
		{"Видин", "Vidin"},
		{"Монтана", "Montana"},
		{"Ловеч", "Lovech"},
		{"Разград", "Razgrad"},
		{"Силистра", "Silistra"},
		{"Търговище", "Targovishte"},
		{"Кюстендил", "Kyustendil"},
		{"Пазарджик", "Pazardzhik"},
		{"Смолян", "Smolyan"},
		{"Кърджали", "Kardzhali"},
		{"Велинград", "Velingrad"},
		{"Дупница", "Dupnitsa"},
		{"Петрич", "Petrich"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For city name '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func TestPeopleNames(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Самуил", "Samuil"},
		{"Синтия", "Sintia"},
		{"Марияна ИваноВа", "Mariana IvanoVa"},
		{"Явор", "Yavor"},
		{"Саня", "Sanya"},
		{"СанЯ", "SanYA"},
		{"САНЯ", "SANYA"},
		// 50 random names
		{"Иван", "Ivan"},
		{"Георги", "Georgi"},
		{"Мария", "Maria"},
		{"Димитър", "Dimitar"},
		{"Николай", "Nikolay"},
		{"Петър", "Petar"},
		{"Анна", "Anna"},
		{"Васил", "Vasil"},
		{"Стефан", "Stefan"},
		{"Елена", "Elena"},
		{"Александър", "Aleksandar"},
		{"Таня", "Tanya"},
		{"Стефания", "Stefania"},
		{"Виктория", "Viktoria"},
		{"Илия", "Ilia"},
		{"Даниел", "Daniel"},
		{"Михаил", "Mihail"},
		{"Радослав", "Radoslav"},
		{"Йордан", "Yordan"},
		{"Валентин", "Valentin"},
		{"Светлана", "Svetlana"},
		{"Христо", "Hristo"},
		{"Маргарита", "Margarita"},
		{"Надя", "Nadya"},
		{"Павел", "Pavel"},
		{"Виолета", "Violeta"},
		{"Симеон", "Simeon"},
		{"Румяна", "Rumyana"},
		{"Ивелина", "Ivelina"},
		{"Веселин", "Veselin"},
		{"Емилия", "Emilia"},
		{"Ангел", "Angel"},
		{"Стойко", "Stoyko"},
		{"Соня", "Sonya"},
		{"Любомир", "Lyubomir"},
		{"Магдалена", "Magdalena"},
		{"Анастасия", "Anastasia"},
		{"Красимир", "Krasimir"},
		{"Десислава", "Desislava"},
		{"Галина", "Galina"},
		{"Евгени", "Evgeni"},
		{"Росица", "Rositsa"},
		{"Кирил", "Kiril"},
		{"Дарина", "Darina"},
		{"Вера", "Vera"},
		{"Живка", "Zhivka"},
		{"Борис", "Boris"},
		{"Яна", "Yana"},
		{"Пламен", "Plamen"},
		{"Милена", "Milena"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For name '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}

func Test100Words(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"Здравей", "Zdravey"},
		{"Благодаря", "Blagodarya"},
		{"Моля", "Molya"},
		{"Добре", "Dobre"},
		{"Ден", "Den"},
		{"Нощ", "Nosht"},
		{"Човек", "Chovek"},
		{"Град", "Grad"},
		{"Страна", "Strana"},
		{"Живот", "Zhivot"},
		{"Любов", "Lyubov"},
		{"Семейство", "Semeystvo"},
		{"Вода", "Voda"},
		{"Храна", "Hrana"},
		{"Къща", "Kashta"},
		{"Работа", "Rabota"},
		{"Училище", "Uchilishte"},
		{"Университет", "Universitet"},
		{"Книга", "Kniga"},
		{"Музика", "Muzika"},
		{"Изкуство", "Izkustvo"},
		{"История", "Istoria"},
		{"Време", "Vreme"},
		{"Новина", "Novina"},
		{"Път", "Pat"},
		{"Автомобил", "Avtomobil"},
		{"Телефон", "Telefon"},
		{"Ресторант", "Restorant"},
		{"Летище", "Letishte"},
		{"Парк", "Park"},
		{"Животно", "Zhivotno"},
		{"Риба", "Riba"},
		{"Птица", "Ptitsa"},
		{"Цвете", "Tsvete"},
		{"Дърво", "Darvo"},
		{"Планина", "Planina"},
		{"Река", "Reka"},
		{"Слънце", "Slantse"},
		{"Луна", "Luna"},
		{"Звезда", "Zvezda"},
		{"Цвят", "Tsvyat"},
		{"Денс", "Dens"},
		{"Пеене", "Peene"},
		{"Гледам", "Gledam"},
		{"Слушам", "Slusham"},
		{"Говоря", "Govorya"},
		{"Спорт", "Sport"},
		{"Игра", "Igra"},
		{"Смях", "Smyah"},
		{"Природа", "Priroda"},
		{"Обичам", "Obicham"},
		{"Сърце", "Sartse"},
		{"Тяло", "Tyalo"},
		{"Ръка", "Raka"},
		{"Крак", "Krak"},
		{"Очи", "Ochi"},
		{"Уста", "Usta"},
		{"Коса", "Kosa"},
		{"Нос", "Nos"},
		{"Ухо", "Uho"},
		{"Език", "Ezik"},
		{"Зъб", "Zab"},
		{"Глава", "Glava"},
		{"Кожа", "Kozha"},
		{"Нокът", "Nokat"},
		{"Кръв", "Krav"},
		{"Въздух", "Vazduh"},
		{"Огън", "Ogan"},
		{"Земя", "Zemya"},
		{"Вода", "Voda"},
		{"Мъгла", "Magla"},
		{"Вятър", "Vyatar"},
		{"Сняг", "Snyag"},
		{"Дъжд", "Dazhd"},
		{"Топло", "Toplo"},
		{"Студено", "Studeno"},
		{"Храна", "Hrana"},
		{"Питие", "Pitie"},
		{"Спане", "Spane"},
		{"Бързо", "Barzo"},
		{"Бавно", "Bavno"},
		{"Силно", "Silno"},
		{"Слабо", "Slabo"},
		{"Голям", "Golyam"},
		{"Малък", "Malak"},
		{"Дълъг", "Dalag"},
		{"Кратък", "Kratak"},
		{"Тежък", "Tezhak"},
		{"Лек", "Lek"},
		{"Стар", "Star"},
		{"Млад", "Mlad"},
		{"Нов", "Nov"},
		{"Стар", "Star"},
		{"Добър", "Dobar"},
		{"Лош", "Losh"},
		{"Горещ", "Goresht"},
		{"Студен", "Studen"},
		{"Богат", "Bogat"},
		{"Беден", "Beden"},
		{"Висок", "Visok"},
		{"Нисък", "Nisak"},
		{"Дълбок", "Dalbok"},
		{"Плитък", "Plitak"},
	}

	tr := New()

	for _, tc := range testCases {
		output, _ := tr.Encode(tc.input)

		if output != tc.expected {
			t.Errorf("For word '%s', expected '%s', got '%s'", tc.input, tc.expected, output)
		}
	}
}
