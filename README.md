# goApp
Консольна програма як шукає вільний час в Google Calendar користувача та пропонує фільми, які користувач може відвідати.
# Завантаження і запуск проекту
1. Встановити VScode https://code.visualstudio.com
2. Встановити golang. В VScode достатньо встановити розширення, посібник - https://code.visualstudio.com/docs/languages/go
3. Завантажте архів з файлами, розархівуйте в довільну папку та відкрийте її в VScode
# Запуск, збірка та робота з програмою
Щоб запустити програму з консолі введіть "go run .", щоб зібрати в виконуваний файл введіть "go build".
При першому запуску потрібно авторизуватися. Спочатку програма попросить ім'я користувача, воно буде пов'язана з вашою аворизацією, якщо при наступних запусках вводити одне і теж ім'я авторизація буде не потрібна.

Ввевши ім'я ми отримуємо посилання на авторизації google(поки вона доступна для акаунтів, які є тестувальниками). Після успішної авторизації залишиться вкладка з відповідю від google де буде код авторизації, він потрібен програмі для отримання доступу до Вашого календаря.

Сторінка може виглядати так:
![image](https://github.com/TroFa0/goApp/assets/117114700/decef5ed-1621-4210-b30a-9c461e759e59)
В посиланні знаходиться код, між code= та &
![45643](https://github.com/TroFa0/goApp/assets/117114700/7b33bac2-314d-4748-813d-ac26aae43d95)
Скопіюйте цей код та вставте в консоль, якщо все вірно то програма почне видавати фільми які ви можете відвідати.
