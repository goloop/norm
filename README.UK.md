[![Go Reference](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/goloop/norm) [![License](https://img.shields.io/badge/license-MIT-brightgreen)](https://github.com/goloop/norm/blob/v1/LICENSE) [![Stay with Ukraine](https://img.shields.io/static/v1?label=Stay%20with&message=Ukraine%20♥&color=ffD700&labelColor=0057B8&style=flat)](https://u24.gov.ua/)

# norm

`norm` нормалізує введені користувачем рядки. Це «пишучий» компаньйон до
[`goloop/is`](https://github.com/goloop/is): де `is` читає (чи це email?), `norm`
чистить значення до канонічної форми цього типу й перевіряє результат.

```go
value, ok := norm.Email("  Example @ Gmail.com ")
// value == "Example@gmail.com", ok == true
```

Лише стандартна бібліотека плюс `goloop/is` для валідації.

## Встановлення

```bash
go get github.com/goloop/norm
```

## Типізовані нормалізатори

Кожен дзеркалить предикат `is` і повертає `(очищене, валідне)`. Очищене
значення повертається навіть коли невалідне, щоб UI міг підказати «можливо, ви
мали на увазі X?».

| Функція | Нормалізація |
|---------|--------------|
| `Email` | trim, lowercase домену (локальну частину збережено), ремонт внутрішніх пробілів |
| `EmailFold` | `Email`, тоді lowercase усієї адреси (регістро-незалежний ключ-ідентичність) |
| `URL` | trim, lowercase лише схеми й хоста (схема не вигадується) |
| `UUID` | trim, зняти `{}` і `urn:uuid:`, lowercase |
| `MAC` | перегрупувати в lowercase colon-пари (приймає `:`, `-`, `.`, без) |
| `IBAN` | прибрати пробіли, uppercase, строга контрольна сума |
| `BankCard` | прибрати пробіли й дефіси, Luhn (і опційні види) |
| `E164`, `Phone` | прибрати пробіли, дужки, дефіси, крапки (код країни не вигадується) |
| `IMSI` | прибрати пробіли й дефіси |
| `IP`, `IPv4`, `IPv6` | канонічна форма (стиснений, lowercase IPv6) |
| `Domain`, `Hostname` | trim, lowercase |
| `Base64`, `Base64URL` | прибрати пробіли (розгорнути) |
| `JWT` | trim, scrub |
| `HexColor` | trim, lowercase (`#` не вигадується) |
| `RGBColor` | lowercase ключового слова `rgb` |
| `Bin` | прибрати пробіли, lowercase префікса `0b` |
| `Decimal` | прибрати пробіли (розгрупувати цифри) |
| `Float` | лише trim і scrub (граматику не ремонтуємо) |
| `Hex`, `MD5`, `SHA1`, `SHA256`, `SHA512` | trim, lowercase |

Кожне значення спершу очищається від керівних і невидимих символів (zero-width
пробіл в email прибирається до валідації). Значення ніколи не позначається
валідним, якщо `is` не згоден, і `norm` ніколи не прибирає довільну пунктуацію
й не вигадує дані.

## Символьний toolkit

Для вводу, що не є фіксованим зовнішнім форматом, є інструменти груп символів.
`Class` - це бітова маска (комбінуйте через `|`):

```go
norm.Trim(s, norm.Spaces|norm.Controls) // прибрати групи з обох кінців
norm.Remove(s, norm.Invisible)          // прибрати групу всюди
norm.Keep(s, norm.Letters|norm.Digits)  // лишити тільки ці групи
norm.CollapseSpace(s)                   // серії пробілів -> один пробіл
norm.Clean(s)                           // прибрати невидимі/контрол, стиснути пробіли, тримити
norm.Lower(s), norm.Upper(s)            // Unicode-регістр
```

`Clean` - очищення в один виклик для вільного текстового поля (ім'я, заголовок,
коментар): прибирає невидимі й non-space контрол-символи, стискає кожну серію
пробілів до одного й тримить краї, тож вставлений текст не занесе в сховище
zero-width байти, приховані переноси рядків чи рвану розбивку.

Групи: `Spaces`, `Controls`, `Invisible` (zero-width / формат-символи, BOM,
м'який дефіс, variation selectors), `Punct`, `Symbols`, `Letters`, `Digits`
(десяткові), `Numbers` (усі Unicode-числа).

Для character-class предикатів `is` чесні втратні відповідники названо `*Only`
і повертають простий рядок (не `(string, bool)`), щоб ім'я сигналізувало
фільтрацію, а не канонізацію:

```go
norm.AlphaOnly("a1-b2")   // "ab"
norm.AlnumOnly("a1-b2!")  // "a1b2"
norm.DigitsOnly("+1 415") // "1415"
```

Свідомо немає `norm.Alpha`/`norm.Digit` (ім'я, що натякало б на канонічну форму
для втратного фільтра), і немає дзеркал для граматичних чи політикових
предикатів (`Numeric`, `Title`, `Nickname`, `VariableName`, `SelectorName`), де
фільтрація змінила б сенс чи ідентичність.

## Документація

- Англійський довідник: [DOC.md](DOC.md)
- Український довідник: [DOC.UK.md](DOC.UK.md)

## Ліцензія

MIT - див. [LICENSE](LICENSE).

## Реліз v1.0.2

IP, IPv4 та IPv6 перевіряють і канонізують адресу одним розбором netip.
