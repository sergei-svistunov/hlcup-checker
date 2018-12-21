# Unofficial Highload Cup (https://highloadcup.ru) checker

## Install
```bash
 go get -u github.com/sergei-svistunov/hlcup-checker
 go install github.com/sergei-svistunov/hlcup-checker
```

## Usage
#### Help
```bash
 `go env GOPATH`/bin/hlcup-checker --help
 ```
 #### Example
 ```bash
 `go env GOPATH`/bin/hlcup-checker --data-dir /tmp/data
```
```
2018/12/21 14:33:14 [1] /accounts/filter/?query_id=840&country_size=%D0%A0%D1%83%D0%BC%D0%B8%D0%B7%D0%B8%D1%8F&limit=16&phone_code=961
2018/12/21 14:33:14 [2] /accounts/hoitlilembunta/cothaannaftefain/?query_id=360
2018/12/21 14:33:14 [3] /accounts/group/?order=1&keys=status&query_id=120&birth=2005&limit=50&country=%D0%A0%D0%BE%D1%81%D1%86%D0%B8%D1%8F
2018/12/21 14:33:14 [4] /accounts/group/?likes=8076&order=1&keys=interests&limit=30&query_id=240
2018/12/21 14:33:14 [5] /accounts/filter/?email_gt=ha&limit=6&query_id=480&city_eq=%D0%9D%D0%BE%D0%B2%D0%BE%D0%BA%D1%8F%D1%80%D1%81%D0%BA
2018/12/21 14:33:14 [6] /accounts/1022/suggest/?limit=goseut&query_id=720
2018/12/21 14:33:14 [7] /accounts/filter/?status_neq=%D0%B2%D1%81%D1%91+%D1%81%D0%BB%D0%BE%D0%B6%D0%BD%D0%BE&query_id=600&limit=38&country_eq=%D0%93%D0%B5%D1%80%D0%BB%D1%8F%D0%BD%D0%B4%D0%B8%D1%8F&interests_any=%D0%A1%D0%B0%D0%BB%D0%B0%D1%82%D1%8B%2CFacebook
2018/12/21 14:33:14 [8] /accounts/4846/suggest/?query_id=0&limit=18&country=%D0%98%D1%81%D0%BF%D1%86%D0%B8%D1%8F
2018/12/21 14:33:14 	Exptected: map[accounts:[map[email:fuecesnavomtopseren@ya.ru fname:Олег id:9749 sname:Феташелан status:всё сложно] map[email:itenwetterdedowceeb@ya.ru fname:Сидор id:9737 sname:Фететатин status:всё сложно] map[email:venotwarmotahatmepfum@yandex.ru fname:Антон id:9673 sname:Фаатотев status:свободны] map[email:alhewberyf@mail.ru fname:Святослав id:9511 sname:Клерыкало status:всё сложно] map[email:dunecur@email.com fname:Кирилл id:9367 sname:Клерушуный status:свободны] map[email:anhiireh@inbox.ru fname:Артём id:9039 sname:Фаыкакий status:свободны] map[email:loninelo@rambler.ru fname:Владимир id:8857 sname:Фетленвен status:свободны] map[email:efpikardehrehanil@inbox.ru fname:Степан id:8801 sname:Фаушуный status:всё сложно] map[email:resyten@icloud.com fname:Игорь id:8793 sname:Пенатоло status:свободны] map[email:opgeniroodde@icloud.com fname:Александр id:8475 sname:Пенолович status:свободны] map[email:talanteehalhacotfi@inbox.ru fname:Виталий id:8427 sname:Стаметаный status:свободны] map[email:meottelitsuackini@icloud.com fname:Мирослав id:8121 status:всё сложно] map[email:nelwegsa@me.com fname:Никита id:7927 sname:Кисушучан status:заняты] map[email:caraumwulvenosdo@ymail.com fname:Дамир id:7783 sname:Клерушуко status:свободны] map[email:owuclimpa@ymail.com fname:Борислав id:7769 sname:Лукушукий status:всё сложно] map[email:patatisoco@yahoo.com fname:Святослав id:7563 sname:Фаолоный status:заняты] map[email:otatwe@inbox.ru fname:Никита id:7379 sname:Стаматоный status:заняты] map[email:ynrohupno@ya.ru fname:Фёдор id:7331 sname:Феташевич status:свободны]]]
2018/12/21 14:33:14 	      Got: map[accounts:[map[email:tygbegithimi@email.com fname:Егор id:9995 sname:Данушутин status:всё сложно] map[email:hocheap@ya.ru fname:Борис id:9865 sname:Колушучан status:заняты] map[email:hofehahdohertusedtec@mail.ru id:9623 status:заняты] map[email:dunecur@email.com fname:Кирилл id:9367 sname:Клерушуный status:свободны] map[email:safiltenci@gmail.com fname:Георгий id:9361 sname:Лукетаный status:свободны] map[email:toorwiwafpufo@yandex.ru fname:Андрей id:9145 sname:Пенолотев status:всё сложно] map[email:hiwadre@yandex.ru fname:Дмитрий id:9095 sname:Хопатотин status:заняты] map[email:toitsyrdelnahatanma@ya.ru fname:Аркадий id:8805 sname:Хопушувен status:заняты] map[email:domirad@list.ru fname:Виталий id:8757 sname:Лебушутин status:заняты] map[email:itpeisnefkedocet@icloud.com fname:Михаил id:8637 sname:Лукатолан status:свободны] map[email:featnehsersasar@inbox.com fname:Сидор id:8503 sname:Данашеный status:всё сложно] map[email:loneryoh@yahoo.com id:8195 status:свободны] map[email:oduhnenno@icloud.com fname:Вячеслав id:7971 sname:Кисололан status:свободны] map[email:sawotoludfelat@mail.ru fname:Святослав id:7763 status:заняты] map[email:cathonsi@email.com fname:Владислав id:7695 sname:Хопололо status:свободны] map[email:gaeded@inbox.com fname:Григорий id:7627 status:свободны] map[email:esfuretdeetolorben@list.ru fname:Данила id:7621 sname:Тератоко status:свободны] map[email:haxhovafien@yandex.ru fname:Владимир id:7615 sname:Фалентев status:свободны]]]
2018/12/21 14:33:14 	      Diff:
--- Expected
+++ Got
@@ -3,6 +3,6 @@
   (map[string]interface {}) (len=5) {
-   (string) (len=5) "email": (string) (len=25) "fuecesnavomtopseren@ya.ru",
-   (string) (len=5) "fname": (string) (len=8) "Олег",
-   (string) (len=2) "id": (float64) 9749,
-   (string) (len=5) "sname": (string) (len=18) "Феташелан",
+   (string) (len=5) "email": (string) (len=22) "tygbegithimi@email.com",
+   (string) (len=5) "fname": (string) (len=8) "Егор",
+   (string) (len=2) "id": (float64) 9995,
+   (string) (len=5) "sname": (string) (len=18) "Данушутин",
    (string) (len=6) "status": (string) (len=19) "всё сложно"
@@ -10,21 +10,12 @@
   (map[string]interface {}) (len=5) {
-   (string) (len=5) "email": (string) (len=25) "itenwetterdedowceeb@ya.ru",
-   (string) (len=5) "fname": (string) (len=10) "Сидор",
-   (string) (len=2) "id": (float64) 9737,
-   (string) (len=5) "sname": (string) (len=18) "Фететатин",
-   (string) (len=6) "status": (string) (len=19) "всё сложно"
+   (string) (len=5) "email": (string) (len=13) "hocheap@ya.ru",
+   (string) (len=5) "fname": (string) (len=10) "Борис",
+   (string) (len=2) "id": (float64) 9865,
+   (string) (len=5) "sname": (string) (len=18) "Колушучан",
+   (string) (len=6) "status": (string) (len=12) "заняты"
   },
-  (map[string]interface {}) (len=5) {
-   (string) (len=5) "email": (string) (len=31) "venotwarmotahatmepfum@yandex.ru",
-   (string) (len=5) "fname": (string) (len=10) "Антон",
-   (string) (len=2) "id": (float64) 9673,
-   (string) (len=5) "sname": (string) (len=16) "Фаатотев",
-   (string) (len=6) "status": (string) (len=16) "свободны"
-  },
-  (map[string]interface {}) (len=5) {
-   (string) (len=5) "email": (string) (len=18) "alhewberyf@mail.ru",
-   (string) (len=5) "fname": (string) (len=18) "Святослав",
-   (string) (len=2) "id": (float64) 9511,
-   (string) (len=5) "sname": (string) (len=18) "Клерыкало",
-   (string) (len=6) "status": (string) (len=19) "всё сложно"
+  (map[string]interface {}) (len=3) {
+   (string) (len=5) "email": (string) (len=28) "hofehahdohertusedtec@mail.ru",
+   (string) (len=2) "id": (float64) 9623,
+   (string) (len=6) "status": (string) (len=12) "заняты"
   },
@@ -38,6 +29,6 @@
   (map[string]interface {}) (len=5) {
-   (string) (len=5) "email": (string) (len=17) "anhiireh@inbox.ru",
-   (string) (len=5) "fname": (string) (len=10) "Артём",
-   (string) (len=2) "id": (float64) 9039,
-   (string) (len=5) "sname": (string) (len=16) "Фаыкакий",
+   (string) (len=5) "email": (string) (len=20) "safiltenci@gmail.com",
+   (string) (len=5) "fname": (string) (len=14) "Георгий",
+   (string) (len=2) "id": (float64) 9361,
+   (string) (len=5) "sname": (string) (len=18) "Лукетаный",
    (string) (len=6) "status": (string) (len=16) "свободны"
......
```