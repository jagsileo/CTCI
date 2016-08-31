/* package whatever; // don't place package name! */

import java.util.*;
import java.lang.*;
import java.io.*;

/* Name of the class has to be "Main" only if the class is public. */
class IsUnique
{
	
	public static void main(String[] args) {
		String str = "abc cef";
		IsUnique obj = new IsUnique();
		System.out.println(obj.isUnique(str));
	}

	public boolean isUnique(String st) {
		int len = st.length();
		boolean unique = true;		
		st.toLowerCase();
		for(int i = 0; i<len; i++) {
			for(int j=i+1; j<len; j++) {
				if(st.charAt(i) == st.charAt(j)) {
					unique = false;
					break;
				}
			}
		}
		return unique;
	}
}
