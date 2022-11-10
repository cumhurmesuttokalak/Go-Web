-net/http package kullanıyorum.


//Go da tanımladıgım degişkeni html dosyasına ekleme
-icine http.Response writer ve *http.Request değişkeni alan bir fonksiyoun oluşturuyorum sonra static olarak html bir dosyayı template.ParseFiles funcı ile cağırıyorum.Sonra geri dönen değeri execute ediyorum writerla.fonksiyon icindeki değişkenleri kullanabilmem için exucute metodunun ikinci parametresi ile göndererip html sayfamda ise {{ . }} ile cağırabilirim.Aynı şekilde ikinci parametreye structda verebiliriz.


//templates kullanım 
-template işlemleri için https://pkg.go.dev/text/template 


//master blade yapısı
-birden fazla html dosyaları ile işlem yapmak için html kodlarımı {{define <degisken_ismi>}} {{end}} bloklarımın içine alıyorum.Kısaca html dosyamı bir değişkene atıyorum.main.go tarafında da parse edeceğim html dosyalarını parse ediyorum.Birden fazla html dosyası ile işlem yapacagım icin tmp.Execute degil de tmp.ExecuteTemplate metodunu kullanıp html e vermiş oldugum değişkeni giriyorum.Ve bu şekilde artık bir html dosyamda tanımladığım değişkeni diğer html dosyamdan da erişebilirim.Cagırırken ise {{template "<diger_html_dosya_degisken_tanımı>"}} şeklinde yapıyorum.Kısaca master blade yapısı bu şekilde yapılıyor go da.

//routing işlemi
-go get  github.com/julienschmidt/httprouter packagemi dahil ediyorum.
-main.go dosyamda bir htpprouter.New() metodu ile bir değişken tanımlayıp routing işlemlerimi bu değişken ile yapacagım.
-bir get isteği gönderelim r.GET("/router",Router)
-Router funcımız func Router(w http.ResponseWriter,r *http.Request,params httprouter.Params) şeklinde olmalı.
-html dosyası ile ilişkenleştiriyorum.
-En son da http.ListenAndServe metodumda nil olan yere httprouter.New() değişkenimi veriyorum.
-Eğer gelen parametrelerimi(mesela path bilgisi olabilir,id değişkeni olabilir) yakalamak istersem Router funcımda params.ByName("id") ile yakalayıp tmp.execute metodumda datamı veriyorum.

//html dosyası formlar ile işlemler(GET,POST...)
-html dosyası oluşturup bir form oluşturuyorum.formun actionu /deneme olacak ve mothodum get ya da post fark etmez.Form icinde bir input ve submit tanımlıyorum inputumun name ise username olacak.
-sonra main.go dosyamda bir httpfuncı oluşturuyorum.Ben bir istekde bulunacagım icin benim funcımdaki (r *http.Request) bu işlemleri yapmamı saglıyor.r.FormValue metodunu kullanıp formdaki değişken ismini veriyorum.ve bu metodu bir degişkene atayıp ekrana bastırıyorum.Bu şekilde input girildiği zaman bana girdiyi dönüyor.

//dosya yükleme işlemi
-ilk olarak formumun bir enctype=multipart/form-data ekliyorum.Ve bir tane input ve button kooyuyorum.Form da yapacağım işlemler bu kadar.
-main.go da deneme fonksiyonum icinde r.Parsemultipartform(10) yapıp dosya büyüklüğünü gb şeklinde belirtiyorum.
-sonra r.FormFile() metodu ile inputdaki nameyi veriyorum.Ve geri dönen değerleri tek tek alıyorum.Artık file benim elimde.
-os.OpenFile("upload",os.O_WRONLY|os.O_CREATE,0666)ile yükleyeceğim dosyayı seçiyorum.Dosya yoksa oluşturuyorum.geri dönen değerlerimi belirtiyorum.
-en son io.Copy()ile yükledigim openFile değişkene elimde olan file ımı copyliyorum.


//mysql ile veritabanı bağlantısı
https://gorm.io/docs/connecting_to_the_database.html#MySQL

//mysql ile veritabanı bağlantısı
https://gorm.io/docs/index.html




//Projeye 
-Go Modülü oluşturdum.Gerekli packageleri yükledim.Bir tane resources dosyası oluşturup indirdigim admin template dosyasını içine atıyorum.Admin tarafı için gerekli dosyaları alacagım icin bir tane admin dosyası oluşturuyorum.Sırasıyla admin klasörü altında dosyalarımı oluşturuyorum(assets,models,views,controllers,helpers,)Admin dosyası içinde oluşturdugum assets dosyası içine resources klasörü içine atmış oldugum (css,js,vendor,img,scss)dosyalarını atıyorum.Views klasörü içine bir index sayfası tanımlayıp resource klasöründe bulunan blank.html dosyasının içeriğini kopyalayıp link,css ve js kodlarımın pathlerini düzeltiyorum.Controllers icinde oluşturdugum Dashboard.go controleri dosyası oluşturup gerekli kontrolleri yapıyorum.Dashboard fonksiyonumda ParseFiles yapacagım html dosyalarının yollarını uzun vermemek için helpers klasörü altında bu pathlerimi yakalayan bir fonksiyon yazıp path yerine o fonksiyonumu cagırıyorum.
-Config dosyası oluşturup routes.go klasörü içinde rotalarımı ayarlıyorum.
-main.go dosyamda serveri dinleyorum(port bilgisi ve router bilgisi ile).
-admin template parçalama işlemlerimi yapıyorum.html dosyalarımın hepsine define ile değişkene atayıp index sayfamda template değişken ismi ile çağırıyorum.









-----Gerekli Packageler ------------
go get  github.com/julienschmidt/httprouter
go get -u gorm.io/gorm
gorm.io/driver/mysql