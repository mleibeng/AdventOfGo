package main

import (
	"fmt"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var pl = fmt.Println

func calculateMapScore(left, right []int) int {
	rigthFreq := make(map[int]int)
	for _, num := range right {
		rigthFreq[num]++
	}

	similarityScore := 0
	for _, num := range left {
		if freq, exists := rigthFreq[num]; exists {
			similarityScore += num * freq
		}
	}
	return similarityScore
}

func main() {

	s := []int{53906, 35867, 61313, 23620, 96434, 70853, 81024, 74096, 95143, 58372, 31935, 49854, 86609, 89364, 74266, 52098, 96964, 49393, 50534, 15769, 80451, 74551, 91700, 93090, 17501, 16736, 46056, 75968, 97701, 28685, 51195, 74109, 61906, 61901, 27949, 80781, 45869, 76250, 92161, 92901, 84135, 36445, 91602, 14229, 26675, 31839, 79535, 36655, 68093, 49052, 38777, 10824, 38968, 63083, 95411, 69727, 93596, 52663, 65999, 59890, 66226, 94707, 13899, 25283, 79946, 80152, 78359, 28675, 51401, 21022, 96272, 22693, 21648, 46178, 17466, 82661, 39775, 48681, 47199, 30660, 18285, 55428, 98906, 86638, 48728, 78268, 32191, 75544, 38692, 51182, 24564, 78193, 14830, 51013, 40389, 80648, 25915, 75278, 73740, 47851, 30743, 30074, 47127, 66658, 81404, 78838, 98988, 61023, 44069, 58419, 88253, 10357, 65647, 56636, 96378, 61167, 34557, 64400, 51614, 12364, 71834, 11214, 84197, 46279, 85435, 87320, 28341, 65983, 94333, 66873, 50771, 91221, 63742, 99770, 37731, 20432, 39263, 99080, 29242, 93843, 53850, 68835, 43354, 48317, 89326, 75434, 70473, 63394, 38016, 78812, 93261, 20014, 11924, 21157, 24802, 55997, 87219, 20988, 39639, 29464, 61672, 96734, 55764, 52658, 66306, 19417, 20305, 33337, 58435, 88271, 72760, 62171, 35420, 85791, 86874, 47942, 65982, 54647, 76794, 71733, 99560, 57741, 19684, 82201, 56598, 78671, 57520, 28794, 70718, 12894, 80420, 75179, 37527, 68865, 54484, 22482, 28787, 23380, 24819, 23039, 34777, 38089, 40091, 65391, 64984, 47803, 54239, 89439, 64104, 42956, 60123, 62887, 50719, 60551, 32730, 80253, 31948, 33170, 43380, 66465, 68963, 74433, 42272, 45882, 62636, 97266, 23083, 27358, 44481, 11314, 64712, 95360, 43100, 22699, 95680, 12170, 51297, 48255, 23355, 55264, 43234, 58663, 24467, 31569, 91720, 17395, 87456, 61344, 50316, 50313, 33328, 70391, 40124, 80337, 66151, 14932, 34464, 18164, 84969, 30898, 76526, 41007, 46997, 74614, 97267, 13525, 83587, 70678, 96613, 39592, 44365, 44840, 99891, 58567, 58851, 87728, 84786, 13329, 93229, 43915, 83982, 19567, 33979, 40170, 97852, 45928, 74935, 97730, 74826, 68208, 53130, 41725, 57089, 98073, 23652, 97201, 87155, 88640, 37745, 14056, 95658, 48755, 13619, 81792, 78244, 30595, 84035, 60013, 30544, 94649, 12343, 37130, 89486, 30879, 82252, 72035, 96757, 12935, 76836, 13035, 64394, 68916, 81474, 90650, 56290, 35758, 58825, 55904, 39969, 45804, 39063, 71312, 98164, 83970, 63860, 19008, 36342, 75519, 70509, 15393, 98995, 80265, 95001, 28863, 65186, 91417, 18213, 46598, 17350, 77732, 43239, 91884, 98023, 30603, 50132, 14657, 89888, 29239, 33269, 86498, 32323, 87234, 91312, 55972, 95671, 47556, 64231, 40010, 78889, 65672, 66829, 58976, 52684, 12106, 40894, 73516, 98758, 78558, 91842, 36111, 81902, 57016, 62322, 12386, 46583, 16653, 90664, 99386, 53504, 88729, 33262, 22996, 49944, 71217, 89409, 81592, 98029, 54437, 73652, 44888, 16495, 20340, 69933, 23100, 27337, 81726, 36381, 13641, 93685, 73502, 95815, 26455, 89509, 43369, 78840, 84665, 27145, 20548, 51087, 24544, 72048, 70432, 66825, 30189, 87176, 69550, 52304, 22031, 41939, 30725, 71537, 53526, 65213, 51028, 87427, 18856, 12027, 28641, 49464, 15977, 58898, 39039, 67691, 60850, 13012, 84069, 32568, 96125, 25181, 73744, 56977, 80745, 68012, 47755, 78533, 23621, 11353, 84311, 88935, 54337, 62396, 83685, 13519, 89345, 91712, 54787, 41442, 57095, 73305, 40222, 18020, 62794, 26368, 11475, 75682, 75582, 97642, 43518, 61262, 65073, 64250, 41393, 43735, 52143, 20553, 88338, 92669, 97752, 30830, 13583, 34097, 95952, 20902, 35750, 79359, 19583, 77422, 63966, 25403, 71523, 93579, 35074, 75840, 44626, 72438, 54577, 66690, 41197, 13612, 25465, 55920, 89969, 58286, 43656, 82249, 30657, 39157, 18894, 11955, 43927, 16920, 29478, 77022, 33574, 52142, 43400, 20918, 59754, 54975, 45684, 70895, 68461, 96575, 39266, 85211, 62831, 58008, 52256, 67407, 11485, 73476, 28411, 58997, 50228, 22731, 60317, 50387, 47125, 32596, 86041, 17153, 47001, 95064, 60020, 97079, 95911, 42771, 31218, 58019, 27090, 42716, 13624, 14872, 15763, 59386, 29342, 85955, 92842, 84757, 76051, 27595, 74249, 76224, 46211, 16434, 95096, 65042, 89918, 57601, 45468, 38812, 51266, 90185, 75344, 68758, 54439, 49138, 72447, 24656, 94875, 22373, 44518, 43686, 32501, 42849, 96366, 97664, 75951, 87742, 26127, 49148, 44880, 48351, 49769, 53310, 97615, 44217, 35613, 37774, 37902, 20513, 22058, 69695, 12676, 78546, 93366, 43779, 42156, 26825, 61551, 79992, 70690, 60303, 47232, 27958, 65137, 81422, 90350, 82498, 34184, 49629, 52046, 68564, 74624, 84055, 78176, 19369, 29606, 56381, 41549, 98646, 10620, 71804, 70094, 22047, 26955, 20916, 70578, 31396, 38677, 51586, 33206, 40909, 86670, 71335, 90443, 76017, 64608, 46749, 89544, 41857, 42787, 73907, 62371, 24918, 51056, 39721, 67175, 30403, 64135, 76994, 66535, 59032, 75831, 30096, 23350, 59712, 38953, 60226, 82255, 36348, 76999, 10164, 15946, 77511, 95601, 18152, 32771, 20404, 61545, 34285, 84576, 51878, 88769, 63200, 68855, 65595, 48751, 60205, 96460, 89979, 11820, 83877, 20276, 19438, 94604, 14247, 89516, 44165, 58766, 11456, 11021, 55218, 74173, 46876, 63558, 64004, 72949, 57103, 95565, 16875, 81532, 15150, 74559, 62813, 29326, 81512, 49315, 76486, 42556, 72818, 24550, 64689, 73254, 93699, 53060, 93009, 39670, 61036, 45671, 41312, 49132, 38018, 71072, 22869, 55057, 24941, 84686, 28947, 25665, 78778, 14626, 37342, 59900, 43285, 92065, 20557, 45377, 14449, 97504, 53508, 39463, 40767, 64115, 57015, 10790, 39059, 97687, 80185, 42697, 84512, 88890, 26625, 23057, 86691, 19977, 20885, 82396, 95747, 70767, 36646, 55189, 88283, 61823, 81975, 96120, 27584, 29777, 83851, 68499, 71570, 83910, 72510, 29704, 58042, 10666, 25477, 46516, 11780, 60016, 76061, 45171, 16022, 43064, 74174, 75767, 94834, 43884, 73079, 49817, 84945, 37442, 75204, 88344, 78701, 79748, 74272, 32505, 74004, 97970, 43969, 34338, 56706, 70683, 51351, 15913, 34587, 27138, 81477, 37366, 15828, 29715, 97538, 16406, 68956, 21846, 76758, 59057, 43840, 67201, 49427, 41187, 36754, 23374, 39822, 66021, 42652, 24842, 24280, 88867, 89241, 23591, 59260, 50605, 24330, 25882, 35776, 54141, 75680, 84399, 78265, 87750, 36091, 27814, 63430, 63140, 81319, 89693, 44939, 78937, 34287, 60288, 39738, 66981, 93466, 43980, 50378, 65798, 33048, 74858, 41614, 78108, 95851, 55294, 91809, 18352, 92466, 20367, 94765, 70536, 83561, 80370, 70474, 96303, 12398, 23108, 53083, 99161, 85566, 32257, 46065, 63630, 42636, 94467, 62773, 56672, 25886, 13184, 24742, 69268, 22434, 12984, 27117, 10754, 24873, 39306, 68234, 17427, 97952, 33373, 76327, 45608, 48367, 54606, 24726, 10736, 13367, 95178, 27296, 68391, 98563, 20841, 20931, 60534, 44041, 39264, 12510, 65091, 73605, 43630, 80045, 17857, 71065, 67686, 40016, 96408, 34507, 10576, 62967, 68474, 67663, 27217, 45444, 16316, 53563, 76310, 56858, 96287, 80543, 82061, 44833, 58023, 93209, 19697, 53393, 97156, 13069, 72142, 57137, 48483, 72970, 93043, 56792, 73491, 95205, 84502, 50655, 81326, 30947, 81476, 54995, 88684, 16695, 67522, 14444, 48169, 77304, 10022, 85352, 88896, 80571, 75090, 45429, 38945, 70349, 13113, 34796, 32195, 28292, 14969}
	e := []int{14872, 86182, 43656, 85315, 90834, 80045, 46279, 30947, 21374, 72621, 12389, 67579, 43656, 63407, 14457, 88395, 96234, 18203, 68865, 46056, 31396, 38740, 86211, 21692, 34796, 15045, 37224, 93912, 75324, 71207, 28531, 81902, 30947, 26456, 31935, 31935, 63848, 91948, 88896, 95050, 56554, 85108, 50652, 42099, 53544, 80370, 80370, 21828, 10740, 30403, 87413, 94461, 80370, 81902, 55634, 81902, 70536, 31396, 68564, 40680, 43656, 13595, 79022, 66901, 33564, 25915, 38273, 78207, 24017, 36263, 97201, 27307, 97487, 97979, 60051, 25915, 93045, 60562, 86686, 85874, 69937, 50228, 13151, 57318, 66727, 80222, 18625, 96995, 27814, 72597, 14872, 43656, 98087, 91047, 62018, 34735, 48415, 30947, 76891, 63741, 93960, 81716, 31935, 14609, 13612, 34796, 56858, 95843, 79751, 22373, 33943, 68564, 75968, 49814, 30947, 62569, 37250, 55670, 31935, 73721, 63350, 15293, 67580, 19369, 25099, 40454, 16142, 71345, 79356, 33981, 31396, 54646, 34796, 46279, 80045, 21157, 76136, 56687, 32173, 46279, 18085, 75005, 40124, 29974, 97201, 22373, 25915, 48255, 68564, 80045, 58244, 64971, 11780, 53840, 89474, 38754, 25915, 14872, 28428, 51586, 78802, 65594, 14925, 49815, 81902, 81326, 17693, 15419, 11780, 55426, 70448, 41784, 75968, 80045, 28578, 65464, 43656, 80045, 72117, 56858, 97201, 48255, 38234, 34796, 81902, 30403, 81902, 80045, 36673, 89585, 88896, 80045, 94333, 89708, 70184, 90992, 99580, 50228, 34018, 26417, 30403, 98258, 90301, 81902, 88288, 19170, 23050, 35940, 15601, 50228, 42357, 87025, 56858, 40959, 73652, 80045, 19369, 19369, 88896, 54668, 25915, 40124, 11349, 51505, 26196, 11220, 96587, 76569, 81623, 30728, 57329, 81902, 83387, 72297, 57279, 30947, 59853, 49975, 34796, 96834, 95736, 66937, 93318, 39898, 37056, 97253, 68460, 13591, 50228, 82900, 41831, 98955, 13612, 21838, 17628, 57871, 48255, 70338, 14872, 68564, 11873, 43656, 52069, 79205, 80045, 56858, 23900, 51307, 20308, 31396, 91747, 87970, 90134, 33972, 43402, 30403, 37227, 32507, 90123, 54511, 27814, 88896, 87294, 43656, 48739, 94736, 28466, 84011, 95244, 24359, 92198, 78978, 11780, 35872, 11780, 41849, 71818, 79810, 21157, 18175, 56858, 94333, 45286, 43656, 31396, 45955, 95996, 20192, 13612, 59748, 75968, 68564, 25915, 45350, 76332, 81326, 39307, 22093, 96648, 65492, 56858, 43656, 31396, 28126, 94702, 56672, 41199, 89015, 10888, 80699, 33732, 14949, 36727, 87031, 15740, 80045, 59349, 71974, 49833, 80714, 85760, 80370, 42793, 71880, 59896, 37319, 74991, 73652, 56836, 24742, 96665, 73652, 17915, 81326, 72240, 26301, 46056, 11140, 80370, 43609, 18944, 14483, 84354, 28052, 44209, 94712, 92523, 52869, 65874, 11780, 27274, 23077, 99212, 78809, 79813, 30403, 90245, 81135, 43656, 56858, 57457, 15315, 31642, 91620, 57818, 33234, 13612, 20422, 20328, 68564, 89807, 31396, 21157, 57474, 93464, 54135, 75431, 89337, 85178, 25915, 94354, 30974, 86157, 19122, 71603, 94579, 80370, 39385, 32550, 76018, 49303, 31396, 94333, 60492, 66925, 19369, 30682, 34796, 84227, 89081, 57058, 19369, 40124, 35430, 89952, 68440, 27404, 76800, 80370, 67957, 22373, 66477, 20836, 68564, 46945, 88896, 68564, 39589, 72751, 74910, 75990, 62184, 57314, 49736, 48255, 97201, 68865, 34796, 74948, 43656, 21157, 22716, 34796, 48255, 29915, 68865, 43088, 18978, 20794, 87927, 46248, 15158, 40474, 94333, 56672, 22373, 19369, 49689, 68074, 34796, 22841, 75878, 14901, 40124, 74145, 80370, 80045, 31396, 14872, 35575, 41578, 43656, 59876, 68865, 52835, 81503, 68564, 18898, 94320, 13612, 54600, 66295, 68564, 86876, 57300, 43538, 21157, 88896, 57385, 94333, 11780, 68865, 48255, 47096, 60433, 68564, 56858, 40124, 30497, 86662, 78165, 67123, 70389, 97260, 43755, 96171, 31396, 32478, 10307, 43656, 50228, 94333, 17632, 30653, 59421, 58568, 21157, 48255, 74539, 99214, 33500, 80541, 26109, 97201, 31396, 13612, 89741, 30403, 31396, 36725, 43656, 20566, 24742, 94333, 49787, 31396, 56858, 76520, 40699, 60374, 93396, 97796, 33317, 31715, 13537, 97201, 30403, 94333, 68797, 34796, 89943, 67333, 41803, 99088, 71716, 81287, 37265, 20523, 75968, 41922, 46689, 81347, 97201, 43561, 31396, 74640, 56672, 56858, 94333, 11780, 25915, 55991, 75968, 13612, 50890, 97201, 21157, 14872, 40124, 22613, 43361, 10761, 60879, 92097, 13612, 97201, 18585, 45638, 55727, 90094, 29249, 27199, 34429, 19369, 31313, 97700, 70925, 84873, 33414, 67163, 31396, 11456, 43656, 43656, 44430, 88781, 79593, 94062, 30947, 58309, 80045, 92191, 33345, 40844, 75524, 43403, 88896, 36543, 21157, 84567, 78699, 74192, 51586, 88896, 35652, 24112, 81326, 25074, 25915, 94333, 77958, 49796, 80045, 21157, 43771, 75968, 46087, 37746, 85745, 76047, 80370, 98673, 76255, 80370, 29711, 51586, 23086, 80370, 24742, 43543, 46279, 58205, 67100, 14872, 98011, 79598, 34796, 25915, 36271, 11780, 78902, 15092, 94333, 31184, 11447, 36783, 31994, 28778, 50228, 72338, 68865, 31396, 31710, 18029, 30248, 35450, 49627, 25915, 31396, 46279, 87463, 40124, 51586, 21515, 81326, 28497, 53032, 33285, 88896, 11367, 80045, 22373, 11780, 40124, 81902, 22373, 67536, 18681, 80045, 24742, 26306, 89613, 59490, 22743, 80459, 31935, 24611, 68564, 62092, 68564, 38769, 62587, 22210, 88896, 61861, 13612, 68564, 98827, 25915, 26135, 52964, 43656, 70406, 19265, 68564, 87046, 81902, 67615, 67631, 40413, 80045, 81902, 91922, 75073, 27664, 27814, 27959, 71679, 74369, 57496, 23305, 64462, 43855, 22416, 98392, 90613, 68564, 56858, 55097, 57790, 80370, 40124, 68564, 84078, 38612, 11780, 48255, 23054, 40124, 21344, 30568, 60974, 25773, 60654, 40392, 27723, 35230, 92827, 20497, 19369, 58544, 97285, 66161, 88896, 98911, 58310, 13612, 93381, 81326, 14896, 80045, 33934, 19369, 53243, 94333, 66744, 38332, 71338, 81760, 44454, 88896, 42137, 31396, 80370, 94333, 10987, 41861, 25498, 82486, 93178, 88121, 68004, 52912, 57580, 55148, 40124, 34796, 73879, 40075, 88796, 81902, 27951, 51586, 13612, 41678, 56236, 11780, 92114, 85239, 34453, 68865, 56858, 80370, 55594, 18077, 88896, 36425, 25915, 82650, 41066, 70257, 14872, 24495, 47449, 68865, 70536, 43243, 88896, 24742, 28631, 86694, 62627, 48813, 28489, 84536, 80024, 62926, 50228, 40124, 31530, 80370, 13922, 75968, 53359, 34796, 85215, 68564, 56300, 86165, 88896, 57052, 22373, 97201, 15143, 85978, 23335, 24498, 80045, 70536, 51594, 26882, 87082, 39538, 27347, 51930, 94333, 41921, 84778, 21872, 59326, 73432, 31935, 44396, 44391, 89041, 25947, 21157, 24742, 15007, 68564, 50230, 14872, 97201, 24904, 20111, 21157, 30403, 92058, 94333, 31396, 77346, 76178, 43656, 57124, 75979, 25915, 84487, 40124, 20994, 94333, 14493, 30947, 22373, 71965, 25915, 35060, 53242, 56815, 56858, 64705, 38903, 80370, 11515, 73652, 90593, 69166, 77382, 14872, 78383, 50005, 78823, 66318, 13612, 73652, 11780, 29692, 81902, 22373, 92940, 80370, 49384, 70536, 94221, 80370, 68564, 56858, 69426, 30104, 29117, 84323, 20029, 41150, 56858, 60022, 23230, 56858, 63002, 31122, 49844, 38592, 69357, 25915, 15463, 80370, 30947, 92854, 31935, 86245, 21547, 73108, 86542, 80659, 77477, 17661, 24717, 93182, 62773, 22838, 68865, 34807, 21157, 87473, 57148, 91244, 21157, 56858, 40124, 49176, 57181, 84572}

	slices.Sort(s)
	slices.Sort(e)

	result := calculateMapScore(e, s)

	pl("result", result)
}
