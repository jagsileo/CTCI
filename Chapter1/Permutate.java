import java.util.*;
class Permutate {
	public static void main(String[] args) {
		String str1 = "goddy";
		String str2 = "doggy";
		char[] chars1 = str1.toCharArray();
		Arrays.sort(chars1);
		char[] chars2 = str2.toCharArray();
		Arrays.sort(chars2);
		String strn1 = new String(chars1);
		String strn2 = new String(chars2);
		if(strn1.equals(strn2)) {
			System.out.println("Yes");
		}
		else {
			System.out.println("No");
		}
	}
}
