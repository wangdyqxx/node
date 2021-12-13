/**
 * @Author: 李大双
 * @Description: 剑指 Offer 题解
 * @File: array_test
 * @Date: 2021/6/7 上午11:36
 */
package code

import (
	"fmt"
	"testing"
)

func TestDuplicate(t *testing.T) {
	testCase := []struct {
		Input []int
	}{
		{
			Input: []int{0, 1, 2, 3, 4, 4, 5},
		},
		{
			Input: []int{3, 4, 2, 0, 0, 1},
		},
	}

	for _, v := range testCase {
		fmt.Println(v)
		//output := Duplicate(v.Input)
		output := findRepeatNumber(v.Input)
		t.Log(output)
	}
}

func TestFind(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	testCase := []struct {
		Target int
		Matrix [][]int
		Ret    bool
	}{
		{
			Target: 5,
			Matrix: matrix,
			Ret:    true,
		},
		{
			Target: 31,
			Matrix: matrix,
			Ret:    false,
		},
		{
			Target: 31,
			Matrix: [][]int{},
			Ret:    false,
		},
		{
			Target: 7,
			Matrix: [][]int{
				{1, 2, 8, 9},
				{4, 7, 10, 13},
			},
			Ret: true,
		},
	}

	for _, v := range testCase {
		if ret := findNumberIn2DArray(v.Target, v.Matrix); ret != v.Ret {
			t.Errorf("expect %v,but is %v", v.Ret, ret)
		}
	}
}

func TestReplaceSpace(t *testing.T) {
	testCase := []struct {
		Input  string
		Output string
	}{
		{
			Input:  "A B",
			Output: "A%20B",
		},
		{
			Input:  " A B",
			Output: "%20A%20B",
		},
		{
			Input:  "   ",
			Output: "%20%20%20",
		},
	}

	for _, v := range testCase {
		if output := replaceSpace(v.Input); output != v.Output {
			t.Errorf("expect %v,but is %v", v.Output, output)
		}
	}
}

