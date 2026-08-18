# Глава 1 — Что такое pgx

## Что такое pgx

`pgx` — библиотека для работы с PostgreSQL через её API.

## pgx vs database/sql

Главное отличие в том, что `pgx` предоставляет полноценный API для непосредственной работы с PostgreSQL, а `database/sql` предоставляет универсальный API для работы с различными реляционными БД через драйверы.

## pgx.Conn

`pgx.Conn` — объект, представляющий одно соединение с БД и позволяющий выполнять операции через это соединение.

## pgxpool.Pool

`pgxpool.Pool` — пул соединений, который управляет несколькими соединениями с БД, позволяя использовать их повторно и распределять между запросами.

## pgxpool.Pool vs pgx.Conn

Главное отличие — `pgx.Conn` представляет одно соединение с PostgreSQL, а `pgxpool.Pool` управляет несколькими соединениями.

В production-приложениях обычно используют пул соединений, чтобы несколько запросов могли выполняться одновременно, а существующие соединения можно было повторно использовать.

Это позволяет эффективнее работать с нагрузкой и не создавать новое соединение для каждого запроса.

## Native API

Native API — это непосредственное использование API `pgx` для работы с PostgreSQL без промежуточного слоя `database/sql`.

## QueryRow

`QueryRow` (Query — запрос, Row — строка) — метод для выполнения запроса, который должен вернуть одну строку.

Результат запроса извлекается с помощью метода `Scan()`.

```go
var result int

err := pool.QueryRow(ctx, "SELECT 1").Scan(&result)
```

## Exec

`Exec` — метод для выполнения SQL-запросов, от которых мы не ожидаем получение строк в результате.

Обычно используется для:

- `INSERT`
- `UPDATE`
- `DELETE`
- `CREATE`
- других SQL-команд, не возвращающих строки.

`Exec` возвращает `CommandTag` и ошибку.

`CommandTag` содержит информацию о выполненной SQL-команде. С помощью метода `RowsAffected()` можно получить количество затронутых строк.

```go
// Обновляем пользователя
tag, err = pool.Exec(ctx, queryUpdate)
if err != nil {
	fmt.Println("Can't update user:", err)
	return
}
fmt.Println("Rows affected:", tag.RowsAffected())
```

## Query

Query — метод для выполнения SQL-запроса, результат которого может содержать несколько строк.

Метод возвращает объект Rows. После получения Rows его необходимо закрыть через defer, чтобы корректно освободить связанные с результатом ресурсы.

```go
rows, err := pool.Query(ctx, query)
if err != nil {
	return err
}
defer rows.Close()
```

Для чтения результата используется следующий алгоритм:

`Next()` — переходит к следующей строке результата.
`Scan()` — считывает значения текущей строки в переданные переменные.
`Err()` — проверяет, возникла ли ошибка во время чтения результата.

Основной паттерн:

```go
for rows.Next() {
	if err := rows.Scan(&id, &name); err != nil {
		return err
	}
}

if err := rows.Err(); err != nil {
	return err
}
```

## PostgreSQL в Docker

В production и при локальной разработке PostgreSQL часто запускают в Docker-контейнере.

Основные элементы:

- **Docker image** — образ PostgreSQL, например `postgres:18.1-bookworm`;
- **container** — запущенный экземпляр образа;
- **volume** — постоянное хранилище данных PostgreSQL;
- **port mapping** — проброс порта контейнера на хост;
- **environment variables** — настройки PostgreSQL при первом запуске.

Пример:

```bash
docker run -d \
  --name pgx-postgres \
  -p 5432:5432 \
  --env-file .env \
  -v pgx-data:/var/lib/postgresql \
  postgres:18.1-bookworm
```

## Graceful Shutdown

Graceful Shutdown — корректное завершение приложения: мы получаем сигнал о завершении, перестаём принимать новые запросы, даём текущим операциям завершиться и после этого закрываем ресурсы.

1. Ждём сигнал завершения
```go
ctx, cancel := signal.NotifyContext(
    context.Background(),
    syscall.SIGINT,
    syscall.SIGTERM,
)
defer cancel()

go func() {
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Can't starting server", err)
		cancel()
	}
}()

<-ctx.Done()
```

Идея: SIGINT / SIGTERM → context отменяется → ctx.Done() разблокируется → понимаем, что пора завершаться.

2. Выполняем Shutdown
```go
shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

if err := server.Shutdown(shutdownCtx); err != nil {
	fmt.Println("Can't shutdown server", err)
	return
}
```

Идея: создаём отдельный context с timeout и через него управляем временем graceful shutdown.