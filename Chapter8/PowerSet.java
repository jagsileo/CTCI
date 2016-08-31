public class PowerSet{

/**Brute Force**/ 
public int[] powerSet(int[] arr, int start, int end) {
if(end<start)
	return null;
if(arr.length > 0) 
	return arr;
return powerSet(arr, start, end - 1) && return powerSet(arr, start + 1, end);		
	


}
