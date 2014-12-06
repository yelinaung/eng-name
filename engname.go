package engname

import (
	"math/rand"
	"time"
)

func int() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func GetMenName() string {
	GN_MEN := []string{"Adam", "Andrewe", "Andro", "Anthonye", "Anthony", "Anthonie", "Archbould", "Arthure", "Ambrose", "Alexander", "Barthe", "Bartram", "Bryan", "Christofer", "Christopher", "Christofor", "Cuthbert", "Cuthberte", "Charles", "Deane", "Den", "Edwarde", "Edward", "Edmunde", "Edmund", "Euyn", "Francis", "Gabriell", "Gawen", "George", "Georg", "Gerard", "Germayne", "Gilbert", "Gylbert", "Gillian", "Godfrey", "Gyeles", "Harye", "Hector", "Henry", "Henrye", "Helyng", "Hewgh", "Heugh", "Hewghe", "Hughe", "Hewe", "Jamys", "James", "Jaymes", "Jarard", "Jarrard", "Jerrerd", "Jarret", "Jaspar", "John", "Johne", "Johan", "Jhon", "Jasper", "Lancelot", "Lanclott", "Lanslet", "Lanslett", "Lancelote", "Lamwell", "Leonard", "Lanard", "Lewis", "Laurence", "Lawrence", "Lyones", "Lyonell", "Mathew", "Matthew", "Mathewe", "Marke", "Mark", "Martayn", "Martine", "Martyn", "Martyne", "Marmaduke", "Marmeduke", "Michell", "Mychaell", "Mychell", "Myghall", "Mydfurthe", "Nycoles", "Nicholas", "Nycholas", "Odnall", "Oswyne", "Oswold", "Oswolde", "Oswald", "Uswolde", "Patryke", "Percyvall", "Percivell", "Peires", "Peter", "Randell", "Rundall", "Ralph", "Rauffe", "Rawfe", "Raff", "Rauf", "Rawff", "Rauff", "Ralffe", "Raphe", "Ranolde", "Reginolde", "Reignolde", "Robert", "Robarte", "Richard", "Rycherde", "Richerd", "Rychard", "Rycherd", "Richarde", "Robert", "Roberta", "Roger", "Rowland", "Rouland", "Rynyone", "Sampson", "Stevne", "Stephen", "Symond", "Thomas", "Tristram", "Tymothie", "Umfray", "Umphrey", "Umfraie", "Umfraye", "Humphrey", "Walter", "William", "Wyllm"}

	return GN_MEN[rand.Intn(len(GN_MEN))] + getSName()
}

func GetWomenName() string {
	GN_WOMEN := []string{"Agnesse", "Alice", "Alyce", "Allis", "Allice", "Alles", "Alison", "Alleson", "Allysone", "Auinson", "Anne", "Ann", "An", "Annesse", "Annes", "Annas", "Auesonne", "Barbara", "Barbare", "Barberye", "Beil", "Beyll", "Besse", "Bessie", "Brydgitt", "Cecilie", "Cecill", "Cicell", "Cicilie", "Sissilie", "Sisseley", "Christien", "Christobell", "Cristye", "Dorothy", "Dorothye", "Dorithy", "Dorthe", "Dorathye", "Doritie", "Dyna", "Ellen", "Elynor", "Elyoner", "Elioner", "Ellinor", "Elyner", "Elianer", "Elloner", "Elsabith", "Esabeth", "Elizabeth", "Elisabethe", "Elsabythe", "Elyzabeth", "Elisabeth", "Elsapeth", "Fayth", "Florence", "Fortune", "Forton", "Frances", "Francis", "Grace", "Helen", "Hellenor", "Isabell", "Esabell", "Elsebell", "Essabell", "Esable", "Isbell", "Isabelle", "Yssabell", "Isabella", "Jane", "Jan", "Jayn", "Jhayne", "Jaine", "Jayne", "Jayn", "Jeane", "Jean", "Janet", "Janett", "Jennet", "Jennat", "Genet", "Jenate", "Jenatt", "Johannet", "Jellyan", "Julian", "Joan", "Juliane", "Katheryn", "Katherine", "", "Kathirine", "Katheron", "Catherine", "Kusteris", "Luce", "Mabell", "Margery", "Margerye", "Margerre", "Margerie", "Margaret", "Margat", "Margrate", "Margarett", "Margret", "Margreat", "Marie", "Marye", "Mary", "Marion", "Maryan", "Maryon", "Martha", "Mauglen", "Mawld", "Mawde", "Meriall", "Phillis", "Phillis", "Rebecca", "Rebacca", "Suzane", "Susana", "Thomazin", "Thomasine", "Thomazin", "Ursuly", "Ursula"}
	return GN_WOMEN[rand.Intn(len(GN_WOMEN))] + getSName()
}

