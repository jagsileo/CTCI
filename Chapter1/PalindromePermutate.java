 import java.util.*;


public class PalindromePermutate {
		
	private static HashMap buildMap(String s) {
	    HashMap<Character, Integer> hm = new HashMap<Character, Integer>();
	    for(int i = 0; i < s.length(); i++) {
	        char key = s.charAt(i);
	        if(hm.containsKey(key)) {
	            hm.put(key, hm.get(key) + 1);
	        }
	        else {
	            hm.put(key, 1);
	        }
	    }
	    return hm;
	}
	
	private static Boolean checkPalindrome(HashMap hm, String s) {
	    int tolerance = 0;
	    Iterator it = hm.entrySet().iterator();
	    while(it.hasNext()) {
	        Map.Entry pair = (Map.Entry) it.next();
	        if((int) pair.getValue() % 2 != 0) {
	            tolerance++;
	        }
	    }
	        if(s.length() % 2 == 0) {
	            return tolerance == 0;
	        }
	        else {
	            return tolerance == 1;
	        }
	   
	}
	
	
	public static void main(String args[]){
		String st = "taco t";
		st = st.replace(" ", "");
		HashMap h = buildMap(st);
		System.out.println(checkPalindrome(h, st));
	}
}