func TestPrintMatrix(t *testing.T) {
	testCase := []struct {
		Matrix [][]int
		Ret    []int
	}{
		{
			Matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			Ret: []int{1, 2, 4, 3},
		},
		{
			Matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			Ret: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			Matrix: [][]int{
				{2, 5, 8},
				{4, 0, -1},
			},
			Ret: []int{2, 5, 8, -1, 0, 4},
		},
	}

	var isEqual = func(a, b []int) bool {
		if len(a) != len(b) {
			return false
		}
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
	for _, v := range testCase {
		if ret := spiralOrder(v.Matrix); !isEqual(ret, v.Ret) {
			t.Errorf("expect %v,but is %v", v.Ret, ret)
		}
	}
}

func TestFirstNotRepeatingChar(t *testing.T) {
	testCase := []struct {
		Input  string
		Output byte
	}{
		{
			Input:  "uindrseqbljlhqvlwvgdebeihttirikuahlikgnahvgnptmqwbovmuwesxkvcitcwrwrucsbbfqvldridfviduqvmfcmeiphoqupbitnwdbvevouoaetisdmgvvvwoglwtgjrpcbghxkrkjthetxeexbphbjiehaicuicgnirslhdstgmqcdnlulpdpadjdltfouwhfqicfcqntnpeqaohslwkhbvflxaudembsrsindluthxapnmrqinsivbxmkohubvmmmpmklbfrmeuvdrhptdmelmjjefgbsqqlbqhvsmswwxrlkutadqbbeisbgfrcivvoxmxxptrscxnjjvtajfhqiucdihihcutxhlonlomfdnbwanrcnbarojsajseqrgkuqgkcnvrghxnmbclfomktwfaakodeecsglufvobkgoqsbrhdiuhxqbcndkmxuertupngvnkgwrfwgdbiurbooxariklwarjgavsuddoveipltessrssxwxgvrrtkisnbavkbpaphicxhapjjpkccakepuafjdaswfwfmbdmuuaveukxvnpkgmhcjcpqntssthjlvrngugbhrivtwmbrvrprrlfmvwjdkonfmgnepqoxwfcvefjihisarwskjfmqrjlkbbugfawmkqgfxgpokkxxivqbqimwsccdekjegwcxhmmuhpstxfpofwmrmhrxeptxwbvxsesheijjfsshgrrwckjkmbslpbngnwhokjnujtiepfrdhiwwgbffixaidabacaibummwxgxowsewlqfxrkarjrkqxmxwobqbcoowmbggtpoqadhqdxlhvxrkfuwpxxgnchudreoeuefkqhlrmwwfvjexvbxdhtdngwvoohpjtdbdbesceiafrhenfljleegsencfmbauxlltfueudnjxsiiggsfwiuidgktsbtcvxtdllviocfdbonjciosucbjidwxnogmnveqkcrxbpamfxwxiugjgrfpstromuxxdiodqeoqdlfulttraquskhwfholbrcbchijlgqvuwxvejedikvgetlmrcmeampdgjvmbdovkcjilbralhmniwvbeandldnudkpjvhoqtfdwllridwljfvflfdrbqadvabggwsiexrthrngexpebuhtefkqgkkjoopsmunesfotsprxuaswwenhkdvsspkppulecahkkvccqngeaoijjgvsfqfpvphvhdnkkcqsebhijkvfpqjmarbkpejagtbtisnflrrvawmvfxeccxtrmorihgslxlqrqcmouojjfhcieuwlrcqhevmveookgonxqdbtgqjimsidnaaiuchwkfkpxfptuvhfqemojixqvhgokdekdwbomptqqlfiaiptxvgfmovdqxupjxjpoxuplsagxpgpmvtcpawkrrthvclhdbpqeucchxptdceswhnqocmeocpgthkgxmxggwlnantwwuoqpmnpvgateitxlocmhnihmfgjnvrggenbhnjfubtoteojojjkjpcnwqthxhlfukingjletwnxdnjwrgjopqoxtqcvwsakqxxumbtcblufdmdvxfkshitenmenkxjvblsoiyhjpakinimwxhcebabgsvqftfvwjnstltjawhwipkubadtoxqrutkwxjnmfoowtnvqplvqokcuwlkmxxhboampcdwokjfxggtnojebagxlwaeowtomubtbfsrufkttugfpnxmipkcsphqafjuxovwpcgonhiqmomsweunoeqkpkxxsdksmufowqpmkontccckcdbwrfwamikananakgjkahndrepemcgxecgpltvdbpoexemnrejdephuuxhfcubxlbdrmhvmeqmtdhnbkwnidigxdantmkckijiecavkpumegrbveffclcltmibjcstrrauphfxxssgxkkirapiihnlbrodvfostahqdkajqtrrfwdsemwxlucbbjjspnnqjnpnpmimhulgoskwpsactexmkfdhaihnlggqeunqevxfrpiwskhrhgfluelshqejavomjshfomvgpugqesbtakvxqwrtguuebcgqphocglfrircfvuikfbviomfrsnpvvlftwrkbmjpvdpgvohtmxcxwnuhwojnfvsfwvdlaxxlmafihpussffcawjpaxdwerwfbvsbipljualcnhseealvqiqfiaiskaafufaubvhjglktfhbdsaskmedroxkrxshnggumcbtdcablawkodmnafkjekuiecqlvbxocfcwipaicgndfafjhtjcelakrecmntjxeqjqgxkxapuobbcaxfrpsktjswxdfvugmmgjwiphnsclqwmranthpiueffaxvhplajqsrtoxbixbdpcfkbpkjkelemubwachcpxcniaqmkmasmvlfcubcdkxfkupgcgbrvmbmgnjgfbawmmdritdrkppswatwtdjemhifhmshunkvaivhteqnwkdcrpfxmrafupfhbgligbvrqjkvcgbiqudvtblhfgetffduvsfhsmjimgruxrvbqniluapaoniwhqhltbxvrphmlisfaomqoecmdbbrgujgsbdelkbddcgmpggogfonssxisphbwljlvhwhewmpqugxgbubqfeilobxxhtcxmxvjtbuavxlaghhjiemihvrsjxbleutpuumgndtnocwkpkdaupbmcsahcbwoelmgnwqummmekwhpahtdvehoprfcciwqphfwppscruimikiximhbkdomovbdalsihioncpaclevxawqcqthtmuogolkrvlipropmnmecttmlecpapdlrcqxopfjobgsvhcadqiudqjoscpvjguddnorldkpqtoocrtkcutkvdpbkkekenpqjmxmuxccamkwwxrdrbdmameptvcgsltqdasicoouvdtbranexokmtkioptqblrruaihvfvvebjigkuwwgearxvmimlgmxgpvbnotepprknjxcunheclsvxxmmoufvdhrwxqlrkmvcaxtbiqrjrnrsfrkvttnpxbspnetcofhpcbcphrchivnrskkalsvdllhtisnaerxbcxutcpmnquwrutmmvrtchhheplhdmbfkrrrxalosjbebtdqtsjbdcuuatgldnubigijbkehxsmhsadwtcbwinlnilthiqslvsduspkjuunpetcofujbqhtkmalnjmllhulckqodbgffxdiqjajecbmlekvepkirxnguqnpnfjdptinuphcoxkoqhqrmkkawmngqwvfpcuxgpekkprupplsbwshmsduaitlhcjargacduewtqgnxbtpwrkawpncarxlsnardkftmsdjflmhoplghandjisupacixxthwokctbxekqwbmhtoutgrogxsdnmcaqrsqrpigbigqhsmrksxgkdbdbxkwimpehltcffuhrphbaqgvabjduudbrxgwiljkwijvaugnleamskjijhnvcspagsfnhjielrvxddxqfgaipfflffksmvgaioaclfsjhovcotoaiwhhnmrudsqmsepsgdofnxjjcgdejutmdtgauuvuepajlikebtpgfimudswplfwgaodtwqjuaujlbqqartevsqesmrpknoanfprqcwochqdrcvquovtsogptduvfiaislfvpxwsqitodxwckprpfxtqhbdixklhtsajqgbinfpdqaehrwhnwggmggnbstolrcbmlhpfjurbknpwbrfhvdanrjcwnvsxudagbkhdlwujqvwpbbgmsfrkgutkqwxqfemaaokrfdocsbbktwaugoadgvqcdhmqkworbvxqwdpmbjkqghmapevrsnaaparwnbvcfnhhgfqplfwiswrxxistamulcbxcmvkbcxsgftbmbvqlxrhbmcujkktdvcgnjkfaclxndtjkdbgqvbqxicgmpowbjsfvspvxtrovjwbbkaapfgcqbifdosnmutvfbgdmhiqvflasedhafqnmrlmwbuhmptnkvvojnjngiepmdktewwsuxernpfmktsgwwnxfvmuotpfxfntcrsjpldmringfpkpievvbgbghevlmuticukrfvahrmtkvckjbxxrmkilbeojsxgbvbhsgvkhgtjddoreipbkwiqbkecobfexswiwttlohnfwokixfmfrpadtudrssiobbvhctpbeesiogwhdqojoalxpxikwuheilvsrxbfbnckjeswhkfviihvwrfxkxcomhvsjojhltospooumkstqqnvxmrowlrtufhxkipxnfecchklrqljfwnglxeharemawsbvnuentnapewjhnzbieisjtomdqtjkigisgopmgcfqxcxninrisupchhmfrosxgdtakhrrnbbnxovvbfjgwdgnlxoxswhsfjeijvivqemoekkmrttscthlpglaorarhibrucuhikfkphqmhowloobeumlikqmmlatesupiekwjhdvntnjbakohrbobimmujkbatfxpmgfmmqsgwsffcbenwgdrhlqrsmfghevnjfdlxupihmkptsefubmixhacnxngpjrckxhvhrelolqwevoqklhxbplakmsxdetvpuddmvaodpglmsknbvfftstckjurbntwqenuhcurxqvotrxfpootqkobeatoduihwhpjxxrvboggkcxagnrrfwfuaqplirlnafprumnkcjxlpvanfmqwuoaupqptqeulpwvnhbahbkihustenkbdushjakemufdhjllnowmvgpdxbrhxfonpgcjslvwgdmajicqqmdxubrofvdodfsedjghpbjncgkckeudxwascljlraoodwvgpvojcqgbalntslsfrwnfcdvidsdrvwhscskpalubxeobapkgpdsqcjpwkednraijmbcreplwijofniggavpdlwfmwnvsaridhbppeolakquhamneffnfmbruivassdaaikxiaxupgdpgkfojvkkagcsmqnweofwikvevrsottkbtcldoruakajinnlgxmpddcrmohaktsdrxjelrbmfdthspikxeocqqdordrqwjtaxihswcubtfomksaxddcvtadnrtqkmdnacgduudtdbhsgpfomrdiaotcfwqlxccqeelgtnkflgixpjcriroruhbbgpidwkggdevpfratqebewaxcudseaikqjuuhpsguvtwxniofskkslfnrphatmxngneuituefcxufisrsjitqxwufrgidbwlmkrqojpwpljasaiukwwthrudhcocoguerakhajdncxbrnuavoqeuwsamwdqgutbaoixcgeibpoajhedooqcewiqvddedmanxljjjcbojgbmwabkfbvamgnfpncdcxoaqhmgurifpwpgrpctlqpwmqraknjltneknsphtwbnbiahknxipsovljkivlpggvldeveeopvoqlvfjbratadttlcecomllpkdgiloeedquivsnmxkfcvrkwaohgbrbvjklkktgwtlfgafqgbigheajrvffvvkkmibcedfmnwasopoqgxjohjoqnijeaifuwiwmogkwqlfmibuwecvnulvhkbsukscqlqlsxjulillmlgjkrbmllcunhbqmbujftqgkpwkvemthigednfohxpsduotwfnknfjbexflcfieaosnlhndsorbkcdlahovwbccshmrlcofowrmqajluiqaarnqtxokbaddswftiexiahmstpbonwmhvqgcpmfrocvtwfhjrtsupscfmvwfvaoolanrlgdsvgoseltdnoxrdglockhwlvffaakgrxjfnfbxnbprfvpwmexnhfekjnkenbohhmwlqoteiwgtjsrnceptbgwkfcmtkliwwqskmsoihmnbjsvnmfkwbwemijdtpnajgrousbrdaagenqlgeaiixfgcbfhceaxkwxbpksfouuqcvcerqecpdtvtsubbadmaatdnmnqhladeiapejkxffoagbwqbvppssvvvhlvhntatxlvqgjxbejpvxemqbdjknuogumrwbognklmvjsldrktpeowiuaapbjgktxwfjiferxgmafbcerteqlvpqvxbgdwiufdctkwptadtujmifppudubpmdtiedmqhnihquansnjufpbuumhhmidphkwusjdsdaocavtauhsvtgcoqhufmacwcbxvjmagkounkoqpcnoanhgwsjvgtlwgvbpdbufekosgagfsmadmvbonkrtcbspoabugkcjeebqhqwfjcqlqjvabaqecofgwskqplgup",
			Output: 'y',
		},
	}

	for _, v := range testCase {
		if output := firstUniqChar(v.Input); output != v.Output {
			t.Errorf("expect %v,but is %v", string(v.Output), string(output))
		}
	}
}

func TestIncreasingTriplet(t *testing.T) {
	numslist := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 2, 1},
	}
	for _, nums := range numslist {
		fmt.Println(increasingTriplet(nums))
	}
}

