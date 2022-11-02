CREATE DATABASE IF NOT EXISTS test;
use test;

CREATE TABLE users (
    id INT(255) auto_increment,
    mail VARCHAR(255) UNIQUE,
    name VARCHAR(255),
    password VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert initial data to users table
INSERT INTO users(mail, name, password, updated_at, created_at, deleted_at)
  VALUES("test@gmail.com", "testname","password", NUll, "2020-08-01", NULL);

CREATE TABLE travels (
    id INT(255) auto_increment NOT NULL,
    user_id INT(255) NOT NULL,
    destination_id INT(255) NOT NULL,
    departure_id INT(255) NOT NULL,
    max_person INT(8) NOT NULL,
    max_day INT(8) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE prefectures (
    id INT(3) auto_increment NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert initial data of prefectures into prefectures table
INSERT INTO prefectures(id, name, code, updated_at, created_at, deleted_at)
  VALUES(1, "北海道","hokkaido", NULL, "2020-08-01", NULL),
        (2, "青森県", "aomori", NULL, "2020-08-01", NULL),
        (3, "岩手県", "iwate", NULL, "2020-08-01", NULL),
        (4, "宮城県", "miyagi", NULL, "2020-08-01", NULL),
        (5, "秋田県", "akita", NULL, "2020-08-01", NULL),
        (6, "山形県", "yamagata", NULL, "2020-08-01", NULL),
        (7, "福島県", "hukushima", NULL, "2020-08-01", NULL),
        (8, "茨城県", "ibaragi", NULL, "2020-08-01", NULL),
        (9, "栃木県", "tochigi", NULL, "2020-08-01", NULL),
        (10, "群馬県", "gunma", NULL, "2020-08-01", NULL),
        (11, "埼玉県", "saitama", NULL, "2020-08-01", NULL),
        (12, "千葉県", "tiba", NULL, "2020-08-01", NULL),
        (13, "東京都", "tokyo", NULL, "2020-08-01", NULL),
        (14, "神奈川県", "kanagawa", NULL, "2020-08-01", NULL),
        (15, "新潟県", "niigata", NULL, "2020-08-01", NULL),
        (16, "富山県", "toyama", NULL, "2020-08-01", NULL),
        (17, "石川県", "ishikawa", NULL, "2020-08-01", NULL),
        (18, "福井県", "hukui", NULL, "2020-08-01", NULL),
        (19, "山梨県", "yamanasi", NULL, "2020-08-01", NULL),
        (20, "長野県", "nagano", NULL, "2020-08-01", NULL),
        (21, "岐阜県", "gihu", NULL, "2020-08-01", NULL),
        (22, "静岡県", "shizuoka", NULL, "2020-08-01", NULL),
        (23, "愛知県", "aichi", NULL, "2020-08-01", NULL),
        (24, "三重県", "mie", NULL, "2020-08-01", NULL),
        (25, "滋賀県", "shiga", NULL, "2020-08-01", NULL),
        (26, "京都府", "kyoto", NULL, "2020-08-01", NULL),
        (27, "大阪府", "osaka", NULL, "2020-08-01", NULL),
        (28, "兵庫県", "hyogo", NULL, "2020-08-01", NULL),
        (29, "奈良県", "nara", NULL, "2020-08-01", NULL),
        (30, "和歌山県", "wakayama", NULL, "2020-08-01", NULL),
        (31, "鳥取県", "tottori", NULL, "2020-08-01", NULL),
        (32, "島根県", "simane", NULL, "2020-08-01", NULL),
        (33, "岡山県", "okayama", NULL, "2020-08-01", NULL),
        (34, "広島県", "hiroshima", NULL, "2020-08-01", NULL),
        (35, "山口県", "yamaguchi", NULL, "2020-08-01", NULL),
        (36, "徳島県", "tokushima", NULL, "2020-08-01", NULL),
        (37, "香川県", "kagawa", NULL, "2020-08-01", NULL),
        (38, "愛媛県", "ehime", NULL, "2020-08-01", NULL),
        (39, "高知県", "kouchi", NULL, "2020-08-01", NULL),
        (40, "福岡県", "hukuoka", NULL, "2020-08-01", NULL),
        (41, "佐賀県", "saga", NULL, "2020-08-01", NULL),
        (42, "長崎県", "nagasaki", NULL, "2020-08-01", NULL),
        (43, "熊本県", "kumamoto", NULL, "2020-08-01", NULL),
        (44, "大分県", "ooita", NULL, "2020-08-01", NULL),
        (45, "宮崎県", "miyazaki", NULL, "2020-08-01", NULL),
        (46, "鹿児島県", "kagoshima", NULL, "2020-08-01", NULL),
        (47, "沖縄県", "okinawa", NULL, "2020-08-01", NULL);


CREATE TABLE cities (
    id INT(255) auto_increment NOT NULL,
    prefecture_id INT(3) NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO cities (id, prefecture_id, name, code, updated_at, created_at, deleted_at)
  VALUES(1, 1, "札幌","sapporo", NULL, "2020-08-01", NULL),
        (2, 1, "定山渓","jozankei", NULL, "2020-08-01", NULL),
        (3, 1, "稚内・留萌・利尻・礼文","wakkanai", NULL, "2020-08-01", NULL),
        (4, 1, "網走・紋別・北見・知床","abashiri", NULL, "2020-08-01", NULL),
        (5, 1, "釧路・阿寒・川湯・根室","kushiro", NULL, "2020-08-01", NULL),
        (6, 1, "帯広・十勝","obihiro", NULL, "2020-08-01", NULL),
        (7, 1, "日高・えりも","hidaka", NULL, "2020-08-01", NULL),
        (8, 1, "富良野・美瑛・トマム","furano", NULL, "2020-08-01", NULL),
        (9, 1, "旭川・層雲峡・旭岳","asahikawa", NULL, "2020-08-01", NULL),
        (10, 1, "千歳・支笏・苫小牧・滝川・夕張・空知","chitose", NULL, "2020-08-01", NULL),
        (11, 1, "小樽・キロロ・積丹・余市","otaru", NULL, "2020-08-01", NULL),
        (12, 1, "ルスツ・ニセコ・倶知安","niseko", NULL, "2020-08-01", NULL),
        (13, 1, "函館・湯の川・大沼・奥尻","hakodate", NULL, "2020-08-01", NULL),
        (14, 1, "洞爺・室蘭・登別","noboribetsu", NULL, "2020-08-01", NULL),
        (15, 2, "青森・浅虫温泉","aomori", NULL, "2020-08-01", NULL),
        (16, 2, "津軽半島・五所川原","tsugaru", NULL, "2020-08-01", NULL),
        (17, 2, "白神山地・西津軽","ntsugaru", NULL, "2020-08-01", NULL),
        (18, 2, "弘前・黒石","hirosaki", NULL, "2020-08-01", NULL),
        (19, 2, "八甲田・十和田湖・奥入瀬","towadako", NULL, "2020-08-01", NULL),
        (20, 2, "八戸・三沢・七戸十和田","hachinohe", NULL, "2020-08-01", NULL),
        (21, 2, "下北・大間・むつ","shimokita", NULL, "2020-08-01", NULL),
        (22, 3, "盛岡","morioka", NULL, "2020-08-01", NULL),
        (23, 3, "雫石","shizukuishi", NULL, "2020-08-01", NULL),
        (24, 3, "安比高原・八幡平・二戸","appi", NULL, "2020-08-01", NULL),
        (25, 3, "宮古・久慈・岩泉","kuji", NULL, "2020-08-01", NULL),
        (26, 3, "北上・花巻・遠野","kitakami", NULL, "2020-08-01", NULL),
        (27, 3, "奥州・平泉・一関","ichinoseki", NULL, "2020-08-01", NULL),
        (28, 4, "仙台・多賀城・名取","sendai", NULL, "2020-08-01", NULL),
        (29, 4, "秋保・作並","akiu", NULL, "2020-08-01", NULL),
        (30, 4, "鳴子・古川・くりこま高原","naruko", NULL, "2020-08-01", NULL),
        (31, 4, "松島・塩釜・石巻・南三陸・気仙沼","matsushima", NULL, "2020-08-01", NULL),
        (32, 4, "白石・宮城蔵王","shiroishi", NULL, "2020-08-01", NULL),
        (33, 5, "秋田","akita", NULL, "2020-08-01", NULL),
        (34, 5, "能代・男鹿・白神","noshiro", NULL, "2020-08-01", NULL),
        (35, 5, "大館・鹿角・十和田大湯・八幡平","odate", NULL, "2020-08-01", NULL),
        (36, 5, "角館・大曲・田沢湖","tazawa", NULL, "2020-08-01", NULL),
        (37, 5, "湯沢・横手","yuzawa", NULL, "2020-08-01", NULL),
        (38, 5, "由利本荘・鳥海山","honjo", NULL, "2020-08-01", NULL),
        (39, 6, "山形・蔵王・天童・上山","yamagata", NULL, "2020-08-01", NULL),
        (40, 6, "米沢・赤湯・高畠・長井","yonezawa", NULL, "2020-08-01", NULL),
        (41, 6, "寒河江・月山","sagae", NULL, "2020-08-01", NULL),
        (42, 6, "尾花沢・新庄・村山","mogami", NULL, "2020-08-01", NULL),
        (43, 6, "酒田・鶴岡・湯野浜・温海","shonai", NULL, "2020-08-01", NULL),
        (44, 7, "福島・二本松","fukushima", NULL, "2020-08-01", NULL),
        (45, 7, "会津若松・喜多方","aizu", NULL, "2020-08-01", NULL),
        (46, 7, "猪苗代・表磐梯","bandai", NULL, "2020-08-01", NULL),
        (47, 7, "磐梯高原・裏磐梯","urabandai", NULL, "2020-08-01", NULL),
        (48, 7, "郡山・磐梯熱海","koriyama", NULL, "2020-08-01", NULL),
        (49, 7, "南会津・下郷・只見・檜枝岐","minami", NULL, "2020-08-01", NULL),
        (50, 7, "白河・須賀川","nakadori", NULL, "2020-08-01", NULL),
        (51, 7, "いわき・南相馬・相馬","hamadori", NULL, "2020-08-01", NULL),
        (52, 8, "水戸・笠間","mito", NULL, "2020-08-01", NULL),
        (53, 8, "大洗・ひたちなか","oarai", NULL, "2020-08-01", NULL),
        (54, 8, "日立・北茨城・奥久慈","hitachi", NULL, "2020-08-01", NULL),
        (55, 8, "つくば・土浦・取手","tsukuba", NULL, "2020-08-01", NULL),
        (56, 8, "古河・結城・筑西・常総","yuki", NULL, "2020-08-01", NULL),
        (57, 8, "鹿嶋・神栖・潮来・北浦","kashima", NULL, "2020-08-01", NULL),
        (58, 9, "宇都宮・さくら","utsunomiya", NULL, "2020-08-01", NULL),
        (59, 9, "日光・中禅寺湖・奥日光・今市","nikko", NULL, "2020-08-01", NULL),
        (60, 9, "鬼怒川・川治・湯西川・川俣","kinugawa", NULL, "2020-08-01", NULL),
        (61, 9, "那須・板室・黒磯","nasu", NULL, "2020-08-01", NULL),
        (62, 9, "塩原・矢板・大田原・西那須野","shiobara", NULL, "2020-08-01", NULL),
        (63, 9, "真岡・益子・茂木","mashiko", NULL, "2020-08-01", NULL),
        (64, 9, "小山・足利・佐野・栃木","koyama", NULL, "2020-08-01", NULL),
        (65, 10, "前橋・赤城","maebashi", NULL, "2020-08-01", NULL),
        (66, 10, "伊香保温泉・渋川","ikaho", NULL, "2020-08-01", NULL),
        (67, 10, "万座･嬬恋･北軽井沢","manza", NULL, "2020-08-01", NULL),
        (68, 10, "草津温泉・白根","kusatsu", NULL, "2020-08-01", NULL),
        (69, 10, "四万温泉","shimaonsen", NULL, "2020-08-01", NULL),
        (70, 10, "水上・猿ヶ京・沼田","numata", NULL, "2020-08-01", NULL),
        (71, 10, "尾瀬・丸沼","oze", NULL, "2020-08-01", NULL),
        (72, 10, "伊勢崎・太田・館林・桐生","kiryu", NULL, "2020-08-01", NULL),
        (73, 10, "高崎","takasaki", NULL, "2020-08-01", NULL),
        (74, 10, "富岡・藤岡・安中・磯部温泉","fujioka", NULL, "2020-08-01", NULL),
        (75, 11, "大宮・浦和・川口・上尾","saitama", NULL, "2020-08-01", NULL),
        (76, 11, "草加・越谷・春日部・羽生","kasukabe", NULL, "2020-08-01", NULL),
        (77, 11, "熊谷・深谷・本庄","kumagaya", NULL, "2020-08-01", NULL),
        (78, 11, "川越・東松山・志木・和光","kawagoe", NULL, "2020-08-01", NULL),
        (79, 11, "所沢・狭山・飯能","tokorozawa", NULL, "2020-08-01", NULL),
        (80, 11, "秩父・長瀞","chichibu", NULL, "2020-08-01", NULL),
        (81, 12, "千葉","chiba", NULL, "2020-08-01", NULL),
        (82, 12, "舞浜・浦安・船橋・幕張","keiyo", NULL, "2020-08-01", NULL),
        (83, 12, "松戸・柏","kashiwa", NULL, "2020-08-01", NULL),
        (84, 12, "成田空港・佐倉","narita", NULL, "2020-08-01", NULL),
        (85, 12, "銚子・旭・九十九里・東金・茂原","choshi", NULL, "2020-08-01", NULL),
        (86, 12, "鴨川・勝浦・御宿・養老渓谷","sotobo", NULL, "2020-08-01", NULL),
        (87, 12, "南房総・館山・白浜・千倉","tateyama", NULL, "2020-08-01", NULL),
        (88, 12, "市原・木更津・君津・富津・鋸南","uchibo", NULL, "2020-08-01", NULL),
        (89, 13, "東京２３区内","tokyo", NULL, "2020-08-01", NULL),
        (90, 13, "立川・八王子・町田・府中・吉祥寺","nishi", NULL, "2020-08-01", NULL),
        (91, 13, "奥多摩・青梅・羽村","okutama", NULL, "2020-08-01", NULL),
        (92, 13, "八丈島","ritou", NULL, "2020-08-01", NULL),
        (93, 13, "大島","oshima", NULL, "2020-08-01", NULL),
        (94, 13, "神津島・新島・式根島","kouzu", NULL, "2020-08-01", NULL),
        (95, 13, "三宅島","miyake", NULL, "2020-08-01", NULL),
        (96, 13, "小笠原諸島","Ogasawara", NULL, "2020-08-01", NULL),
        (97, 14, "横浜","yokohama", NULL, "2020-08-01", NULL),
        (98, 14, "川崎","kawasaki", NULL, "2020-08-01", NULL),
        (99, 14, "箱根","hakone", NULL, "2020-08-01", NULL),
        (100, 14, "小田原","odawara", NULL, "2020-08-01", NULL),
        (101, 14, "湯河原・真鶴","yugawara", NULL, "2020-08-01", NULL),
        (102, 14, "相模湖・丹沢","sagamiko", NULL, "2020-08-01", NULL),
        (103, 14, "大和・相模原・町田西部","sagamihara", NULL, "2020-08-01", NULL),
        (104, 14, "厚木・海老名・伊勢原","ebina", NULL, "2020-08-01", NULL),
        (105, 14, "湘南・鎌倉・江ノ島・藤沢・平塚","shonan", NULL, "2020-08-01", NULL),
        (106, 14, "横須賀・三浦","miura", NULL, "2020-08-01", NULL),
        (107, 15, "新潟","niigata", NULL, "2020-08-01", NULL),
        (108, 15, "月岡・瀬波・咲花","kaetsu", NULL, "2020-08-01", NULL),
        (109, 15, "長岡・燕三条・柏崎・弥彦・岩室・寺泊","kita", NULL, "2020-08-01", NULL),
        (110, 15, "魚沼・十日町・津南・六日町・大湯","minami", NULL, "2020-08-01", NULL),
        (111, 15, "越後湯沢・苗場","yuzawa", NULL, "2020-08-01", NULL),
        (112, 15, "上越・糸魚川・妙高","joetsu", NULL, "2020-08-01", NULL),
        (113, 15, "佐渡","sado", NULL, "2020-08-01", NULL),
        (114, 16, "富山・八尾・立山","toyama", NULL, "2020-08-01", NULL),
        (115, 16, "滑川・魚津・朝日・黒部・宇奈月","goto", NULL, "2020-08-01", NULL),
        (116, 16, "高岡・氷見・砺波","gosei", NULL, "2020-08-01", NULL),
        (117, 17, "金沢","kanazawa", NULL, "2020-08-01", NULL),
        (118, 17, "加賀・小松・辰口","kaga", NULL, "2020-08-01", NULL),
        (119, 17, "能登・輪島・珠洲","noto", NULL, "2020-08-01", NULL),
        (120, 17, "七尾・和倉・羽咋","nanao", NULL, "2020-08-01", NULL),
        (121, 18, "福井","hukui", NULL, "2020-08-01", NULL),
        (122, 18, "あわら・三国","awara", NULL, "2020-08-01", NULL),
        (123, 18, "永平寺・勝山・大野","katsuyama", NULL, "2020-08-01", NULL),
        (124, 18, "敦賀・美浜","tsuruga", NULL, "2020-08-01", NULL),
        (125, 18, "若狭・小浜・高浜","obama", NULL, "2020-08-01", NULL),
        (126, 19, "甲府・湯村・昇仙峡","kofu", NULL, "2020-08-01", NULL),
        (127, 19, "山梨・石和・勝沼・塩山","yamanashi", NULL, "2020-08-01", NULL),
        (128, 19, "大月・都留・道志渓谷","otsuki", NULL, "2020-08-01", NULL),
        (129, 19, "山中湖・忍野","yamanakako", NULL, "2020-08-01", NULL),
        (130, 19, "河口湖・富士吉田・本栖湖・西湖・精進湖","kawaguchiko", NULL, "2020-08-01", NULL),
        (131, 19, "下部・身延・早川","minobu", NULL, "2020-08-01", NULL),
        (132, 19, "韮崎・南アルプス","nirasaki", NULL, "2020-08-01", NULL),
        (133, 19, "八ヶ岳・小淵沢・清里・大泉","kiyosato", NULL, "2020-08-01", NULL),
        (134, 20, "長野・小布施・信州高山・戸隠・飯綱","nagano", NULL, "2020-08-01", NULL),
        (135, 20, "斑尾・飯山・信濃町・野尻湖・黒姫","madara", NULL, "2020-08-01", NULL),
        (136, 20, "野沢温泉・木島平・秋山郷","nozawa", NULL, "2020-08-01", NULL),
        (137, 20, "志賀高原･湯田中･渋","shiga", NULL, "2020-08-01", NULL),
        (138, 20, "上田・別所・鹿教湯","ueda", NULL, "2020-08-01", NULL),
        (139, 20, "戸倉上山田・千曲","chikuma", NULL, "2020-08-01", NULL),
        (140, 20, "菅平・峰の原","sugadaira", NULL, "2020-08-01", NULL),
        (141, 20, "軽井沢・佐久･小諸","karui", NULL, "2020-08-01", NULL),
        (142, 20, "八ヶ岳・野辺山・富士見・原村","yatsu", NULL, "2020-08-01", NULL),
        (143, 20, "蓼科・白樺湖・霧ヶ峰・車山","kirigamine", NULL, "2020-08-01", NULL),
        (144, 20, "諏訪湖","suwa", NULL, "2020-08-01", NULL),
        (145, 20, "伊那・駒ヶ根・飯田・昼神","ina", NULL, "2020-08-01", NULL),
        (146, 20, "木曽","kiso", NULL, "2020-08-01", NULL),
        (147, 20, "松本・塩尻・浅間温泉・美ヶ原温泉","matsumo", NULL, "2020-08-01", NULL),
        (148, 20, "上高地・乗鞍・白骨","kamiko", NULL, "2020-08-01", NULL),
        (149, 20, "安曇野・穂高・大町・豊科","hotaka", NULL, "2020-08-01", NULL),
        (150, 20, "白馬・八方尾根・栂池高原・小谷","hakuba", NULL, "2020-08-01", NULL),
        (151, 21, "岐阜・各務原","gifu", NULL, "2020-08-01", NULL),
        (152, 21, "奥飛騨・新穂高","kamitakara", NULL, "2020-08-01", NULL),
        (153, 21, "高山・飛騨","takayama", NULL, "2020-08-01", NULL),
        (154, 21, "下呂温泉・濁河温泉","gero", NULL, "2020-08-01", NULL),
        (155, 21, "中津川・多治見・恵那・美濃加茂","tajimi", NULL, "2020-08-01", NULL),
        (156, 21, "郡上八幡・関・美濃","gujo", NULL, "2020-08-01", NULL),
        (157, 21, "白川郷","shirakawago", NULL, "2020-08-01", NULL),
        (158, 21, "大垣・岐阜羽島","ogaki", NULL, "2020-08-01", NULL),
        (159, 22, "静岡・清水","shizuoka", NULL, "2020-08-01", NULL),
        (160, 22, "熱海","atami", NULL, "2020-08-01", NULL),
        (161, 22, "伊東","ito", NULL, "2020-08-01", NULL),
        (162, 22, "伊豆高原","izukogen", NULL, "2020-08-01", NULL),
        (163, 22, "下田・南伊豆","shimoda", NULL, "2020-08-01", NULL),
        (164, 22, "西伊豆・戸田・土肥・堂ヶ島","nishi", NULL, "2020-08-01", NULL),
        (165, 22, "伊豆長岡・修善寺・天城湯ヶ島","naka", NULL, "2020-08-01", NULL),
        (166, 22, "富士・富士宮","fuji", NULL, "2020-08-01", NULL),
        (167, 22, "御殿場・沼津・三島","numazu", NULL, "2020-08-01", NULL),
        (168, 22, "焼津・藤枝・御前崎・寸又峡","yaizu", NULL, "2020-08-01", NULL),
        (169, 22, "東伊豆・河津","higashi", NULL, "2020-08-01", NULL),
        (170, 22, "浜松・浜名湖・天竜","hamamatsu", NULL, "2020-08-01", NULL),
        (171, 22, "掛川・袋井・磐田","kikugawa", NULL, "2020-08-01", NULL),
        (172, 23, "名古屋","nagoyashi", NULL, "2020-08-01", NULL),
        (173, 23, "豊橋・豊川・蒲郡・伊良湖","mikawawan", NULL, "2020-08-01", NULL),
        (174, 23, "奥三河・新城・湯谷温泉","okumikawa", NULL, "2020-08-01", NULL),
        (175, 23, "豊田・刈谷・知立・安城・岡崎","mikawa", NULL, "2020-08-01", NULL),
        (176, 23, "一宮・犬山・小牧・瀬戸・春日井","owari", NULL, "2020-08-01", NULL),
        (177, 23, "セントレア・東海・半田・知多","chita", NULL, "2020-08-01", NULL),
        (178, 23, "南知多・日間賀島・篠島","minamichita", NULL, "2020-08-01", NULL),
        (179, 24, "津･鈴鹿･亀山","tsu", NULL, "2020-08-01", NULL),
        (180, 24, "四日市・桑名・湯の山・長島温泉","yunoyama", NULL, "2020-08-01", NULL),
        (181, 24, "伊賀・名張","iga", NULL, "2020-08-01", NULL),
        (182, 24, "松阪","matsusaka", NULL, "2020-08-01", NULL),
        (183, 24, "伊勢・二見","ise", NULL, "2020-08-01", NULL),
        (184, 24, "鳥羽","toba", NULL, "2020-08-01", NULL),
        (185, 24, "志摩・賢島","shima", NULL, "2020-08-01", NULL),
        (186, 24, "熊野・尾鷲・紀北","kumano", NULL, "2020-08-01", NULL),
        (187, 25, "大津・雄琴・草津・栗東","ootsu", NULL, "2020-08-01", NULL),
        (188, 25, "湖西・高島・マキノ","kosei", NULL, "2020-08-01", NULL),
        (189, 25, "長浜・米原","kohoku", NULL, "2020-08-01", NULL),
        (190, 25, "彦根・近江八幡・守山・東近江","kotou", NULL, "2020-08-01", NULL),
        (191, 25, "信楽・甲賀","shigaraki", NULL, "2020-08-01", NULL),
        (192, 26, "京都","shi", NULL, "2020-08-01", NULL),
        (193, 26, "宇治・長岡京","nannbu", NULL, "2020-08-01", NULL),
        (194, 26, "亀岡・湯の花・美山・京丹波","yunohana", NULL, "2020-08-01", NULL),
        (195, 26, "福知山・綾部","fukuchiyama", NULL, "2020-08-01", NULL),
        (196, 26, "丹後・久美浜","hokubu", NULL, "2020-08-01", NULL),
        (197, 26, "天橋立・宮津・舞鶴","miyazu", NULL, "2020-08-01", NULL),
        (198, 27, "大阪","shi", NULL, "2020-08-01", NULL),
        (199, 27, "高槻・茨木・箕面・伊丹空港","hokubu", NULL, "2020-08-01", NULL),
        (200, 27, "枚方・守口・東大阪","toubu", NULL, "2020-08-01", NULL),
        (201, 27, "八尾・藤井寺・河内長野","nantou", NULL, "2020-08-01", NULL),
        (202, 27, "堺・岸和田・関空・泉佐野","nanbu", NULL, "2020-08-01", NULL),
        (203, 28, "神戸・有馬温泉・六甲山","kobe", NULL, "2020-08-01", NULL),
        (204, 28, "宝塚・西宮・甲子園・三田・篠山","nantou", NULL, "2020-08-01", NULL),
        (205, 28, "明石・加古川・三木","minamichu", NULL, "2020-08-01", NULL),
        (206, 28, "姫路・相生・赤穂","nannansei", NULL, "2020-08-01", NULL),
        (207, 28, "和田山・竹田城・ハチ高原","chubu", NULL, "2020-08-01", NULL),
        (208, 28, "城崎温泉・豊岡・出石・神鍋","kita", NULL, "2020-08-01", NULL),
        (209, 28, "香住・浜坂・湯村","kasumi", NULL, "2020-08-01", NULL),
        (210, 28, "淡路島","awaji", NULL, "2020-08-01", NULL),
        (211, 29, "奈良・大和高原","nara", NULL, "2020-08-01", NULL),
        (212, 29, "橿原・大和郡山・天理・生駒","hokubu", NULL, "2020-08-01", NULL),
        (213, 29, "吉野・十津川・天川・五條","nanbu", NULL, "2020-08-01", NULL),
        (214, 30, "和歌山・加太・和歌浦","wakayama", NULL, "2020-08-01", NULL),
        (215, 30, "高野山・橋本","Kihoku", NULL, "2020-08-01", NULL),
        (216, 30, "御坊・有田・海南・日高","gobo", NULL, "2020-08-01", NULL),
        (217, 30, "南紀白浜・紀伊田辺・龍神","shirahama", NULL, "2020-08-01", NULL),
        (218, 30, "勝浦・串本・すさみ","Katsuura", NULL, "2020-08-01", NULL),
        (219, 30, "熊野古道・新宮・本宮・中辺路","hongu", NULL, "2020-08-01", NULL),
        (220, 31, "鳥取・岩美・浜村","tottori", NULL, "2020-08-01", NULL),
        (221, 31, "倉吉・三朝温泉","chubu", NULL, "2020-08-01", NULL),
        (222, 31, "米子・皆生温泉・大山","seibu", NULL, "2020-08-01", NULL),
        (223, 32, "松江・玉造・安来・奥出雲","matsue", NULL, "2020-08-01", NULL),
        (224, 32, "出雲・大田・石見銀山","toubu", NULL, "2020-08-01", NULL),
        (225, 32, "津和野・益田・浜田・江津","masuda", NULL, "2020-08-01", NULL),
        (226, 32, "隠岐諸島","ritou", NULL, "2020-08-01", NULL),
        (227, 33, "岡山","okayama", NULL, "2020-08-01", NULL),
        (228, 33, "牛窓・瀬戸内・備前","bizen", NULL, "2020-08-01", NULL),
        (229, 33, "津山・湯郷・美作・奥津","tsuyama", NULL, "2020-08-01", NULL),
        (230, 33, "湯原・蒜山・新見・高梁","niimi", NULL, "2020-08-01", NULL),
        (231, 33, "倉敷・総社・玉野・笠岡","kurashiki", NULL, "2020-08-01", NULL),
        (232, 34, "広島","hiroshima", NULL, "2020-08-01", NULL),
        (233, 34, "東広島・竹原・三原・広島空港","higashihiroshima", NULL, "2020-08-01", NULL),
        (234, 34, "福山・尾道・しまなみ海道","fukuyama", NULL, "2020-08-01", NULL),
        (235, 34, "呉・江田島","kure", NULL, "2020-08-01", NULL),
        (236, 34, "三次・庄原・帝釈峡","shohara", NULL, "2020-08-01", NULL),
        (237, 34, "三段峡・芸北・北広島","sandankyo", NULL, "2020-08-01", NULL),
        (238, 34, "宮島・宮浜温泉・廿日市","miyajima", NULL, "2020-08-01", NULL),
        (239, 35, "山口・湯田温泉・防府","yamaguchi", NULL, "2020-08-01", NULL),
        (240, 35, "下関・宇部","shimonoseki", NULL, "2020-08-01", NULL),
        (241, 35, "岩国・周南・柳井","iwakuni", NULL, "2020-08-01", NULL),
        (242, 35, "萩・長門・秋吉台","hagi", NULL, "2020-08-01", NULL),
        (243, 36, "徳島・鳴門","tokushima", NULL, "2020-08-01", NULL),
        (244, 36, "大歩危・祖谷・剣山・吉野川","hokubu", NULL, "2020-08-01", NULL),
        (245, 36, "阿南・日和佐・宍喰","nanbu", NULL, "2020-08-01", NULL),
        (246, 37, "高松・さぬき・東かがわ","takamatsu", NULL, "2020-08-01", NULL),
        (247, 37, "坂出・宇多津・丸亀","sakaide", NULL, "2020-08-01", NULL),
        (248, 37, "琴平・観音寺","kotohira", NULL, "2020-08-01", NULL),
        (249, 37, "小豆島・直島","ritou", NULL, "2020-08-01", NULL),
        (250, 38, "松山・道後","chuuyo", NULL, "2020-08-01", NULL),
        (251, 38, "今治・しまなみ海道","touyo", NULL, "2020-08-01", NULL),
        (252, 38, "西条・新居浜・四国中央","saijo", NULL, "2020-08-01", NULL),
        (253, 38, "宇和島・八幡浜","nanyo", NULL, "2020-08-01", NULL),
        (254, 39, "高知・南国・香南・伊野","kouchi", NULL, "2020-08-01", NULL),
        (255, 39, "安芸・室戸","toubu", NULL, "2020-08-01", NULL),
        (256, 39, "足摺・四万十・宿毛・須崎","seibu", NULL, "2020-08-01", NULL),
        (257, 40, "博多・キャナルシティ・海の中道・太宰府・二日市","fukuoka", NULL, "2020-08-01", NULL),
        (258, 40, "天神・中洲・薬院・福岡ドーム・糸島","seibu", NULL, "2020-08-01", NULL),
        (259, 40, "北九州","kitakyusyu", NULL, "2020-08-01", NULL),
        (260, 40, "宗像・宮若・飯塚","chikuzen", NULL, "2020-08-01", NULL),
        (261, 40, "久留米・甘木・原鶴温泉・筑後川温泉","kurume", NULL, "2020-08-01", NULL),
        (262, 40, "北九州空港・苅田・行橋・豊前","buzen", NULL, "2020-08-01", NULL),
        (263, 40, "大牟田・柳川・八女・筑後","chikugo", NULL, "2020-08-01", NULL),
        (264, 41, "佐賀・古湯温泉","saga", NULL, "2020-08-01", NULL),
        (265, 41, "鳥栖","tosu", NULL, "2020-08-01", NULL),
        (266, 41, "嬉野・武雄・伊万里・有田・太良","ureshino", NULL, "2020-08-01", NULL),
        (267, 41, "唐津・呼子","karatsu", NULL, "2020-08-01", NULL),
        (268, 42, "長崎","nagasaki", NULL, "2020-08-01", NULL),
        (269, 42, "雲仙・島原・小浜","unzen", NULL, "2020-08-01", NULL),
        (270, 42, "諫早・大村・長崎空港","airport", NULL, "2020-08-01", NULL),
        (271, 42, "ハウステンボス・佐世保・平戸","sasebo", NULL, "2020-08-01", NULL),
        (272, 42, "五島列島","ritou", NULL, "2020-08-01", NULL),
        (273, 42, "対馬","tsushima", NULL, "2020-08-01", NULL),
        (274, 42, "壱岐島","iki", NULL, "2020-08-01", NULL),
        (275, 43, "熊本","kumamoto", NULL, "2020-08-01", NULL),
        (276, 43, "大津・玉名・山鹿・荒尾・菊池","kikuchi", NULL, "2020-08-01", NULL),
        (277, 43, "阿蘇","aso", NULL, "2020-08-01", NULL),
        (278, 43, "宇土・八代・水俣","yatsushiro", NULL, "2020-08-01", NULL),
        (279, 43, "人吉・球磨","kuma", NULL, "2020-08-01", NULL),
        (280, 43, "天草･本渡","amakusa", NULL, "2020-08-01", NULL),
        (281, 43, "黒川温泉・杖立","kurokawa", NULL, "2020-08-01", NULL),
        (282, 44, "大分","oita", NULL, "2020-08-01", NULL),
        (283, 44, "別府・日出","beppu", NULL, "2020-08-01", NULL),
        (284, 44, "佐伯・臼杵・豊後大野","usuki", NULL, "2020-08-01", NULL),
        (285, 44, "湯布院・湯平","yufuin", NULL, "2020-08-01", NULL),
        (286, 44, "久住・竹田","taketa", NULL, "2020-08-01", NULL),
        (287, 44, "九重・日田・天瀬","hita", NULL, "2020-08-01", NULL),
        (288, 44, "国東・中津・宇佐・耶馬渓","kunisaki", NULL, "2020-08-01", NULL),
        (289, 45, "宮崎","miyazaki", NULL, "2020-08-01", NULL),
        (290, 45, "高千穂・延岡・日向・高鍋","hokubu", NULL, "2020-08-01", NULL),
        (291, 45, "都城・えびの・日南・綾","nanbu", NULL, "2020-08-01", NULL),
        (292, 46, "鹿児島・桜島","kagoshima", NULL, "2020-08-01", NULL),
        (293, 46, "霧島・国分・鹿児島空港","oosumi", NULL, "2020-08-01", NULL),
        (294, 46, "鹿屋・垂水・志布志","kanoya", NULL, "2020-08-01", NULL),
        (295, 46, "川内・出水","hokusatsu", NULL, "2020-08-01", NULL),
        (296, 46, "指宿・枕崎・南さつま","nansatsu", NULL, "2020-08-01", NULL),
        (297, 46, "屋久島","yakushima", NULL, "2020-08-01", NULL),
        (298, 46, "種子島","ritou", NULL, "2020-08-01", NULL),
        (299, 46, "奄美大島･喜界島・徳之島","amami", NULL, "2020-08-01", NULL),
        (300, 46, "沖永良部島・与論島","okinoerabu", NULL, "2020-08-01", NULL),
        (301, 47, "那覇","nahashi", NULL, "2020-08-01", NULL),
        (302, 47, "恩納・名護・本部・今帰仁","hokubu", NULL, "2020-08-01", NULL),
        (303, 47, "宜野湾・北谷・読谷・沖縄市・うるま","chubu", NULL, "2020-08-01", NULL),
        (304, 47, "糸満・豊見城・南城","nanbu", NULL, "2020-08-01", NULL),
        (305, 47, "慶良間・渡嘉敷・座間味・阿嘉","kerama", NULL, "2020-08-01", NULL),
        (306, 47, "久米島","kumejima", NULL, "2020-08-01", NULL),
        (307, 47, "宮古島・伊良部島","Miyako", NULL, "2020-08-01", NULL),
        (308, 47, "石垣・西表・小浜島","ritou", NULL, "2020-08-01", NULL),
        (309, 47, "与那国島","yonaguni", NULL, "2020-08-01", NULL),
        (310, 47, "大東島","daito", NULL, "2020-08-01", NULL);
      

CREATE TABLE detailed_cities (
    id INT(255) auto_increment NOT NULL,
    city_id INT(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    code VARCHAR(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- insert initial data to users table
INSERT INTO detailed_cities(id, city_id, name, code, updated_at, created_at, deleted_at)
  VALUES(1, 1,"札幌・新札幌・琴似", "A", NUll, "2020-08-01", NULL),
        (2, 1,"大通公園・時計台・狸小路", "B", NUll, "2020-08-01", NULL),
        (3, 1,"すすきの・中島公園", "C", NUll, "2020-08-01", NULL),
        (4, 89,"東京駅・銀座・秋葉原・東陽町・葛西", "A", NUll, "2020-08-01", NULL),
        (5, 89,"新橋・汐留・浜松町・お台場", "B", NUll, "2020-08-01", NULL),
        (6, 89,"赤坂・六本木・霞ヶ関・永田町", "C", NUll, "2020-08-01", NULL),
        (7, 89,"渋谷・恵比寿・目黒・二子玉川", "D", NUll, "2020-08-01", NULL),
        (8, 89,"品川・大井町・蒲田・羽田空港", "E", NUll, "2020-08-01", NULL),
        (9, 89,"新宿・中野・荻窪・四谷", "F", NUll, "2020-08-01", NULL),
        (10, 89,"池袋・赤羽・巣鴨・大塚", "G", NUll, "2020-08-01", NULL),
        (11, 89,"東京ドーム・飯田橋・御茶ノ水", "H", NUll, "2020-08-01", NULL),
        (12, 89,"上野・浅草・錦糸町・新小岩・北千住", "I", NUll, "2020-08-01", NULL),
        (13, 172,"名古屋駅・伏見・丸の内", "A", NUll, "2020-08-01", NULL),
        (14, 172,"栄・錦・名古屋城", "B", NUll, "2020-08-01", NULL),
        (15, 172,"金山・熱田・千種・ナゴヤドーム", "C", NUll, "2020-08-01", NULL),
        (16, 192,"京都駅", "D", NUll, "2020-08-01", NULL),
        (17, 192,"嵐山・嵯峨野・太秦・高雄", "A", NUll, "2020-08-01", NULL),
        (18, 192,"河原町・四条烏丸・二条城・御所", "B", NUll, "2020-08-01", NULL),
        (19, 192,"祇園・東山・北白川", "C", NUll, "2020-08-01", NULL),
        (20, 192,"大原・鞍馬・貴船", "E", NUll, "2020-08-01", NULL),
        (21, 198,"大阪駅・梅田・ユニバーサルシティ・尼崎", "B", NUll, "2020-08-01", NULL),
        (22, 198,"なんば・心斎橋・天王寺・阿倍野・長居", "D", NUll, "2020-08-01", NULL),
        (23, 198,"京橋・淀屋橋・本町・ベイエリア・弁天町", "C", NUll, "2020-08-01", NULL),
        (24, 198,"新大阪・江坂", "A", NUll, "2020-08-01", NULL);




CREATE TABLE budgets (
    id INT(255) NOT NULL,
    transportations INT(8) NOT NULL,
    accommodation INT(255) NOT NULL,
    sightseeing INT(255) NOT NULL,
    meal INT(255) NOT NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE visits (
    id INT(255) NOT NULL,
    visit_day DATE,
    address VARCHAR(255),
    charge INT(255),
    info VARCHAR(255),
    name VARCHAR(255),
    visit_type varchar(255),
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE transportations (
    id INT(255) NOT NULL,
    destination VARCHAR(255) NOT NULL,
    departure VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    charge INT(255) NOT NULL,
    order_no INT(8) NOT NULL,
    way_there Boolean not NULL,
    updated_at datetime,
	created_at datetime,
	deleted_at datetime,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

