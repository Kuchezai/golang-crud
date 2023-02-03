# golang-crud



На данный момент в проект реализовано:

    1. CRUD через GET-запросы:

CREATE (INSERT) :

[Добавляем запись](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/c1.png)
[Результат](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/c2.png)

READ (SELECT):

[Вывод всех](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/sa1.png)

UPDATE:

[Обновляем запись](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/u1.png)
[Результат](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/u2.png)

DELETE:

[Удаляем запись](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/d1.png)
[Результат](https://github.com/Kuchezai/golang-crud/raw/master/exmpl/d2.png)

    2. Конфиг в yaml-файле:

internal/app/config.go и configs/config.yaml
    
    3. Вывод SELECT (*) в html-шаблон:

internal\app\controllers\user.go ShowUI()

    4. CREATE через POST-запрос из формы:

internal\app\controllers\user.go CreateFromUi()

    5. Логирование ошибок и информации в .txt файл:

internal\app\logs

    6. Middleware admin:

internal\app\middleware

    7. JSON Web Token:

internal\auth

    8. Верификация пользователя через email:

internal\smtp\EmailSender.go

    9. Миграции БД

internal\app\db\migration

    10. Загрузка изображений на сервер

internal\app\service\ImageUploader.go