func TestProductExceptSelf(t *testing.T) {
	numsList := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 2, 1},
	}
	for _, nums := range numsList {
		fmt.Println(productExceptSelf(nums))
	}
}

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		//{1, 4, 7, 11, 15},
		//{2, 5, 8, 12, 19},
		//{3, 6, 9, 16, 22},
		{10},
		//{18, 21, 23, 26, 30},
	}
	target := 10
	fmt.Println(searchMatrix(matrix, target))
	//time.Sleep(time.Second)
}

func TestBinarySearch(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	for _, nums := range matrix {
		target := nums[0]
		fmt.Println(binarySearch(nums, target))
	}
}

func TestSingleNumber(t *testing.T) {
	matrix := [][]int{
		{1, 1, 2, 2, 3, 4, 4},
	}
	for _, nums := range matrix {
		fmt.Println(singleNumber(nums))
	}
}

func TestMajorityElement(t *testing.T) {
	matrix := [][]int{
		{1, 1, 2, 2, 2, 3, 4, 4},
	}
	for _, nums := range matrix {
		fmt.Println(majorityElement(nums))
	}
}

func TestMerge(t *testing.T) {
	//a1 := []int{0}
	//a2 := []int{8}
	//a1 := []int{ 1,2,3,0,0}
	//a2 := []int{ 2,3}
	a1 := []int{0, 0, 0, 0, 0}
	a2 := []int{1, 2, 3, 4, 5}
	merge(a1, 0, a2, 5)
	fmt.Println(a1)
}

