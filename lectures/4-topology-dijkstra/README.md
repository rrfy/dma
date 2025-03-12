# Графы - топологическая сортировка и пути

## Топологическая сортировка

Топологическая сортировка - упорядочивание вершин графа так, чтобы для любого ребра u -> v вершина u шла в порядке раньше v.

Если в графе есть цикл, то этот порядок не определён, поэтому алгоритм работает только с ацикличными направленными графами (directed acyclic graph или DAG).

Зачем она вообще? Вопрос хороший!

Предположим, что мы пишем компилятор, и вот нам надо собрать нашу программу, которая состоит из нескольких файлов. Их может быть немного (1-2, 3), где понятно в каком порядке их собирать:

<img width="334" alt="image" src="https://github.com/user-attachments/assets/e81194d9-86d3-4883-9ad6-528a3ce4de14" />

Но вот в больших проектах, с тысячами файлов будет тяжело сходу определить точный порядок. (на самом деле даже проект из 10 файлов тяжело собрать)

<img width="744" alt="image" src="https://github.com/user-attachments/assets/ddeeabc9-9c02-4f36-8d78-d4a9a567ab38" />

Этот порядок подскажет топологическая сортировка!

Задачу топологической сортировки можно решить несколькими способами, но сегодня рассмотрим Алгоритм Кана, который использует BFS. Кому интересно, можете почитать про Алгоритм Тарьяна (который использует DFS).

Алгоритм Кана выглядет примерно следующим образом:

1. Считаем входящие степени (in-degree) для всех вершин.
2. Добавляем вершины с in-degree = 0 в очередь.
3. Удаляем вершину из очереди, уменьшаем in-degree её соседей.
4. Если in-degree соседа стал 0, добавляем его в очередь.
5. Повторяем, пока очередь не пуста.

Давайте рассмотрим пример:

<img width="978" alt="image" src="https://github.com/user-attachments/assets/2cd1effb-2b36-483e-bf5a-dfdb9022c441" />

забираем из очереди первую вершину и добавляем её в результат. уменьшаем in-degree вершины 4, видим, что его значение in-degree стало 0, добавляем в очередь

<img width="970" alt="image" src="https://github.com/user-attachments/assets/b66fecc8-725a-4682-bc3f-88f9c4dcfc6e" />

забираем из очереди вершину 4 и добавляем её в результат. уменьшаем in-degree вершин 2 и 3, видим, что значение in-degree для 3 стало 0, добавляем в очередь

<img width="983" alt="image" src="https://github.com/user-attachments/assets/9c44deb2-dd68-41cc-8bcf-6712281fd409" />

забираем из очереди вершину 3 и добавляем её в результат. уменьшаем in-degree вершины 2, видим, что значение in-degree для 2 стало 0, добавляем в очередь

<img width="990" alt="image" src="https://github.com/user-attachments/assets/2cd44640-8bd7-4f53-8f29-747e59b09a21" />

забираем из очереди вершину 2 и добавляем её в результат. рёбер из неё нет. алгоритм заканчивается, потому что очередь опустела

<img width="975" alt="image" src="https://github.com/user-attachments/assets/e144602e-0b7c-4a4c-9203-3444f614e9fc" />

в итоге получаем следующий порядок вершин:

<img width="1005" alt="image" src="https://github.com/user-attachments/assets/a40c2502-d8d8-48f5-9acc-819bdfdbbd21" />

реализация

<img width="596" alt="image" src="https://github.com/user-attachments/assets/d91fccbb-3ca6-4f76-99b2-39884220b15c" />

## Пути в графах

Начнём с невзвешенных графов.

Нужно найти кратчайший путь от 1 до 2. Как это сделать?

<img width="475" alt="image" src="https://github.com/user-attachments/assets/c0d8b63c-9306-4331-ba4a-a402837e31a9" />

На самом деле мы можем ипользовать BFS для этого.

По определению BFS проходит по "слоям" графа, то есть сначала по вершинам на расстоянии 1 от начальной, потом - на расстоянии 2, 3 и т.д.

