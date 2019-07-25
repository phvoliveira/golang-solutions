package main

import (
	"fmt"
)

func minimumSwaps(arr []int32) int32 {
	var count int32
	var i int

	for i < len(arr) {
		itemIndex := int(arr[i]) - 1
		if itemIndex != i {
			arr[i], arr[itemIndex] = arr[itemIndex], arr[i]
			count++
		} else {
			i++
		}
	}

	return count
}

func main() {
	fmt.Printf("Primeiro Resultado: %v\n", minimumSwaps([]int32{4, 3, 1, 2}))
	fmt.Printf("Segundo Resultado: %v\n", minimumSwaps([]int32{2, 3, 4, 1, 5}))
	fmt.Printf("Terceiro Resultado: %v\n", minimumSwaps([]int32{1, 3, 5, 2, 4, 6, 7}))
	fmt.Printf("Quarto Resultado: %v\n", minimumSwaps([]int32{4, 3, 5, 7, 1, 6, 2}))
	fmt.Printf("Teste Final: %v\n", minimumSwaps([]int32{974, 328, 380, 817, 304, 214, 787, 319, 291, 144, 393, 581, 838, 745, 946, 535, 61, 413, 441, 703, 901, 609, 114, 10, 972, 806, 218, 275, 300, 792, 249, 554, 407, 741, 979, 539, 164, 163, 874, 552, 428, 130, 480, 24, 746, 382, 306, 952, 766, 55, 423, 881, 813, 253, 203, 294, 467, 337, 368, 798, 281, 523, 311, 998, 518, 911, 309, 601, 782, 743, 479, 199, 930, 179, 727, 777, 707, 973, 595, 165, 633, 255, 639, 26, 532, 30, 400, 840, 252, 748, 737, 567, 562, 965, 825, 857, 791, 955, 716, 991, 142, 210, 886, 100, 408, 773, 174, 730, 91, 38, 759, 987, 313, 755, 483, 182, 833, 204, 119, 713, 21, 951, 781, 197, 71, 229, 702, 768, 361, 489, 731, 807, 941, 398, 495, 687, 75, 99, 624, 31, 153, 541, 647, 143, 508, 2, 405, 578, 982, 112, 152, 873, 525, 610, 329, 900, 856, 889, 357, 574, 51, 953, 537, 121, 470, 648, 835, 794, 14, 240, 348, 883, 421, 868, 454, 751, 16, 937, 507, 202, 243, 872, 438, 104, 80, 764, 645, 254, 804, 8, 637, 910, 772, 151, 653, 28, 515, 836, 132, 810, 996, 606, 544, 778, 655, 738, 880, 50, 97, 757, 921, 612, 187, 579, 430, 435, 896, 960, 843, 347, 148, 908, 723, 673, 346, 206, 821, 138, 585, 569, 43, 620, 656, 318, 137, 672, 841, 449, 822, 580, 76, 550, 213, 947, 231, 797, 63, 310, 117, 750, 516, 396, 641, 608, 823, 816, 514, 568, 614, 217, 783, 450, 377, 297, 340, 561, 73, 734, 136, 799, 162, 774, 559, 864, 205, 453, 23, 472, 892, 424, 140, 642, 924, 524, 909, 975, 596, 847, 760, 146, 317, 695, 186, 697, 172, 420, 805, 717, 170, 451, 372, 417, 780, 842, 485, 986, 375, 270, 322, 636, 227, 44, 201, 488, 267, 236, 410, 216, 983, 963, 771, 434, 719, 221, 37, 17, 33, 362, 519, 626, 762, 25, 993, 747, 302, 812, 191, 110, 93, 597, 828, 682, 819, 282, 312, 679, 869, 761, 367, 712, 351, 557, 345, 686, 859, 865, 749, 546, 167, 677, 861, 916, 1, 158, 314, 409, 326, 343, 464, 239, 12, 917, 257, 237, 383, 837, 926, 894, 528, 674, 553, 260, 139, 299, 74, 852, 635, 968, 887, 378, 77, 448, 888, 124, 854, 220, 211, 154, 141, 681, 978, 455, 419, 66, 126, 870, 127, 27, 827, 298, 970, 444, 815, 185, 366, 820, 264, 906, 700, 769, 107, 305, 57, 147, 534, 685, 928, 285, 334, 198, 316, 533, 664, 276, 878, 849, 591, 7, 666, 728, 529, 458, 277, 395, 629, 376, 560, 129, 551, 381, 200, 111, 935, 471, 244, 621, 994, 512, 944, 691, 790, 789, 494, 890, 90, 669, 262, 684, 920, 19, 95, 271, 195, 246, 307, 176, 287, 501, 20, 720, 209, 284, 330, 526, 785, 721, 359, 497, 352, 88, 948, 219, 715, 443, 565, 587, 161, 371, 358, 903, 68, 885, 456, 384, 323, 758, 689, 42, 705, 542, 251, 932, 86, 696, 22, 387, 134, 386, 517, 754, 15, 336, 447, 663, 177, 867, 676, 169, 9, 32, 196, 616, 113, 70, 593, 509, 826, 365, 575, 69, 289, 801, 997, 155, 505, 230, 775, 250, 957, 962, 288, 427, 333, 884, 594, 279, 500, 178, 607, 83, 829, 59, 47, 649, 536, 324, 452, 215, 690, 603, 971, 683, 120, 584, 64, 247, 48, 942, 283, 58, 466, 735, 82, 660, 652, 308, 412, 349, 511, 605, 89, 722, 46, 658, 339, 487, 81, 29, 72, 832, 431, 904, 92, 498, 611, 877, 266, 793, 905, 445, 228, 545, 803, 661, 54, 109, 699, 934, 36, 558, 13, 665, 320, 491, 929, 404, 222, 389, 5, 259, 98, 586, 954, 226, 156, 795, 740, 40, 234, 429, 265, 335, 476, 530, 238, 604, 897, 657, 753, 331, 390, 701, 62, 543, 776, 619, 11, 765, 65, 549, 893, 159, 602, 41, 876, 274, 956, 115, 232, 188, 131, 851, 898, 866, 439, 570, 521, 493, 360, 363, 788, 520, 85, 714, 469, 212, 985, 858, 736, 207, 379, 589, 183, 506, 461, 477, 388, 39, 980, 871, 907, 919, 855, 913, 492, 977, 615, 853, 415, 157, 149, 96, 327, 659, 355, 35, 102, 646, 292, 834, 411, 401, 710, 862, 927, 651, 680, 708, 625, 341, 667, 992, 698, 692, 671, 824, 573, 263, 711, 590, 171, 964, 87, 128, 332, 844, 984, 563, 6, 976, 194, 503, 693, 931, 600, 356, 496, 599, 742, 463, 638, 918, 945, 879, 538, 150, 369, 613, 891, 399, 640, 634, 739, 522, 566, 193, 915, 208, 296, 325, 571, 425, 397, 135, 504, 709, 391, 233, 961, 4, 863, 123, 831, 278, 106, 175, 969, 950, 426, 933, 811, 364, 474, 694, 547, 678, 78, 668, 422, 241, 486, 258, 1000, 116, 718, 414, 752, 767, 848, 160, 784, 922, 293, 478, 577, 502, 225, 943, 168, 630, 995, 839, 184, 374, 446, 850, 173, 786, 315, 662, 457, 949, 583, 654, 52, 989, 350, 242, 406, 105, 403, 468, 354, 392, 592, 618, 617, 733, 527, 45, 290, 133, 845, 49, 295, 846, 990, 548, 224, 576, 145, 442, 770, 190, 724, 181, 166, 402, 433, 484, 338, 118, 818, 531, 235, 988, 437, 344, 67, 261, 555, 385, 923, 966, 3, 510, 999, 796, 256, 875, 490, 280, 914, 459, 726, 588, 223, 643, 465, 981, 744, 627, 272, 936, 180, 248, 394, 675, 79, 499, 460, 53, 125, 353, 860, 732, 103, 779, 60, 84, 108, 473, 925, 122, 540, 373, 189, 808, 269, 814, 273, 582, 632, 622, 650, 725, 436, 18, 631, 475, 34, 101, 756, 321, 704, 556, 688, 342, 959, 895, 418, 432, 763, 899, 882, 729, 938, 912, 644, 482, 670, 94, 902, 303, 416, 967, 939, 56, 623, 802, 370, 572, 830, 301, 958, 513, 809, 268, 245, 286, 192, 564, 481, 462, 940, 628, 706, 800, 440, 598}))
}