func TestLargestNumber(t *testing.T) {
	nums := []int{3, 30, 34, 5, 9}
	t.Log(largestNumber(nums))
}

func TestLengthOfLongestSubstring(t *testing.T) {
	strs := []string{
		"dvdf",
		"pwwkew",
		" ",
		"aaaa",
	}
	for _, str := range strs {
		t.Log(lengthOfLongestSubstring(str))
	}
}

func TestMax(t *testing.T) {
	arra := [][]int{
		{1, 1},
		{1, 8, 6, 2, 5, 4, 8, 3, 7},
		{4, 3, 2, 1, 4},
		{1, 2, 1},
		{1, 2, 1, 0},
		{},
	}
	for _, arr := range arra {
		t.Log(maxArea(arr))
	}
}

func TestThreeSum(t *testing.T) {
	arrs := [][]int{
		{-1, 0, 1, 2, -1, -4},
		//{1, -1, -1, 0},
		//{0,0,0,0},
	}
	for _, arr := range arrs {
		t.Log(threeSum(arr))
	}
}

func TestCoinChange(t *testing.T) {
	coins := []int{1, 2, 5}
	fmt.Println(coinChange(coins, 12))
	fmt.Println(Dp_CoinChange(coins, 12))
	a := make(chan int, 1)
	b := make(chan int, 1)
	a <- 1
	b <- 2
	for {
		select {
		case aa := <-a:
			fmt.Println(aa)
		case bb := <-b:
			fmt.Println(bb)
		}
	}

}

//func  coinChange(vector<int>& coins, int amount) int {
//	if (amount == 0) return 0;
//	int ans = INT_MAX;
//	for (int coin : coins) {
//	// 金额不可达
//	if (amount - coin < 0) continue;
//	int subProb = coinChange(coins, amount - coin);
//	// 子问题无解
//	if (subProb == -1) continue;
//	ans = min(ans, subProb + 1);
//	}
//	return ans == INT_MAX ? -1 : ans;
//}

func TestFib1(t *testing.T) {
	//fmt.Println(fib1(20))
	//fmt.Println(fib2(20))
	fmt.Println(fib3(20))
}

func TestRob(t *testing.T) {
	arr := []int{1, 2, 4, 1, 7, 8, 3}
	fmt.Println(rob(arr))
}

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 2, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}

func TestArray(t *testing.T) {
	a1 := [5]int{1, 1, 1, 0, 0}
	arr1(a1)
	fmt.Println(a1)
	a2 := []int{1, 1, 1}
	arr2(a2)
	fmt.Println(a2)
}
