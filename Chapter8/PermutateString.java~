import java.util.*;
import java.lang.*;
public class PermutateString {

	public static ArrayList<String> permutate(String s) {
		ArrayList<String> arr = new ArrayList<String>();
		if(s == null) {
			arr.add("Null Input");
		}

		if(s.length() == 1) {
			arr.add(s);
			//return arr;
		}

		if(s.length() == 2) {
			arr.add(s);
			char[] ch = s.toCharArray();
			char buf = ch[0];
			ch[0] = ch[1];
			ch[1] = buf;			
			arr.add(new String(ch));			
		}
		if(s.length() > 2) {
			ArrayList<String> prev = permutate(s.substring(0, s.length() - 1));
			if(prev != null) {
				char last = s.charAt(s.length() - 1);
				for(int i = 0; i<prev.size(); i++) {

					String p = prev.get(i);
					arr.add(last + p);
					for(int j = 1; j<p.length(); j++) {
						//System.out.println(j + p.substring(0, j) + last + p.substring(j));
						arr.add(p.substring(0, j) + last + p.substring(j));
					}
					arr.add(p + last);			
				}
			}
		}
		return arr;


	}

	public static void main(String[] args) {
		ArrayList<String> result = permutate("vignesh");
		System.out.println(result.size());		
	}
}





