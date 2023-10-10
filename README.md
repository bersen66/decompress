# decompress

Создать Go-функцию, осуществляющую примитивную распаковку
строки, содержащую повторяющиеся символы/руны, например:
* ```"a4bc2d5e" => "aaaabccddddde"```
* ```"abcd" => "abcd"```
* ```"45" => "" (некорректная строка)```
* ```"" => ""```
  Реализовать поддержку escape-последовательностей.
  Например:
  
 * ```qwe\4\5 => qwe45``` (\*)
 * ```qwe\45 => qwe44444``` (\*)
 * ```qwe\\5 => qwe\\\\\``` (\*)
 
 В случае если была передана некорректная строка, функция
 должна возвращать ошибку. Написать unit-тесты.