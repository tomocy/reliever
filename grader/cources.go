package grader

type Course struct {
	Kind    string
	Name    string
	Credits int
}

var courses = map[string]Course{
	"日本の歴史": {"基礎教養1", "日本の歴史", 2}, "東洋の歴史": {"基礎教養1", "東洋の歴史", 2}, "西洋の歴史": {"基礎教養1", "西洋の歴史", 2},
	"日本の文学": {"基礎教養1", "日本の文学", 2}, "東洋の文学": {"基礎教養1", "東洋の文学", 2}, "西洋の文学": {"基礎教養1", "西洋の文学", 2},
	"教養としての日本語": {"基礎教養1", "教養としての日本語", 2}, "東洋の芸術": {"基礎教養1", "東洋の芸術", 2}, "西洋の芸術": {"基礎教養1", "西洋の芸術", 2},
	"世界の思想": {"基礎教養1", "世界の思想", 2}, "日本国憲法": {"基礎教養1", "日本国憲法", 2}, "法の世界": {"基礎教養1", "法の世界", 2},
	"政治の世界": {"基礎教養1", "政治の世界", 2}, "経済の世界": {"基礎教養1", "経済の世界", 2}, "思考の世界": {"基礎教養1", "思考の世界", 2},
	"人間の世界": {"基礎教養1", "人間の世界", 2}, "心の世界": {"基礎教養1", "心の世界", 2}, "世界の中の日本語日本文化": {"基礎教養1", "世界の中の日本語日本文化", 2},
	"ことばの世界": {"基礎教養1", "ことばの世界", 2}, "Basic Learning Skills 1(Quantitative Research Methods)": {"基礎教養1", "Basic Learning Skills 1(Quantitative Research Methods)", 2}, "Basic Learning Skills 2(Qualitative Research Methods)": {"基礎教養1", "Basic Learning Skills 2(Qualitative Research Methods)", 2},
	"数学の考え方": {"基礎教養2", "数学の考え方", 2}, "生命科学の考え方": {"基礎教養2", "生命科学の考え方", 2}, "化学の考え方": {"基礎教養2", "化学の考え方", 2},
	"物理学の考え方": {"基礎教養2", "物理学の考え方", 2}, "宇宙地球科学の考え方": {"基礎教養2", "宇宙地球科学の考え方", 2}, "グラフィックスの世界": {"基礎教養2", "グラフィックスの世界", 2},
	"健康科学の考え方": {"基礎教養2", "健康科学の考え方", 2}, "文系学生のための科学実験": {"基礎教養2", "文系学生のための科学実験", 2},
	"現代数学の基礎": {"基礎教養3", "現代数学の基礎", 2}, "現代生命科学の基礎": {"基礎教養3", "現代生命科学の基礎", 2}, "現代化学の基礎": {"基礎教養3", "現代化学の基礎", 2},
	"現代物理学の基礎": {"基礎教養3", "現代物理学の基礎", 2}, "現代宇宙地球科学の基礎": {"基礎教養3", "現代宇宙地球科学の基礎", 2}, "グラフィックスの基礎": {"基礎教養3", "グラフィックスの基礎", 2},
	"健康科学の基礎": {"基礎教養3", "健康科学の基礎", 2}, "環境科学の基礎": {"基礎教養3", "環境科学の基礎", 2},
	"現代の法と政治を考える": {"現代教養科目", "現代の法と政治を考える", 2}, "経済現象を読み解く": {"現代教養科目", "経済現象を読み解く", 2}, "現代社会を読み解く": {"現代教養科目", "現代社会を読み解く", 2},
	"現代文化を読み解く": {"現代教養科目", "現代文化を読み解く", 2}, "現代の差別を考える": {"現代教養科目", "現代の差別を考える", 2}, "現代の医療と生命を考える": {"現代教養科目", "現代の医療と生命を考える", 2},
	"現代の環境を考える": {"現代教養科目", "現代の環境を考える", 2}, "現代の生命倫理・法・経済を考える": {"現代教養科目", "現代の生命倫理・法・経済を考える", 2},
	"大阪大学の歴史": {"先端教養科目", "大阪大学の歴史", 2}, "知性への誘い": {"先端教養科目", "知性への誘い", 2}, "関西は今": {"先端教養科目", "関西は今", 2},
	"女性リーダーとの対話": {"先端教養科目", "女性リーダーとの対話", 2}, "現代キャリアデザイン論": {"先端教養科目", "現代キャリアデザイン論", 2}, "震災の知・復興の知": {"先端教養科目", "震災の知・復興の知", 2},
	"共生社会とアクセシビリティ": {"先端教養科目", "共生社会とアクセシビリティ", 2}, "大学生活と社会": {"先端教養科目", "大学生活と社会", 2}, "健康・医療イノベーション学Ⅰ": {"先端教養科目", "健康・医療イノベーション学Ⅰ", 2},
	"健康・医療イノベーション学II": {"先端教養科目", "健康・医療イノベーション学II", 2}, "今、がんは": {"先端教養科目", "今、がんは", 2}, "おいしく食べる": {"先端教養科目", "おいしく食べる", 2},
	"「くすり」と生体": {"先端教養科目", "「くすり」と生体", 2}, "再生医学": {"先端教養科目", "再生医学", 2}, "病気のバイオサイエンスⅠ": {"先端教養科目", "病気のバイオサイエンスⅠ", 2},
	"病気のバイオサイエンスII": {"先端教養科目", "病気のバイオサイエンスII", 2}, "物理学・宇宙地球科学の先端科学": {"先端教養科目", "物理学・宇宙地球科学の先端科学", 2}, "安全･安心･快適な社会づくり": {"先端教養科目", "安全･安心･快適な社会づくり", 2},
	"環境・エネルギー問題を考える": {"先端教養科目", "環境・エネルギー問題を考える", 2}, "自然環境学": {"先端教養科目", "自然環境学", 2}, "科学・技術と情報の最前線": {"先端教養科目", "科学・技術と情報の最前線", 2},
	"データ科学": {"先端教養科目", "データ科学", 2}, "データ科学と意思決定": {"先端教養科目", "データ科学と意思決定", 2}, "データ科学による課題解決実践": {"先端教養科目", "データ科学による課題解決実践", 2},
	"データ科学による課題解決入門": {"先端教養科目", "データ科学による課題解決入門", 2}, "データ科学のための数理": {"先端教養科目", "データ科学のための数理", 2}, "データ科学特講": {"先端教養科目", "データ科学特講", 2},
	"データ解析の実際": {"先端教養科目", "データ解析の実際", 2}, "工学と現代数学の接点": {"先端教養科目", "工学と現代数学の接点", 2}, "工学への数値シミュレーション": {"先端教養科目", "工学への数値シミュレーション", 2},
	"数値シミュレーションの基礎": {"先端教養科目", "数値シミュレーションの基礎", 2}, "数理モデリングの基礎": {"先端教養科目", "数理モデリングの基礎", 2}, "文理融合に向けた数理科学Ⅰ": {"先端教養科目", "文理融合に向けた数理科学Ⅰ", 2},
	"文理融合に向けた数理科学II": {"先端教養科目", "文理融合に向けた数理科学II", 2}, "先端ビーム科学": {"先端教養科目", "先端ビーム科学", 2}, "「ものづくり」における接合の科学と工学": {"先端教養科目", "「ものづくり」における接合の科学と工学", 2},
	"生命を担う物質-蛋白質": {"先端教養科目", "生命を担う物質-蛋白質", 2}, "実践的危機管理入門": {"先端教養科目", "実践的危機管理入門", 2}, "イノベーション": {"先端教養科目", "イノベーション", 2},
	"技術と社会": {"先端教養科目", "技術と社会", 2}, "Frontier Lectures from University of California Ⅰ": {"先端教養科目", "Frontier Lectures from University of California Ⅰ", 2}, "Frontier Lectures from University of California II": {"先端教養科目", "Frontier Lectures from University of California II", 2},
	"国際政治を考える": {"国際教養１", "国際政治を考える", 2}, "平和の問題を考える": {"国際教養１", "平和の問題を考える", 2}, "国際社会の法を考える": {"国際教養１", "国際社会の法を考える", 2},
	"欧米の文化と社会を知る": {"国際教養１", "欧米の文化と社会を知る", 2}, "アジアの文化と社会を知る": {"国際教養１", "アジアの文化と社会を知る", 2}, "中東の文化と社会を知る": {"国際教養１", "中東の文化と社会を知る", 2},
	"ユーラシアの文化と社会を知る": {"国際教養１", "ユーラシアの文化と社会を知る", 2}, "アフリカの文化と社会を知る": {"国際教養１", "アフリカの文化と社会を知る", 2}, "世界は今": {"国際教養１", "世界は今", 2},
	"世界の事情を英語で学ぶ": {"国際教養１", "世界の事情を英語で学ぶ", 2}, "世界の事情を英語で学ぶ　中級編": {"国際教養１", "世界の事情を英語で学ぶ　中級編", 2},
	"国際コミュニケーション演習（ドイツ語）": {"国際教養２", "国際コミュニケーション演習（ドイツ語）", 2}, "国際コミュニケーション演習（フランス語）": {"国際教養２", "国際コミュニケーション演習（フランス語）", 2}, "国際コミュニケーション演習（ロシア語）": {"国際教養２", "国際コミュニケーション演習（ロシア語）", 2},
	"国際コミュニケーション演習（中国語）": {"国際教養２", "国際コミュニケーション演習（中国語）", 2}, "国際コミュニケーション演習（朝鮮語）": {"国際教養２", "国際コミュニケーション演習（朝鮮語）", 2},
	"国際コミュニケーション演習（スペイン語）": {"国際教養２", "国際コミュニケーション演習（スペイン語）", 2},
	"国際コミュニケーション演習（イタリア語）": {"国際教養２", "国際コミュニケーション演習（イタリア語）", 2}, "地域言語文化演習（ドイツ語）": {"国際教養２", "地域言語文化演習（ドイツ語）", 2}, "地域言語文化演習（フランス語）": {"国際教養２", "地域言語文化演習（フランス語）", 2},
	"地域言語文化演習（ロシア語）": {"国際教養２", "地域言語文化演習（ロシア語）", 2}, "地域言語文化演習（中国語）": {"国際教養２", "地域言語文化演習（中国語）", 2}, "地域言語文化演習（朝鮮語）": {"国際教養２", "地域言語文化演習（朝鮮語）", 2},
	"地域言語文化演習（スペイン語）": {"国際教養２", "地域言語文化演習（スペイン語）", 2}, "地域言語文化演習（イタリア語）": {"国際教養２", "地域言語文化演習（イタリア語）", 2}, "多文化コミュニケーション(日本語）": {"国際教養２", "多文化コミュニケーション(日本語）", 2},
	"特別外国語演習（広東語）Ⅰ": {"国際教養2", "特別外国語演習（広東語）Ⅰ", 2}, "特別外国語演習（インドネシア語）Ⅰ": {"国際教養2", "特別外国語演習（インドネシア語）Ⅰ", 2}, "特別外国語演習（モンゴル語）Ⅰ": {"国際教養2", "特別外国語演習（モンゴル語）Ⅰ", 2},
	"特別外国語演習（モンゴル語）II": {"国際教養2", "特別外国語演習（モンゴル語）II", 2}, "特別外国語演習（タイ語）Ⅰ": {"国際教養2", "特別外国語演習（タイ語）Ⅰ", 2}, "特別外国語演習（ビルマ語）Ⅰ": {"国際教養2", "特別外国語演習（ビルマ語）Ⅰ", 2},
	"特別外国語演習（ビルマ語）II": {"国際教養2", "特別外国語演習（ビルマ語）II", 2}, "特別外国語演習（ヒンディー語）Ⅰ": {"国際教養2", "特別外国語演習（ヒンディー語）Ⅰ", 2}, "特別外国語演習（アラビア語）Ⅰ": {"国際教養2", "特別外国語演習（アラビア語）Ⅰ", 2},
	"特別外国語演習（アラビア語）II": {"国際教養2", "特別外国語演習（アラビア語）II", 2}, "特別外国語演習（トルコ語）Ⅰ": {"国際教養2", "特別外国語演習（トルコ語）Ⅰ", 2}, "特別外国語演習（トルコ語）II": {"国際教養2", "特別外国語演習（トルコ語）II", 2},
	"特別外国語演習（スワヒリ語）Ⅰ": {"国際教養2", "特別外国語演習（スワヒリ語）Ⅰ", 2}, "特別外国語演習（スワヒリ語）II": {"国際教養2", "特別外国語演習（スワヒリ語）II", 2}, "特別外国語演習（ハンガリー語）Ⅰ": {"国際教養2", "特別外国語演習（ハンガリー語）Ⅰ", 2},
	"特別外国語演習（デンマーク語）Ⅰ": {"国際教養2", "特別外国語演習（デンマーク語）Ⅰ", 2}, "特別外国語演習（デンマーク語）II": {"国際教養2", "特別外国語演習（デンマーク語）II", 2}, "特別外国語演習（スウェーデン語）Ⅰ": {"国際教養2", "特別外国語演習（スウェーデン語）Ⅰ", 2},
	"特別外国語演習（スウェーデン語）II": {"国際教養2", "特別外国語演習（スウェーデン語）II", 2}, "特別外国語演習（ポルトガル語）Ⅰ": {"国際教養2", "特別外国語演習（ポルトガル語）Ⅰ", 2}, "特別外国語演習（ポルトガル語）II": {"国際教養2", "特別外国語演習（ポルトガル語）II", 2},
	"英語(Reading)": {"大学英語", "英語(Reading)", 1}, "英語(Writing)": {"大学英語", "英語(Writing)", 1}, "英語(Listening)": {"大学英語", "英語(Listening)", 1},
	"英語(Speaking": {"大学英語", "英語(Speaking)", 1}, "英語上級(Reading)": {"大学英語", "英語上級(Reading)", 1}, "英語上級(Writing)": {"大学英語", "英語上級(Writing)", 1},
	"英語上級(Listening)": {"大学英語", "英語上級(Listening)", 1}, "英語上級(Speaking)": {"大学英語", "英語上級(Speaking)", 1},
	"英語検定訓練コース": {"大学英語", "英語検定訓練コース", 1}, "英語基礎訓練コース": {"大学英語", "英語基礎訓練コース", 1},
	"実践英語": {"実践英語・専門英語", "実践英語", 1}, "専門英語基礎": {"実践英語・専門英語", "専門英語基礎", 1},
	"ドイツ語初級Ⅰ": {"ドイツ語", "ドイツ語初級Ⅰ", 1}, "ドイツ語初級II": {"ドイツ語", "ドイツ語初級II", 1}, "ドイツ語中級": {"ドイツ語", "ドイツ語中級", 1},
	"ドイツ語上級":   {"ドイツ語", "ドイツ語上級", 1},
	"フランス語初級Ⅰ": {"フランス語", "フランス語初級Ⅰ", 1}, "フランス語初級II": {"フランス語", "フランス語初級II", 1}, "フランス語中級": {"フランス語", "フランス語中級", 1},
	"フランス語上級": {"フランス語", "フランス語上級", 1},
	"ロシア語初級Ⅰ": {"ロシア語", "ロシア語初級Ⅰ", 1}, "ロシア語初級II": {"ロシア語", "ロシア語初級II", 1}, "ロシア語中級": {"ロシア語", "ロシア語中級", 1},
	"ロシア語中級２": {"ロシア語", "ロシア語中級２", 1}, "ロシア語上級": {"ロシア語", "ロシア語上級", 1},
	"中国語初級Ⅰ": {"中国語", "中国語初級Ⅰ", 1}, "中国語初級II": {"中国語", "中国語初級II", 1}, "中国語中級": {"中国語", "中国語中級", 1},
	"中国語上級":  {"中国語", "中国語上級", 1},
	"朝鮮語初級Ⅰ": {"朝鮮語", "朝鮮語初級Ⅰ", 1}, "朝鮮語初級II": {"朝鮮語", "朝鮮語初級II", 1}, "朝鮮語中級": {"朝鮮語", "朝鮮語中級", 1},
	"スペイン語初級Ⅰ": {"スペイン語", "スペイン語初級Ⅰ", 1}, "スペイン語初級II": {"スペイン語", "スペイン語初級II", 1}, "スペイン語中級": {"スペイン語", "スペイン語中級", 1},
	"イタリア語初級I": {"イタリア語", "イタリア語初級I", 1}, "イタリア語初級II": {"イタリア語", "イタリア語初級II", 1}, "イタリア語中級": {"イタリア語", "イタリア語中級", 1},
	"総合日本語": {"日本語", "総合日本語", 1}, "専門日本語": {"日本語", "専門日本語", 1},
	"英語選択":      {"第１外国語", "英語選択", 1},
	"ドイツ語初級Ⅰ選択": {"第２外国語", "ドイツ語初級Ⅰ選択", 1}, "ドイツ語初級II選択": {"第２外国語", "ドイツ語初級II選択", 1}, "ドイツ語中級選択": {"第２外国語", "ドイツ語中級選択", 1},
	"フランス語初級Ⅰ選択": {"第２外国語", "フランス語初級Ⅰ選択", 1}, "フランス語初級II選択": {"第２外国語", "フランス語初級II選択", 1},
	"フランス語中級選択": {"第２外国語", "フランス語中級選択", 1}, "ロシア語初級Ⅰ選択": {"第２外国語", "ロシア語初級Ⅰ選択", 1},
	"ロシア語初級II選択": {"第２外国語", "ロシア語初級II選択", 1}, "ロシア語中級選択": {"第２外国語", "ロシア語中級選択", 1},
	"中国語初級Ⅰ選択": {"第２外国語", "中国語初級Ⅰ選択", 1}, "中国語初級II選択": {"第２外国語", "中国語初級II選択", 1}, "中国語中級選択": {"第２外国語", "中国語中級選択", 1},
	"ラテン語初級Ⅰ選択": {"第３外国語", "ラテン語初級Ⅰ選択", 1}, "ラテン語初級II選択": {"第３外国語", "ラテン語初級II選択", 1}, "ラテン語中級Ⅲ選択": {"第３外国語", "ラテン語中級Ⅲ選択", 1},
	"ラテン語中級Ⅳ選択": {"第３外国語", "ラテン語中級Ⅳ選択", 1}, "ギリシャ語初級Ⅰ選択": {"第３外国語", "ギリシャ語初級Ⅰ選択", 1}, "ギリシャ語初級II選択": {"第３外国語", "ギリシャ語初級II選択", 1},
	"ギリシャ語中級Ⅲ選択": {"第３外国語", "ギリシャ語中級Ⅲ選択", 1}, "ギリシャ語中級Ⅳ選択": {"第３外国語", "ギリシャ語中級Ⅳ選択", 1},
	"情報活用基礎": {"情報処理教育科目", "情報活用基礎", 2}, "コンピュータのしくみ": {"情報処理教育科目", "コンピュータのしくみ", 2}, "情報科学入門": {"情報処理教育科目", "情報科学入門", 2},
	"情報社会と倫理": {"情報処理教育科目", "情報社会と倫理", 2}, "情報探索入門": {"情報処理教育科目", "情報探索入門", 2}, "計算機シミュレーション入門": {"情報処理教育科目", "計算機シミュレーション入門", 2},
	"アドバンスド情報リテラシー": {"情報処理教育科目", "アドバンスド情報リテラシー", 2},
	"スポーツ実習 Ａ":      {"健康・スポーツ教育科目", "スポーツ実習 Ａ", 1}, "スポーツ科学": {"健康・スポーツ教育科目", "スポーツ科学", 1}, "健康科学実習A": {"健康・スポーツ教育科目", "健康科学実習A", 1},
	"健康科学": {"健康・スポーツ教育科目", "健康科学", 1}, "スポーツ実習 B": {"健康・スポーツ教育科目", "スポーツ実習 B", 1},
	"哲学基礎Ａ": {"専門基礎教育科目", "哲学基礎Ａ", 2}, "哲学基礎Ｂ": {"専門基礎教育科目", "哲学基礎Ｂ", 2}, "中国哲学基礎": {"専門基礎教育科目", "中国哲学基礎", 2},
	"倫理学基礎": {"専門基礎教育科目", "倫理学基礎", 2}, "インド学基礎": {"専門基礎教育科目", "インド学基礎", 2}, "心理学実験": {"専門基礎教育科目", "心理学実験", 2},
	"感情・人格心理学": {"専門基礎教育科目", "感情・人格心理学", 2}, "心理・行動科学入門": {"専門基礎教育科目", "心理・行動科学入門", 2}, "日本史学基礎Ａ": {"専門基礎教育科目", "日本史学基礎Ａ", 2},
	"日本史学基礎Ｂ": {"専門基礎教育科目", "日本史学基礎Ｂ", 2}, "考古学基礎Ａ": {"専門基礎教育科目", "考古学基礎Ａ", 2}, "考古学基礎Ｂ": {"専門基礎教育科目", "考古学基礎Ｂ", 2},
	"アジア史学基礎Ａ": {"専門基礎教育科目", "アジア史学基礎Ａ", 2}, "アジア史学基礎Ｂ": {"専門基礎教育科目", "アジア史学基礎Ｂ", 2}, "アジア史学基礎Ｃ": {"専門基礎教育科目", "アジア史学基礎Ｃ", 2},
	"西洋史学基礎Ａ": {"専門基礎教育科目", "西洋史学基礎Ａ", 2}, "西洋史学基礎Ｂ": {"専門基礎教育科目", "西洋史学基礎Ｂ", 2}, "人文地理学基礎Ａ": {"専門基礎教育科目", "人文地理学基礎Ａ", 2},
	"人文地理学基礎Ｂ": {"専門基礎教育科目", "人文地理学基礎Ｂ", 2}, "国文学": {"専門基礎教育科目", "国文学", 2}, "国文学資料演習Ａ": {"専門基礎教育科目", "国文学資料演習Ａ", 2},
	"国文学資料演習Ｂ": {"専門基礎教育科目", "国文学資料演習Ｂ", 2}, "国語学": {"専門基礎教育科目", "国語学", 2}, "英米文学入門": {"専門基礎教育科目", "英米文学入門", 2},
	"英語学の基礎": {"専門基礎教育科目", "英語学の基礎", 2}, "フランス文学入門": {"専門基礎教育科目", "フランス文学入門", 2}, "ドイツ文学入門": {"専門基礎教育科目", "ドイツ文学入門", 2},
	"中国の文学": {"専門基礎教育科目", "中国の文学", 2}, "中国古典演習": {"専門基礎教育科目", "中国古典演習", 2}, "比較文学入門": {"専門基礎教育科目", "比較文学入門", 2},
	"芸術の始まり": {"専門基礎教育科目", "芸術の始まり", 2}, "文芸学": {"専門基礎教育科目", "文芸学", 2}, "演劇学入門": {"専門基礎教育科目", "演劇学入門", 2},
	"音楽学": {"専門基礎教育科目", "音楽学", 2}, "日本美術史": {"専門基礎教育科目", "日本美術史", 2}, "東洋美術史": {"専門基礎教育科目", "東洋美術史", 2},
	"西洋美術史": {"専門基礎教育科目", "西洋美術史", 2}, "民俗学": {"専門基礎教育科目", "民俗学", 2}, "日本学基礎": {"専門基礎教育科目", "日本学基礎", 2},
	"日本語学基礎": {"専門基礎教育科目", "日本語学基礎", 2}, "法学概論": {"専門基礎教育科目", "法学概論", 2}, "政治学概論": {"専門基礎教育科目", "政治学概論", 2},
	"国際関係論入門": {"専門基礎教育科目", "国際関係論入門", 2}, "知的財産モラル": {"専門基礎教育科目", "知的財産モラル", 2}, "経済学Ａ": {"専門基礎教育科目", "経済学Ａ", 2},
	"経済学Ｂ": {"専門基礎教育科目", "経済学Ｂ", 2}, "社会思想史": {"専門基礎教育科目", "社会思想史", 2}, "社会学入門": {"専門基礎教育科目", "社会学入門", 2},
	"社会思想史入門": {"専門基礎教育科目", "社会思想史入門", 2}, "現代社会論入門": {"専門基礎教育科目", "現代社会論入門", 2}, "教育学概論": {"専門基礎教育科目", "教育学概論", 2},
	"基礎人間科学概論": {"専門基礎教育科目", "基礎人間科学概論", 2}, "アジア言語文化研究入門": {"専門基礎教育科目", "アジア言語文化研究入門", 2}, "中東・アフリカ言語文化研究入門": {"専門基礎教育科目", "中東・アフリカ言語文化研究入門", 2},
	"ヨーロッパ・アメリカ言語文化研究入門": {"専門基礎教育科目", "ヨーロッパ・アメリカ言語文化研究入門", 2}, "統計学Ａ-Ⅰ": {"専門基礎教育科目", "統計学Ａ-Ⅰ", 2}, "統計学Ａ-II": {"専門基礎教育科目", "統計学Ａ-II", 2},
	"図学Ａ": {"専門基礎教育科目", "図学Ａ", 2}, "図学Ｂ-Ⅰ": {"専門基礎教育科目", "図学Ｂ-Ⅰ", 2}, "図学Ｂ-II": {"専門基礎教育科目", "図学Ｂ-II", 2},
	"Japan in the World ": {"専門基礎教育科目", "Japan in the World ", 2}, "Anthropology and Contemporary Global Issues": {"専門基礎教育科目", "Anthropology and Contemporary Global Issues", 2}, "Introduction to Social Psychology": {"専門基礎教育科目", "Introduction to Social Psychology", 2},
	"Politics in Post-War Japan": {"専門基礎教育科目", "Politics in Post-War Japan", 2}, "Media Sociology": {"専門基礎教育科目", "Media Sociology", 2}, "Cross Cultural Psychology": {"専門基礎教育科目", "Cross Cultural Psychology", 2},
	"アーカイブズの世界に触れる": {"基礎セミナー", "アーカイブズの世界に触れる", 2}, "アカデミック・ライティング入門 Ａ": {"基礎セミナー", "アカデミック・ライティング入門 Ａ", 2}, "アカデミック・ライティング入門 Ｂ": {"基礎セミナー", "アカデミック・ライティング入門 Ｂ", 2},
	"飯舘村環境放射線研修": {"基礎セミナー", "飯舘村環境放射線研修", 2}, "池島プロジェクト": {"基礎セミナー", "池島プロジェクト", 2}, "「囲碁」で論理的思考を養おう": {"基礎セミナー", "「囲碁」で論理的思考を養おう", 2},
	"イノベーション思考": {"基礎セミナー", "イノベーション思考", 2}, "インターネットを使って情報発信しよう": {"基礎セミナー", "インターネットを使って情報発信しよう", 2}, "インドネシアの歴史と社会": {"基礎セミナー", "インドネシアの歴史と社会", 2},
	"Introduction to Laser Engineering": {"基礎セミナー", "Introduction to Laser Engineering", 2}, "宇宙への旅：材料の役割": {"基礎セミナー", "宇宙への旅：材料の役割", 2}, "英語で考える": {"基礎セミナー", "英語で考える", 2},
	"英語の科学番組をインターネットで聞く": {"基礎セミナー", "英語の科学番組をインターネットで聞く", 2}, "映像表現入門": {"基礎セミナー", "映像表現入門", 2}, "エネルギーと環境": {"基礎セミナー", "エネルギーと環境", 2},
	"大阪大学リーダーズ・アカデミー": {"基礎セミナー", "大阪大学リーダーズ・アカデミー", 2}, "化学実験の基礎": {"基礎セミナー", "化学実験の基礎", 2}, "科学者という仕事を考える": {"基礎セミナー", "科学者という仕事を考える", 2},
	"化学フロンティア　Ⅰ": {"基礎セミナー", "化学フロンティア　Ⅰ", 2}, "化学フロンティア　II": {"基礎セミナー", "化学フロンティア　II", 2}, "化学フロンティア　Ⅲ": {"基礎セミナー", "化学フロンティア　Ⅲ", 2},
	"化学フロンティア　Ⅳ": {"基礎セミナー", "化学フロンティア　Ⅳ", 2}, "化学フロンティア　Ⅵ": {"基礎セミナー", "化学フロンティア　Ⅵ", 2}, "化学フロンティア　Ⅸ": {"基礎セミナー", "化学フロンティア　Ⅸ", 2},
	"確率モデルとその応用Ⅰ": {"基礎セミナー", "確率モデルとその応用Ⅰ", 2}, "楽器を作ろう": {"基礎セミナー", "楽器を作ろう", 2}, "カフェ的対話法": {"基礎セミナー", "カフェ的対話法", 2},
	"共生の人間学入門セミナー": {"基礎セミナー", "共生の人間学入門セミナー", 2}, "キラルテクノロジーの基礎": {"基礎セミナー", "キラルテクノロジーの基礎", 2}, "口の中で機能を支えるマテリアル": {"基礎セミナー", "口の中で機能を支えるマテリアル", 2},
	"暮らしの中の放射線": {"基礎セミナー", "暮らしの中の放射線", 2}, "Global Citizenship Seminar": {"基礎セミナー", "Global Citizenship Seminar", 2}, "芸術を通して脳科学を学ぼう": {"基礎セミナー", "芸術を通して脳科学を学ぼう", 2},
	"現象と数理モデルⅠ": {"基礎セミナー", "現象と数理モデルⅠ", 2}, "建築・町を見る": {"基礎セミナー", "建築・町を見る", 2}, "国際協力と人類学": {"基礎セミナー", "国際協力と人類学", 2},
	"子どもの現在": {"基礎セミナー", "子どもの現在", 2}, "壊れる/壊す/創る": {"基礎セミナー", "壊れる/壊す/創る", 2}, "コントラクトブリッジで考える力をつけよう": {"基礎セミナー", "コントラクトブリッジで考える力をつけよう", 2},
	"最先端の情報システム": {"基礎セミナー", "最先端の情報システム", 2}, "裁判員裁判を考える": {"基礎セミナー", "裁判員裁判を考える", 2}, "材料の化学": {"基礎セミナー", "材料の化学", 2},
	"材料プロセス入門": {"基礎セミナー", "材料プロセス入門", 2}, "様々な科学でみられる数理と応用Ⅰ": {"基礎セミナー", "様々な科学でみられる数理と応用Ⅰ", 2}, "産業革命はなぜイギリスで始まったのか": {"基礎セミナー", "産業革命はなぜイギリスで始まったのか", 2},
	"システム・制御の新しいパラダイム": {"基礎セミナー", "システム・制御の新しいパラダイム", 2}, "自然の読み方": {"基礎セミナー", "自然の読み方", 2}, "実践グローバルリーダーシップ": {"基礎セミナー", "実践グローバルリーダーシップ", 2},
	"自分のウェブページを作ろう": {"基礎セミナー", "自分のウェブページを作ろう", 2}, "ジェロントロジーを学ぶ": {"基礎セミナー", "ジェロントロジーを学ぶ", 2}, "社会科学における数学的方法": {"基礎セミナー", "社会科学における数学的方法", 2},
	"数学の楽しみ　１Ａ": {"基礎セミナー", "数学の楽しみ　１Ａ", 2}, "数学の楽しみ　１Ｂ": {"基礎セミナー", "数学の楽しみ　１Ｂ", 2}, "数学の楽しみ　１Ｃ": {"基礎セミナー", "数学の楽しみ　１Ｃ", 2},
	"数学の楽しみ　１Ｄ": {"基礎セミナー", "数学の楽しみ　１Ｄ", 2}, "数理生物学入門": {"基礎セミナー", "数理生物学入門", 2}, "数理モデリングの実践": {"基礎セミナー", "数理モデリングの実践", 2},
	"スマート社会を実現する情報システム": {"基礎セミナー", "スマート社会を実現する情報システム", 2}, "生物化学工学の未来探訪": {"基礎セミナー", "生物化学工学の未来探訪", 2}, "生物工学入門": {"基礎セミナー", "生物工学入門", 2},
	"精密科学の世界Ⅰ": {"基礎セミナー", "精密科学の世界Ⅰ", 2}, "精密科学の世界II": {"基礎セミナー", "精密科学の世界II", 2}, "生命と機械の融合を目指したものづくり": {"基礎セミナー", "生命と機械の融合を目指したものづくり", 2},
	"臓器移植の諸相": {"基礎セミナー", "臓器移植の諸相", 2}, "大学教員という仕事": {"基礎セミナー", "大学教員という仕事", 2}, "大規模災害と危機管理を考える": {"基礎セミナー", "大規模災害と危機管理を考える", 2},
	"多文化コミュニケーションセミナーⅠ": {"基礎セミナー", "多文化コミュニケーションセミナーⅠ", 2},
	"蛋白質科学入門Ⅰ":          {"基礎セミナー", "蛋白質科学入門Ⅰ", 2}, "知能とコンピュータ": {"基礎セミナー", "知能とコンピュータ", 2}, "Discovery Seminar：はじめてのリサーチ": {"基礎セミナー", "Discovery Seminar：はじめてのリサーチ", 2},
	"哲学の森": {"基礎セミナー", "哲学の森", 2}, "電子・イオンが創るエネルギー技術": {"基礎セミナー", "電子・イオンが創るエネルギー技術", 2}, "電子・光科学への招待": {"基礎セミナー", "電子・光科学への招待", 2},
	"特殊相対論から量子色力学まで": {"基礎セミナー", "特殊相対論から量子色力学まで", 2}, "ナノサイエンス": {"基礎セミナー", "ナノサイエンス", 2}, "ナノテクノロジーが拓く量子の世界": {"基礎セミナー", "ナノテクノロジーが拓く量子の世界", 2},
	"ナノテクノロジーの最前線": {"基礎セミナー", "ナノテクノロジーの最前線", 2}, "ナノの科学技術最前線": {"基礎セミナー", "ナノの科学技術最前線", 2}, "ナノマシンを考える": {"基礎セミナー", "ナノマシンを考える", 2},
	"２１世紀の難問を総合的に考える": {"基礎セミナー", "２１世紀の難問を総合的に考える", 2}, "ネットを知り、ネットを使いこなす": {"基礎セミナー", "ネットを知り、ネットを使いこなす", 2}, "脳と行動": {"基礎セミナー", "脳と行動", 2},
	"博物学の世界を覗く": {"基礎セミナー", "博物学の世界を覗く", 2}, "博物館体験コース": {"基礎セミナー", "博物館体験コース", 2}, "はじめてのアカデミック・ライティング": {"基礎セミナー", "はじめてのアカデミック・ライティング", 2},
	"発明・発見？？": {"基礎セミナー", "発明・発見？？", 2}, "ピアサポート入門": {"基礎セミナー", "ピアサポート入門", 2}, "光と物質とエネルギー": {"基礎セミナー", "光と物質とエネルギー", 2},
	"光とプラズマ": {"基礎セミナー", "光とプラズマ", 2}, "非線形力学入門　": {"基礎セミナー", "非線形力学入門　", 2}, "人を動かす仕掛けの仕組み": {"基礎セミナー", "人を動かす仕掛けの仕組み", 2},
	"ビブリオバトル入門": {"基礎セミナー", "ビブリオバトル入門", 2}, "船と海の科学": {"基礎セミナー", "船と海の科学", 2}, "分子と生命": {"基礎セミナー", "分子と生命", 2},
	"放射光とレーザー": {"基礎セミナー", "放射光とレーザー", 2}, "放射線と放射能": {"基礎セミナー", "放射線と放射能", 2}, "街に出てサイエンスカフェをやってみよう": {"基礎セミナー", "街に出てサイエンスカフェをやってみよう", 2},
	"ミクスド・リアリティへの誘い": {"基礎セミナー", "ミクスド・リアリティへの誘い", 2}, "メカトロニクス入門": {"基礎セミナー", "メカトロニクス入門", 2}, "目で観る物性論": {"基礎セミナー", "目で観る物性論", 2},
	"免疫のしくみ：病気と感染": {"基礎セミナー", "免疫のしくみ：病気と感染", 2}, "ものづくりと破壊の科学": {"基礎セミナー", "ものづくりと破壊の科学", 2}, "ものづくりフロンティア": {"基礎セミナー", "ものづくりフロンティア", 2},
	"有機金属化学入門": {"基礎セミナー", "有機金属化学入門", 2}, "流体現象を解きほぐす数理科学Ⅰ": {"基礎セミナー", "流体現象を解きほぐす数理科学Ⅰ", 2}, "量子力学の不思議な世界": {"基礎セミナー", "量子力学の不思議な世界", 2},
	"Academic Writing Seminar Ⅰ": {"基礎セミナー", "Academic Writing Seminar Ⅰ", 2}, "ITの基礎と応用": {"基礎セミナー", "ITの基礎と応用", 2}, "イノベーションのためのパトス・ロゴス・エトス": {"基礎セミナー", "イノベーションのためのパトス・ロゴス・エトス", 2},
	"宇宙物質形成の起源": {"基礎セミナー", "宇宙物質形成の起源", 2}, "大阪の町を読む": {"基礎セミナー", "大阪の町を読む", 2}, "確率モデルとその応用II": {"基礎セミナー", "確率モデルとその応用II", 2},
	"基礎からのアカデミック・スキル": {"基礎セミナー", "基礎からのアカデミック・スキル", 2}, "キャンパスデザインプロジェクト": {"基礎セミナー", "キャンパスデザインプロジェクト", 2}, "Critical Thinking Skills Seminar": {"基礎セミナー", "Critical Thinking Skills Seminar", 2},
	"経営者と学ぶリーダーシップ": {"基礎セミナー", "経営者と学ぶリーダーシップ", 2}, "現象と数理モデルII": {"基礎セミナー", "現象と数理モデルII", 2}, "現代社会と著作権": {"基礎セミナー", "現代社会と著作権", 2},
	"Contemporary Japan Seminar": {"基礎セミナー", "Contemporary Japan Seminar", 2}, "様々な科学でみられる数理と応用II": {"基礎セミナー", "様々な科学でみられる数理と応用II", 2}, "事業戦略と知的財産マネジメント": {"基礎セミナー", "事業戦略と知的財産マネジメント", 2},
	"社会・技術の変容とアートの役割": {"基礎セミナー", "社会・技術の変容とアートの役割", 2}, "数学の楽しみ　２Ａ": {"基礎セミナー", "数学の楽しみ　２Ａ", 2}, "数学の楽しみ　２Ｂ": {"基礎セミナー", "数学の楽しみ　２Ｂ", 2},
	"数学の楽しみ　２Ｃ": {"基礎セミナー", "数学の楽しみ　２Ｃ", 2}, "数学の楽しみ　２Ｄ": {"基礎セミナー", "数学の楽しみ　２Ｄ", 2}, "数理医学入門": {"基礎セミナー", "数理医学入門", 2},
	"相対論的ゲームを作る": {"基礎セミナー", "相対論的ゲームを作る", 2}, "多文化コミュニケーションセミナーII": {"基礎セミナー", "多文化コミュニケーションセミナーII", 2},
	"Presentation Skills Seminar": {"基礎セミナー", "Presentation Skills Seminar", 2}, "平和研究入門": {"基礎セミナー", "平和研究入門", 2},
	"リーダーシップを考える": {"基礎セミナー", "リーダーシップを考える", 2}, "流体現象を解きほぐす数理科学II": {"基礎セミナー", "流体現象を解きほぐす数理科学II", 2}, "ロボットを通して人を知る": {"基礎セミナー", "ロボットを通して人を知る", 2},
}
