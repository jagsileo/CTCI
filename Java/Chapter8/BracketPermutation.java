import java.util.*;
public class BracketPermutation {
	public static ArrayList<String> bracket(Stack s1, Stack s2, int input) {
		ArrayList<String> result = new ArrayList<String>();
		if(input == 0) {
			result.add("Empty");
			return result;
		}

		if(input == 1) {
			char c1 = (Character)s1.pop();
			char c2 = (Character)s2.pop();
			result.add((c1 + c2).toString());
		}

		if(input == 2) {
			char c1 = (Character)s1.pop();
			char c2 = (Character)s1.pop();
			char c3 = (Character)s2.pop();
			char c4 = (Character)s2.pop();
			result.add((c1 + c2 + c3 + c4).toString());
			result.add((c1 + c3 + c2 + c4).toString());
			return result;
		}

		char c1 = (Character)s1.pop();
		char c2 = (Character)s2.pop();
		ArrayList<String> recursed = bracket(s1, s2, input - 1);
		for(int i = 0; i < recursed.size(); i++) {
			String recString = recursed.get(i);
			result.add((c1 + recString + c2).toString());
			result.add((c1 + c2 + recString).toString());
			result.add((recString + c1 + c2).toString());
		}
		return result;
	}
	
	
	public static Stack buildOpenBrackets(int input) {
		Stack<Character> s = new Stack<Character>();
		for(int i = 0; i<input; i++) {
			s.push('(');
		}
		return s;
	}
	
	
	public static Stack buildCloseBrackets(int input) {
		Stack<Character> s = new Stack<Character>();
		for(int i = 0; i<input; i++) {
			s.push(')');
		}
		return s;
	}
	public static void main(String[] args) {
		Stack st1 = buildOpenBrackets(3);
		Stack st2 = buildCloseBrackets(3);
		ArrayList<String> bp = bracket(st1, st2, 3);
		for(int j = 0; j<bp.size(); j++) {
			System.out.println(bp.get(j));
		}
	}
}

