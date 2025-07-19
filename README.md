# Pokemon CLI (Pokedex)

Інтерактивний командний додаток для дослідження світу Pokemon, ловлі покемонів та управління власним Pokedex.

## 🎮 Функціональність

- **Дослідження локацій** - знаходь нові місця та покемонів
- **Ловля покемонів** - використовуй покеболи для ловлі
- **Детальний огляд** - вивчай статистику пійманих покемонів  
- **Pokedex** - веди облік своєї колекції
- **Кешування** - швидкі повторні запити
- **Реалістична ловля** - рандомна логіка на основі складності покемона

## 🚀 Як запустити

### Запуск з джерельного коду
```bash
# Клонувати репозиторій
git clone [your-repo-url]
cd pokedexcli

# Запустити програму
go run .
```

### Збірка бінарника

#### Для поточної платформи
```bash
# Збудувати виконуваний файл
go build -o pokedex .

# Запустити
./pokedex              # Linux/macOS
# або
pokedex.exe            # Windows
```

#### Крос-компіляція для різних платформ
```bash
# Windows (64-bit)
GOOS=windows GOARCH=amd64 go build -o pokedex.exe .

# macOS (64-bit)
GOOS=darwin GOARCH=amd64 go build -o pokedex-mac .

# Linux (64-bit)
GOOS=linux GOARCH=amd64 go build -o pokedex-linux .

# macOS Apple Silicon (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o pokedex-mac-arm64 .
```

#### Оптимізована збірка (менший розмір)
```bash
go build -ldflags="-s -w" -o pokedex .
```

## 📋 Команди

### Навігація по світу
- `map` - показати наступні 20 локацій
- `mapb` - показати попередні 20 локацій  
- `explore <location-name>` - дослідити локацію та знайти покемонів

### Ловля та управління
- `catch <pokemon-name>` - спробувати піймати покемона
- `inspect <pokemon-name>` - детальна інформація про пійманого покемона
- `pokedex` - список всіх пійманих покемонів

### Допоміжні команди
- `help` - список всіх команд
- `exit` - вийти з програми

## 💡 Приклади використання

```bash
Pokedex > map
# Виводить список локацій...

Pokedex > explore pastoria-city-area
Exploring...pastoria-city-area
Found Pokemon:
 - tentacool
 - tentacruel
 - magikarp
 - gyarados

Pokedex > catch magikarp
Throwing a Pokeball at magikarp...
magikarp was caught!
You may now inspect it with the inspect command.

Pokedex > inspect magikarp
Name: magikarp
Height: 9
Weight: 100
Stats:
  -hp: 20
  -attack: 10
  -defense: 55
  -special-attack: 15
  -special-defense: 20
  -speed: 80
Types:
  - water

Pokedex > pokedex
Your Pokedex:
 - magikarp
```

## 🏗 Архітектура проекту

```
.
├── main.go                    # Головний файл з CLI логікою
├── internal/
│   ├── pokeapi/
│   │   ├── client.go         # HTTP клієнт для PokeAPI
│   │   └── types.go          # Структури даних
│   └── pokecache/
│       ├── cache.go          # Система кешування
│       └── cache_test.go     # Тести кешу
├── go.mod                    # Go модуль
└── README.md                 # Цей файл
```

## 🔧 Технічні деталі

### API
Використовує [PokeAPI](https://pokeapi.co/) для отримання даних про:
- Локації та області
- Інформацію про покемонів (статистика, типи, розміри)

### Кешування
- Автоматичне кешування всіх API запитів
- Час життя кешу: 5 хвилин
- Значно пришвидшує повторні запити

### Алгоритм ловлі
```go
// Формула ймовірності ловлі
catchChance := 50 + (300 - baseExperience) / 10
// Мінімум: 5%, Максимум: 95%
```

- Легкі покемони (низький `base_experience`) - легше піймати
- Важкі покемони (високий `base_experience`) - важче піймати

## 🛠 Технології

- **Go 1.21+** - мова програмування
- **PokeAPI** - джерело даних
- **Стандартна бібліотека Go** - HTTP клієнт, JSON парсинг
- **math/rand** - рандомна логіка ловлі

## 📚 Що вивчено

Під час розробки цього проекту було освоєно:

- **HTTP клієнти** в Go
- **JSON unmarshaling** для складних структур
- **Системи кешування** з TTL
- **CLI додатки** з інтерактивним вводом
- **Організація коду** в пакети
- **Обробка помилок** та валідація
- **Тестування** в Go

## 🎯 Можливі покращення

- [ ] Сортування списків покемонів
- [ ] Збереження Pokedex у файл
- [ ] Статистика ловлі (успішні/неуспішні спроби)
- [ ] Кольорний вивід в терміналі
- [ ] Конфігурація складності ловлі
- [ ] Додаткова інформація про покемонів (abilities, moves)

## 📄 Ліцензія

Цей проект створено з навчальною метою. Дані про Pokemon належать The Pokémon Company.
Ідея завдання отримана як частина навчального курсу сервісу boot.dev
Як допоміжний інструмент використовувалась IDE Cursor з ШІ Claude 

---

**Gotta catch 'em all!** 🏆 