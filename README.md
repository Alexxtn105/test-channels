# Тестовый пример использования каналов
## Небуферизированные каналы
- **Отправляющая** горутина блокируется, пока принимающая не возьмет данные
Запуск:
```bash
go run unbuffered.go 
```
## Буферизированные каналы
- **Принимающая** горутина блокируется, если канал пуст
Запуск:
```bash
go run buffered.go 
```

Преимущества буферизованных каналов:
 - отделяют поставщиков и потребителей друг от друга;
 - помогают при всплесках трафика;
 - улучшают производительность;
 - отлично подходят для пула воркеров;
 - хороши в асинхронных сигналах или событиях.