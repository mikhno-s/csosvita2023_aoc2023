# Одяг
Іван любить шопінг. Якось він загорівся ідеєю підібрати собі футболку і штани так, щоб виглядати в них максимально стильно. У розумінні Івана стильність одягу тим більша, чим менша різниця у кольорі елементів його одягу.

В наявності є N футболок та M штанів, про кожен елемент відомий його колір (ціле число від 1 до 10 000 000). Допоможіть Івану вибрати одну футболку і одні штани так, щоб різниця в їхньому кольорі була щонайменша.

## Input Format

Спочатку вводиться інформація про футболки: у першому рядку ціле число N і в другому N цілих чисел від 1 до 10 000 000 - кольори наявних футболок. Гарантується, що номери кольорів подаються в зростаючому порядку та є унікальними.

Далі в тому ж форматі йде опис штанів: їх кількість M та кольори - M цілих чисел від 1 до 10 000 000 у зростаючому порядку.

## Constraints

1 ≤ N ≤ 100 000, 1 ≤ M ≤ 100 000

## Output Format

Виведіть пару невід'ємних чисел – колір майки та колір штанів, які слід обрати Івану. Якщо варіантів вибору є кілька, виведіть перший зліва.

## Sample Input 0

```
2
3 4
3
1 2 3
```

## Sample Output 0

```
3 3
```

## Sample Input 1

```
2
4 5
3
1 2 3
```

## Sample Output 1

```
4 3
```

## Sample Input 2

```
5
1 2 3 4 5
5
1 2 3 4 5
```

## Sample Output 2

```
1 1
```