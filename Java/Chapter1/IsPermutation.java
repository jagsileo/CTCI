 import java.util.*;


public class IsPermutation {
		
	private HashMap arrange(String s) {
	    HashMap<Character, Integer> hm = new HashMap<Character, Integer>();
	    for(int i = 0; i<s.length(); i++) {
	        if(hm.containsKey(s.charAt(i))) {
	            hm.put(s.charAt(i),hm.get(s.charAt(i)) + 1);
	        }
	        else {
	            hm.put(s.charAt(i),1);
	        }
	    }
	    return hm;
	}
	
	private boolean compare(HashMap hm1, HashMap hm2) {
	    return(hm1.equals(hm2));
	}
	
	
	public static void main(String args[]){
		String s1 = "god";
		String s2 = "dog";
		MyClass mc = new MyClass();
		System.out.println(mc.arrange(s1));
		System.out.println(mc.arrange(s2));
		System.out.println(mc.compare(mc.arrange(s1), mc.arrange(s2)));
	}
}
