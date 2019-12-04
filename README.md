[![Build Status](https://travis-ci.org/Serabass/biner.svg?branch=master)](https://travis-ci.org/Serabass/biner)

_документ устарел, обновится позже, когда проработаю основные детали_

# Biner

**Biner** - это моя первая попытка разработать язык.
Назвать это языком программирования пока язык не поворачивается (уж простите за тавтологию), однако фишки вроде описания структур и импортов уже реализованы.

## Для чего он нужен

**Biner** - это, в первую очередь, язык описания структур двоичных данных. На текущий момент имеет встроенные основные типы:

- int8
- int16
- int32
- float32
- float64
- char как аналог int8
- некоторые строковые структуры (опишу их позже)

## Демо

Все примеры буду складывать в папке [examples](./examples)

## Где он может пригодиться

Я задумывал **Biner** как инструмент для описания и разбора двоичных блоков памяти, коими могут быть области в оперативной памяти или файлах.

Предположим, у нас есть три байта с данными о красном цвете RGB: `FF 00 00`. Мы можем описать для него простую структуру:

```go
struct rgb {
  r: int8;
  g: int8;
  b: int8;
}

```

Таким образом, каждый из полученных байт считается один за другим и пропишется в нужные поля на выходе. А на выходе мы получим JSON-объект:

```json
{
  "r": 255,
  "g": 0,
  "b": 0
}
```

## Процесс разработки
Я предлагаю использовать подход TDD []

## Что есть

### Вложенность структур

_описать_

### Модульность

_описать_

### Константы

_описать_

### Директивы

_описать_

### Простейшая логика при чтении данных

Возьмём вот такой пример из тестов:

```go
struct {
  color: rgb {
    when r == 0xFF {
      red = true;
    }
    when g == 0xFF {
      green = true;
    }
    when b == 0xFF {
      blue = true;
    }
  }
}
```

Ключевое слово `when` определяет, что делать если проходит условие.
В данном случае мы считали структуру `rgb`, а затем обрабатываем полученные значения через блоки `when`.
/
Таким образом, прочитав набор байт вида `00 FF 00`, мы должны получить вот такую структуру:

```json
{
  "r": 0,
  "g": 255,
  "b": 0,
  "green": true
}
```

Последнее поле прописывается потому что прошло второе условие `when`.

## Что планируется

### Более гибкая работа со скалярными значениями

Сейчас всё довольно топорно сделано.

### Перегрузка типов

_описать_

### Запись блоков памяти обратно в буфер

Сейчас **Biner** умеет только читать.

_описать_

### Плагины для IDE

### Линтеры

### Варианты названия

* BiSt = от Binary Structurization (созвучно с Beast)


https://www.opennet.ru/opennews/art.shtml?num=46244
