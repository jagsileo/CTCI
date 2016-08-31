 import java.util.*;


public class StringCompress {
		
	private static HashMap buildMap(String st) {
	    HashMap<Character, Integer> hm = new HashMap<Character, Integer>();
	    for(int i = 0; i<st.length(); i++) {
	        char key = st.charAt(i);
	        if(hm.containsKey(key)) {
	            hm.put(key, hm.get(key) + 1);
	        }
	        else {
	            hm.put(key, 1);
	        }
	    }
	    return hm;
	}
	
	private static String compressString(HashMap hm, String st) {
	    String s = "";
	    Iterator it = hm.entrySet().iterator();
	    while(it.hasNext()) {
	    Map.Entry pair = (Map.Entry) it.next();
	    s = s + (char)pair.getKey() + pair.getValue();
	    }
	    if(s.length() > st.length()) {
	        return st;
	    }
	    else {
	        return s;
	    }
	}
	
	
	public static void main(String args[]){
		String a = "aabbbbbcc";
		a = a.toLowerCase();
		HashMap h = buildMap(a);
		System.out.println(compressString(h, a));
	}
}
