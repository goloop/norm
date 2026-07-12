# norm - довідник

`norm` нормалізує рядки до канонічної форми відомого типу й перевіряє результат
через `goloop/is`. Повний український довідник; англійською - [DOC.md](DOC.md).

## Зміст

- [Модель](#модель)
- [Типізовані нормалізатори](#типізовані-нормалізатори)
- [Правила по типах](#правила-по-типах)
- [Символьний toolkit](#символьний-toolkit)
- [Симетрія з is](#симетрія-з-is)
- [Межі](#межі)

## Модель

Кожен типізований нормалізатор повертає `(value string, ok bool)`: best-effort
очищене значення й чи воно валідне. Значення повертається навіть коли `ok`
false, щоб можна було підказати варіант.

Конвеєр кожного типізованого нормалізатора:

1. **Scrub**: прибрати керівні й невидимі символи (ніколи не валідні в цих
   форматах) і обрізати зовнішні пробіли.
2. Застосувати **безпечні канонічні** кроки (напр. lowercase регістро-
   незалежного домену чи схеми). Якщо результат валідний - повернути.
3. Застосувати обмежений **ремонт** (напр. прибрати внутрішні пробіли) лише
   якщо ще невалідно, і перевірити знову.
4. Повернути очищене значення з вердиктом `is`.

Оскільки результат завжди підтверджується предикатом `is`, значення ніколи не
позначається валідним без згоди `is`. Валідний ввід, зокрема email із лапкованою
локальною частиною, що легально містить пробіли, зберігається, бо ремонт
запускається лише коли м'якший результат невалідний.

## Типізовані нормалізатори

```go
func Email(s string) (string, bool)
func URL(s string) (string, bool)
func UUID(s string) (string, bool)
func MAC(s string) (string, bool)
func IBAN(s string) (string, bool)
func BankCard(s string, kinds ...is.CardKind) (string, bool)
func E164(s string) (string, bool)
func Hex(s string) (string, bool)
func MD5(s string) (string, bool)
func SHA1(s string) (string, bool)
func SHA256(s string) (string, bool)
func SHA512(s string) (string, bool)
```

## Правила по типах

- **Email**: trim; lowercase домену (після останнього `@`) зі збереженням
  локальної частини (домени регістро-незалежні, локальні частини технічно ні);
  якщо ще невалідно - прибрати внутрішні пробіли. IDN/punycode не конвертується.
- **URL**: trim; lowercase лише схеми й хоста. Path, query і fragment
  регістрозалежні й недоторкані. Відсутня схема не вигадується, тож
  `example.com` лишається невалідним.
- **UUID**: trim; зняти обрамлення `{...}` і префікс `urn:uuid:`; lowercase.
- **MAC**: перегрупувати в lowercase colon-пари. Приймає `:`, `-`, `.` і форму
  без роздільників; обробляє EUI-48 (12 hex) і EUI-64 (16 hex).
- **IBAN**: прибрати пробіли, uppercase, строга валідація (з контрольною сумою).
- **BankCard**: прибрати пробіли й дефіси, валідація Luhn і будь-якими заданими
  `is.CardKind`. **Безпека:** повернуте значення - повний номер картки (PAN); не
  логуйте й не зберігайте його без потреби, а Luhn-валідність не є доказом
  справжньої чи бажаної картки.
- **E164**: прибрати пробіли, дефіси, дужки й крапки. Без магії: національний
  номер без провідного `+` лишається невалідним.
- **Hex / MD5 / SHA1 / SHA256 / SHA512**: trim і lowercase (канонічна форма hex-
  дайджесту).

## Символьний toolkit

```go
type Class uint
const (
	Spaces Class = 1 << iota // Unicode White_Space (вкл. NBSP)
	Controls                 // Unicode Cc
	Invisible                // Cf + variation selectors + CGJ (zero-width/формат)
	Punct                    // Unicode P
	Symbols                  // Unicode S
	Letters                  // Unicode L
	Digits                   // Unicode Nd
)

func Trim(s string, class Class) string      // прибрати клас з обох кінців
func TrimLeft(s string, class Class) string
func TrimRight(s string, class Class) string
func Remove(s string, class Class) string    // прибрати клас всюди
func Keep(s string, class Class) string      // лишити тільки клас (інверсія Remove)
func CollapseSpace(s string) string          // серії пробілів -> один пробіл
func Lower(s string) string                  // Unicode нижній регістр
func Upper(s string) string                  // Unicode верхній регістр
```

Комбінуйте групи через `|`, напр. `norm.Remove(s, norm.Controls|norm.Invisible)`.
Нульовий `Class` - no-op для `Trim` і `Remove`.

## Симетрія з is

`norm` дзеркалить **форматні** предикати `is` (Email, URL, UUID, ...), де
«нормалізувати» має єдине безпечне значення: дати канонічну форму й перевірити.

Він свідомо **не** дзеркалить **властивісні** предикати (`is.Alpha`, `is.Digit`,
`is.Space`, ...). `norm.Alpha` натякав би на канонічну форму, а насправді
прибирав би не-літери - це втратна фільтрація. Замість цього - toolkit:
`norm.Keep(s, norm.Letters)` для випадку «лише літери», або `norm.Remove` /
`norm.Trim` для решти. Регістрові предикати відповідають простим перетворенням
`norm.Lower` і `norm.Upper`.

## Межі

`norm` чистить і валідує відомі типи й дає інструменти груп символів. Він не
парсить структуровані документи, не транслітерує (див. `goloop/t13n`), не робить
slug (див. `goloop/slug`) і не змінює регістр поза `Lower`/`Upper`.