func getSName() string {
	SNAMES := []string{"Ambrose", "Arrossmythe", "Armestrong", "Armestronge", "Armyn", "Armorer", "Atkinsone", "Atkinson", "Allandson", "Allen", "Anderson", "Appelbie", "Aske", "Ayre", "Ayton", "Bailey", "Baills", "Bedyke", "Baites", "Bateman", "Baylls", "baills", "Bayker", "Baxter", "Burdon", "Barker", "Bartram", "Barton", "Baxter", "Bewyke", "Bewicke", "Benson", "Beckham", "Bednell", "Braccenbe", "Bainbrigge", "Bainebrigge", "Bainbrigg", "Bell", "Bellasses", "Bettlestonne", "Benett", "Bewicke", "Basnett", "Bilton", "Burrell", "Burlessone", "Blackman", "Blenkenshope", "Blythman", "Birkehead", "Blaikden", "Blakiston", "Blaxton", "Blaixton", "Blunt", "Blownt", "Blount", "Bois", "Bone", "Boothe", "Booker", "Bowdon", "Bowswell", "Boutflower", "Bourn", "Bowden", "Bramall", "Brasse", "Brakenby", "Bradeley", "Brewester", "Bryane", "Brigham", "Browne", "Brandling", "Brackenbury", "Buck", "Burlessone", "Burden", "Burrell", "Butter", "Byerley", "Byerlaye", "Cay", "Cassope", "Carnabbe", "Catrick", "Catriok", "Carter", "Chator", "Chaitor", "Charlton", "Colynwood", "Carr", "Carre", "Camber", "Chamber", "Carnabye", "Carnabie", "Chamber", "Chappman", "Chapman", "Chilton", "Clarke", "Claverne", "Clerk", "Cleoburne", "Cleoburne", "Clewghe", "Cocke", "Cloughe", "Coockeson", "Cockson", "Collyson", "Cornefourthe", "Constable", "Coteeforthe", "Collingwood", "Coltard", "Coltman", "Conyares", "Conyers", "Coperwhate", "Cowpelande", "Cramer", "Cralington", "Cramlington", "Cramlyngton", "Crane", "Crayne", "Craister", "Craster", "Cramer", "Crawforthe", "Craycall", "Crage", "Craster", "Collingwood", "Cook", "Cooke", "Conyers", "Cottesworthe", "Claxton", "Crathorne", "Chaitor", "Chator", "Daltone", "Davynson", "Davison", "Davey", "Dearham", "Delavell", "Delaval", "Daniell", "Dente", "Dent", "Dennyson", "Dickinson", "Dixkinson", "Dykkinson", "Digglee", "Dinsdall", "Dacre", "Dalton", "Dyxon", "Dixon", "Dawney", "Duckett", "Dobson", "Donne", "Dunne", "Dune", "Dossy", "Dowghtwhett", "Doughtwhett", "Eirkehouse", "Elyson", "Ellison", "Ellikar", "Elstoppe", "Elstobbe", "Elwood", "Emerson", "Eryngton", "Erington", "Eringtonn", "Ewrie", "Eden", "Edgar", "Egleston", "Ewbank", "Fairbanks", "Farnaby", "Fawell", "Fetherston", "Fetherstonhaugh", "Fennyk", "Fennyke", "Fenyk", "Fenyke", "Fenwik", "Fenwyk", "Fenwicke", "Ferrye", "Fewler", "Fintche", "Fishe", "Flower", "Forster", "Franckleyne", "Francklyn", "Francis", "Franche", "Frissell", "Fuister", "Furrowe", "Fyeffe", "Gack", "Gallon", "Garnett", "Garthe", "Gardiner", "Geilson", "Gibson", "Gipson", "Girlington", "Golle", "Glover", "Gooderyche", "Grene", "Grenewell", "Gregsou", "Gregge", "Gyrdler", "Gucyne", "Gray", "Grayed", "Hacforth", "Haisteye", "Hagthorpe", "Haggthroppe", "Haisteye", "Halliday", "Halliman", "Harbotele", "Harbotell", "Harding", "Hasslewoode", "Harwoode", "Hawton", "Hebborne", "Heath", "Harrigaite", "Harrogait", "Herrisone", "Heryson", "Heighingtone", "Helcott", "Helcot", "Heron", "Harteborne", "Hardynge", "Hall", "Hallowell", "Hebborne", "Hedley", "Heghington", "Heth", "Hicson", "Hilton", "Hindmas", "Harle", "Hodgson", "Hodgeson", "Hodschon", "Hodshon", "Hogerde", "Hondley", "Hopper", "Hoppine", "Horsley", "Horssley", "Houghell", "Howlesworth", "Hewton", "Hevisydes", "Hicson", "Hiltone", "Hudspethe", "Hutcheson", "Hutchinson", "Hubberstee", "Hudspethe", "Hudspeth", "Hunter", "Hutton", "Hutone", "Hyghley", "Ironside", "Jeffrason", "Jenison", "Jennyngson", "Johnson", "Jollye", "Katherick", "Kaye", "Kaylame", "Keathe", "Keye", "Kirkbye", "Kirkehouse", "Kirton", "Killinghall", "Kiplinge", "Kiplin", "Kinge", "Lackenbye", "Lampton", "Lancaster", "Lasse", "Lawes", "Lawson", "Laxe", "Leddell", "Liddle", "Lee", "Lighten", "Lighlye", "Lockye", "Lowghanbye", "Ludge", "Lygthone", "Lewen", "Macame", "Maddeson", "Maire", "Marley", "Mallorye", "Marshall", "Marwodd", "Maxsone", "Maughan", "Maugham", "Medcalf", "Mytford", "Metfford", "Mitford", "Mytfourthe", "Mettfourthe", "Midfourth", "Mydforthe", "Mydfurthe", "Mill", "Milbourne", "Mylburne", "Miller", "Moberley", "Morgaine", "Morland", "Morpeth", "Morley", "Morton", "Mowl", "Mowell", "Moyser", "Muschaunce", "Mylborne", "Myllot", "Midelton", "Middleton", "Newbye", "Newton", "Nichell", "Norman", "Nycholson", "Nicholson", "Ogell", "Ogle", "Oliver", "Olyver", "Ollever", "Orde", "Orpwood", "Pacoke", "Paige", "Palmes", "Pape", "Pasinor", "Paker", "Parkinn", "Parkinson", "Parkingson", "Patteson", "Payckock", "Pawlin", "Pele", "Percye", "Pickeringe", "Pine", "Porter", "Potter", "Pressick", "Preston", "Prestwigg", "Prierman", "Philloppe", "Rand", "Rande", "Ratliffe", "Ratclif", "Rackett", "Ranoldsonne", "Ray", "Rampshawe", "Raw", "Rawe", "Rippes", "Rogerlie", "Rolandsone", "Rounseforth", "Rowthe", "Rede", "Redheade", "Reivelye", "Rey", "Ruddesforth", "Rukebye", "Rukeby", "Russell", "Rutless", "Rutlodge", "Robynson", "Robinsonne", "Riddell", "Rypley", "Sands", "Sanderson", "Salkeld", "Salving", "Sayr", "Saier", "Sawer", "Scott", "Selby", "Selbie", "Salkeld", "Semer", "Shaiftoo", "Shaftoe", "Sharpe", "Shawdfurthe", "Sandefurthe", "Shaldeforthe", "Sherewood", "Sheill", "Shell", "Sigeswike", "Sighwicke", "Sigwick", "Sigswicke", "Sigswick", "Silvertop", "Singleton", "Skaythlocke", "Smathwhaite", "Smythe", "Smith", "Smorquet", "Softly", "Sotheran", "Spence", "Spencer", "Stampfurthe", "Stappltone", "Stele", "Stellin", "Stevensone", "Stewart", "Strangwishe", "Strangwyche", "Still", "Stokar", "Stones", "Storey", "Strother", "Stampe", "Stamp", "Surteis", "Surttes", "Surttees", "Surtes", "Swan", "Swann", "Swynborne", "Swineborne", "Teasdaile", "Teasdell", "Tynmothe", "Tobie", "Todd", "Tood", "Tod", "Tompson", "Thompson", "Toller", "Toplyffe", "Trotter", "Troullope", "Troloppe", "Townes", "Turner", "Turping", "Taylor", "Tailyer", "Tayler", "Teisdaill", "Tempest", "Teddcastell", "Thewe", "Thursbie", "Thursby", "Throckmorton", "Ubancke", "Wales", "Wall", "Walgrauve", "Wandisforde", "Wardayle", "Warde", "Watter", "Weddefeld", "Welfoott", "Whelpden", "Wheatley", "White", "Whitfeild", "Whytfelde", "Whitfelde", "Whitfeld", "Wilson", "Wild", "Williamson", "Wivell", "Woldhave", "Woodfall", "Woller", "Wordye", "Wright", "Wylkeson", "Wilkinson", "Welkenson", "Whetston", "Wedeston", "Wharton", "Whorton", "Watson", "Wythrington", "Wuddrington", "Witheringtone", "Wederington", "Welche", "Whitfeild", "Wyckylffe", "Wylley", "Wympraye", "Wimfrey", "Wyvell", "Wivell", "Vacye", "Volenbye", "Yong"}

	return SNAMES[rand.Intn(len(SNAMES))]
}
