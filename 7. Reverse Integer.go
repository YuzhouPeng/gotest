package leetcode_go
const INT_MAX = math.MaxInt32
const INT_MIN = -math.MaxInt32

func reverse(x int) int {
	var rev int = 0
	for {
		if x==0{
			break
		}

		var pop int= x%10
		x/=10
		if (rev> INT_MAX/10 ||(rev==INT_MAX/10 && pop>7)){
			return 0
		}
		if (rev< INT_MIN/10 ||(rev==INT_MIN/10 && pop<(-8))){
			return 0
		}
		rev = rev*10+pop

	}
	if (rev>=INT_MAX){
		return 0
	}
	if (rev<=INT_MIN){
		return 0
	}
	return rev
}