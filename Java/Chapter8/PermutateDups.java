import java.util.*;
import java.lang.*;
public class PermutateDups {

	public static ArrayList<String> permutateWithDups(String s) {
		ArrayList<String> permutations = new ArrayList<String>();
		if(s.length() <= 0) {
			permutations.add("");
		}
		if(s.length() == 1) {
			permutations.add(s);
		}

		if(s.length() > 1) {
			Map<Character, Integer> charCount = getCount(s);
			if(s.length() == charCount.size()) {
				permutations = PermutateString.permutate(s);				
			}
			else if(charCount.size() == 1) {
				permutations.add(s);
			}
			else {
				for()
			}
			
		}
	}

	public static Map<Character, Integer> getCount(String str) {
		Map<Character, Integer> countChar = new HashMap<Character, Integer>();
		for(char c:str.toCharArray()) {
			if(countChar.containsKey(c)) {
				countChar.put(c, countChar.get(c) + 1);
			}
			else {
				countChar.put(c, 1);
			}
		}
		return countChar;
	}


	public static void main(String[] args) {


	}
}


