 import java.util.*;


public class CountEdits {
		
	private static HashMap buildHash(String s) {
	    HashMap<Character, Integer> hm = new HashMap<Character, Integer>();
	    for(int i = 0; i<s.length(); i++) {
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
	
	private static Boolean compareHash(HashMap hm1, HashMap hm2) {
	    int tolerance = 0;
	    Iterator it = hm1.entrySet().iterator();
	    while(it.hasNext()) {
	    Map.Entry pair = (Map.Entry) it.next();
	    char key = (char) pair.getKey();
	    if((hm2.containsKey(key)) && (pair.getValue() != hm2.get(key))) {
	        tolerance++;
	    }
	    else if(!(hm2.containsKey(key))) {
	        tolerance++;
	    }
	    }
	    
	    return tolerance < 2;
	}
	
	
	public static void main(String args[]){
		Scanner sc = new Scanner(System.in);
		String[] s = new String[2];
		s[0] = "pale";
		s[1] = "pkbe";
		
		HashMap hm1 = buildHash(s[0]);
		HashMap hm2 = buildHash(s[1]);
		if(s[0].length() >= s[1].length()) {
		    System.out.println(compareHash(hm1, hm2));
		}
		else {
		    System.out.println(compareHash(hm2, hm1));
		}
	}
}
