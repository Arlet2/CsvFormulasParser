# CSV Formula Parser

## Тех.задание
Задан CSV-файл (comma-separated values) с заголовком, в котором перечислены названия столбцов. Строки нумеруются
целыми положительными числами, необязательно в порядке возрастания. В ячейках CSV-файла могут хранится или целые
числа или выражения вида
= ARG1 OP ARG2

где ARG1 и ARG2 – целые числа или адреса ячеек в формате Имя_колонки Номер_строки, а OP – арифметическая операция
из списка: +, -, *, /.
Например, таблица

|    | A | B          | Cell |   |
|----|---|------------|------|---|
| 1  | 1 | 0          | 1    |   |
| 2  | 2 | =A1+Cell30 | 0    |   |
| 30 | 0 | =B1+A1     | 5    |   |

Будет представлена в нашем CSV-формате следующим образом:

,A,B,Cell
1,1,0,1
2,2,=A1+Cell30,0
30,0,=B1+A1,5

(обратите внимание на пропуск первого значения в первой строке CSV-представления, он обозначает пустую левую верхнюю
ячейку таблицы).
Требуется написать программу, которая читает произвольную CSV-форму из файла (количество строк и столбцов может быть
любым), вычисляет значения ячеек, если это необходимо, и выводит получившуюся табличку в виде CSV-представления в
консоль.

## Сборка проекта
Проект можно собрать, активировав скрипт pkg.sh в корневой директории (будет две сборки под Windows и Linux).

## Тестирование
Тесты всей программы представлены в папке cmd, в файле main_test.go. Тестовые файлы csv (включая пример из тз) представлены в папке cmd/test_csvs

Юнит-тесты представлены в пакетах, в которых они тестируются, в файлах <package_name>_test.go.

Запустить все юнит-тесты и тесты программы можно, запустив скрипт test.sh.

## Идея решения
Для вычисления формул использовал поиск в глубину и топологическую сортировку ориентированного ациклического графа.
