[![Build Status](https://travis-ci.org/Serabass/biner.svg?branch=master)](https://travis-ci.org/Serabass/biner)

# Biner

**Biner** (рабочее название) - это моя первая попытка разработать язык.
Назвать это языком программирования пока язык не поворачивается (уж простите за тавтологию), однако фишки вроде структур, наследования, импортов и многого другого уже реализованы.

Пока что будем называть его языком описания.

## Для чего он нужен

**Biner** - это, в первую очередь, язык описания структур двоичных данных.

На текущий момент имеет встроенные типы:

- int8
- int16
- int32
- uint8
- uint16
- uint32
- float32
- float64
- char как аналог int8
- некоторые строковые структуры (опишу их позже)

Все эти типы данных должны читаться стандартными функциями класса Buffer.

## Демо

Все примеры буду складывать в папке [examples](./examples)

## Как начать разработку

1. Форкаем к себе репку
2. Клонируем
3. `$ yarn`
4. `$ yarn tw`

Запустятся все тесты, которые можно отфильтровать (см. документацию Jest). Я обычно запускаю `sandbox.tests.ts` и параллельно играюсь с парсером в `examples/sandbox.go`. 

Основной код парсера лежит в `/src/biner-work.pegjs` (за подробностями ко мне или в [документацию PEGjs](https://pegjs.org/documentation))

Пока большинство тестов не проходят, да и написаны криво. К тому же на текущий момент я полностью сломал класс Processor, потому как для столь сложной разработки нужно выполнять большие задачи последовательно: сначала тщательно продумываем синтаксис и работаем только с парсером, а затем навешиваем логику.

Иначе говоря, чем больше примеров будет в `/examples`, тем меньше подводных камней мы встретим в ходе разработки.

**Важно:** не стесняйтесь задавать вопросы. Таким образом я буду знать, что ещё внести в этот README, чтобы он был максимально информативным.

## Где он может пригодиться

Я задумывал **Biner** как инструмент для описания и разбора двоичных блоков памяти, коими могут быть области в оперативной памяти или файлах.

### Простейший пример

Предположим, у нас есть три байта с данными о красном цвете RGB: `FF 00 00`. Мы можем описать для него простую структуру:

```go
struct rgb {
  r: int8;
  g: int8;
  b: int8;
}

```

Таким образом, каждый из полученных байтов будет считан один за другим и прописан в нужные поля на выходе. А на выходе мы получим JSON-объект:

```json
{
  "r": 255,
  "g": 0,
  "b": 0
}
```

## Процесс разработки
Я предлагаю использовать подход [TDD](https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D0%B7%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D0%BA%D0%B0_%D1%87%D0%B5%D1%80%D0%B5%D0%B7_%D1%82%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5) с использованием Jest и его замечательного `--watch`.

## Что под капотом
* [Node.js](https://nodejs.org/)
* [Typescript](https://www.typescriptlang.org/)
* [TSLint](https://palantir.github.io/tslint/)
* [Jest](https://jestjs.io/)
* [PEGjs](https://pegjs.org/) (Его важно изучить. У них в репке есть [несколько замечательных примеров](https://github.com/pegjs/pegjs/tree/master/examples))
* [Travis CI](https://travis-ci.org/)
* 

Пока что я использую расширение `.go` для примеров, потому как его подсветка очень близка к концепции разрабатываемого языка.

## Kaitai

Есть аналогичный проект: [Kaitai](http://kaitai.io/), однако там структуры описываются на языке YAML. Этот язык тоже хорош, да и проект довольно объёмный:
 * Множество готовых примеров (некоторые из которых, я, к слову, у них подсмотрел)
 * Онлайн-песочница
 * Хорошая документация
 * Возможность генерировать парсеры для разных языков (это вообще огонь)

Имею смелость предположить, что стоит устроить с ними коллаборацию и обменяться, например конвертерами форматов (Biner <-> Kaitai). Ну и получить опыт у вполне состоявшегося проекта. YAML, конечно, хорош в определённых задачах, но лично мне кажется, что он слишком громоздок для описания таких структур. В случае же Biner хочется сделать максимально человекочитаемый язык, с которым будет проще разобраться программисту, владеющему любым Си-подобным языком (Biner всё же можно считать Си-подобным, я так полагаю).

## Что есть

### Структуры

_описать_

### Скалярные типы

_описать_

### Условные блоки (пока только switch)

_описать_

### Модульность

(Это песня про импорты/экспорты)

### Константы

_описать_

### Директивы

_описать_

### Декораторы

Я позже опишу чем они могут быть здесь полезны. Однако саму концепцию я пока не продумал.

### Наследование

Структуры можно наследовать от других структур

### Generics

Замечательная концепция во многих языках программирования. К слову, парсер уже обучен их понимать.

## Что планируется

### Простейшая арифметика

На первых этапах нужно реализовать хотя бы основные арифметические операции (+-*/).

### Запись блоков памяти обратно в буфер

На первых этапах мы будем учить Biner только читать данные, но в дальнейшем нужно будет заложить и запись. Я пока смутно представляю себе это, но задача не должна оказаться нерешаемой.

### Плагины для IDE / Линтеры

Подсветка, анализ кода и прочие плюшки

## Варианты названия

Нужно решить, как обозвать проект. Сейчас используется рабочее название Biner, но мне оно не нравится.

* BiSt = от Binary Structurization (созвучно с Beast)
* BSLang - Binary Structurization Language
* ...предлагайте

https://www.opennet.ru/opennews/art.shtml?num=46244