Как только мы попадаем в искомую вершину, возвращаем ответ.

<img width="455" alt="image" src="https://github.com/user-attachments/assets/f4bef01f-bcab-42c1-80ce-584b2dd75c22" />

Это хорошо, но как найти конкретный путь?

Для этого будем хранить массив parent[], где parent[v] = u означает, что в вершину v пришли из u.

После завершения BFS восстанавливаем путь от конечной вершины к начальной, используя parent.

<img width="793" alt="image" src="https://github.com/user-attachments/assets/11a5096e-d8e6-4ca8-af70-30431cb673c2" />

<img width="805" alt="image" src="https://github.com/user-attachments/assets/38cd7eee-c403-43d9-805d-93a528db8efd" />

<img width="805" alt="image" src="https://github.com/user-attachments/assets/8b0429ca-7013-4b4e-855e-33413da98f96" />

<img width="789" alt="image" src="https://github.com/user-attachments/assets/b3d0e2ad-a2cc-405c-aa3e-faf4035ad30e" />

<img width="700" alt="image" src="https://github.com/user-attachments/assets/a4c38aca-41c7-4278-bd13-3d72cb08d9d3" />

Ну а что со взвешенными?

С ними всё сложнее!

Самый "классический" алгоритм для поиска путей - алгоритм Дейкстры

* Зададим стартовую вершину s
* Создадим массив расстояний d, где d[v] - минимальное расстояние от s до v. Изначально d[s] = 0, а все остальные значения - плюс бесконечность.
* Создадим массив a[], где a[v] == True, если найдено минимальное расстояние от v до s, и False, если ещё нет
* Будем проводить n итераций, где на каждой выполняем:

  a. Извлекаем ещё не помеченную вершину с минимальным расстоянием из d

  b. Помечаем, что для выбранной вершины мы нашли минимальную длину до s, (обновляем массив a)

  c. Обновляем расстояния до её соседей (релаксация) следующим образом: d[u] = min(d[u], d[v] + w[v][u]), где w[v][u] - вес ребра {v, u}

<img width="673" alt="image" src="https://github.com/user-attachments/assets/a48617f2-3bc2-4354-9e5e-ffe6b233dea5" />

<img width="663" alt="image" src="https://github.com/user-attachments/assets/e7b2dbdd-f28d-4f21-99b7-03169ff08706" />

<img width="674" alt="image" src="https://github.com/user-attachments/assets/c92901b5-2045-4fdf-83cf-33d904697f63" />

<img width="675" alt="image" src="https://github.com/user-attachments/assets/0f2d7ec0-8c91-499d-85cc-b5ad75211d2f" />

<img width="658" alt="image" src="https://github.com/user-attachments/assets/73744cc5-7eb2-48ef-aee0-451cceb3e739" />

<img width="670" alt="image" src="https://github.com/user-attachments/assets/cf53225d-70ad-4f8c-8649-4e70567049ec" />

<img width="670" alt="image" src="https://github.com/user-attachments/assets/2489a8de-f277-4125-962a-272fe5833d8a" />

Реализация

```python
def dijkstra(graph, start):
    vertices = list(graph.keys())
    d = {v: float('inf') for v in vertices}
    d[start] = 0
    unvisited = set(vertices)

    while unvisited:
        v = None
        min_dist = float('inf')
        for u in unvisited:
            if d[u] < min_dist:
                min_dist = d[u]
                v = u
        if v is None:
            break

        unvisited.remove(v)

        for u, weight in graph[v]:
            if d[u] > d[v] + weight:
                d[u] = d[v] + weight

    return d
```

А как восстановить сам путь?

Примерно так же, как с BFS:

<img width="434" alt="image" src="https://github.com/user-attachments/assets/128812c9-8c14-4974-97e4-6f917de582a7" />

<img width="379" alt="image" src="https://github.com/user-attachments/assets/ddcbba22-92c4-44d9-9657-eafea5b6f8d7" />
