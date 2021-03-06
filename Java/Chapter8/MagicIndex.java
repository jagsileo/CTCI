//Input -> arr
//Output -> magicIndex integer

public class MagicIndex{
	int[] arr;
	/**Brute Force - O(n)**/
	public int magicIndex(int[] arr) {
		for(int i = 0; i<arr.length; i++) {
			if(arr[i] == i) {
				return i;
			}
		}
		return -1;		
	}

	/**Distinct and sorted - O(log n)**/
	public int magicIndex(int[] arr, int start, int end) {
		if(end < start)
			return -1;
		int mid = (start + end)/2;

		if(arr[mid] == mid)
			return mid;

		if(arr[mid] < mid) 
			return magicIndex(arr, mid+1, end);

		if(arr[mid] > mid)
			return magicIndex(arr, start, mid-1);			
	}

	/**Non Distinct and sorted**/
	public int magicIndex(int[] arr, int start, int end) {
		if(end < start)
			return -1;
		int mid = (start + end)/2;
		int midValue = arr[mid];
		if(midValue == mid)
			return mid;
		int left = magicIndex(arr, start, Math.min(midValue, mid - 1));
		if(left > 0)
			return left;
		int right = magicIndex(arr, Math.max(midValue, mid+1), end);	
		if(right > 0)
			return right;

	}

}

