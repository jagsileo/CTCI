import myTree.*;
import java.util.*;
public class FirstAncestor {


	public static String whichSide(MyBTNode root, MyBTNode n) {
		if(root == null || n == null) {
			return "not found";
		}
		if(root == n) {
			return "root";
		}
		if(root.left != null && root.left == n) {
			return "left";
		}
		if(root.right != null && root.right == n) {
			return "right";
		}
		if(root.left != null && isDescendant(root.left, n)) {
			return "left";
		}
		if(root.right != null && isDescendant(root.right, n)) {
			return "right";
		}	

	}



	public static Boolean isDescendant(MyBTNode root, MyBTNode n) {
		if(root == null || n == null) {
			return false;
		}
		if(root == n) {
			return true;
		}
		if(root.left != null && root.left == n) {
			return true;
		} 
		else if(root.left != null && root.right == n) {
			return true;
		}
		else if(root.left != null) {
			return isDescendant(root.left, n);
		}
		else if(root.right != null) {
			return isDescendant(root.right,n);
		}
	}
	
	public static MyBTNode findAncestor(MyBTNode root, MyBTNode n1, MyBTNode n2) {
		if(root == null || n1 == null || n2 == null) {
			return null;
		}
		
		if(root == n1) {
			return n1;
		}
		if(root == n2) {
			return n2;
		}
		if(whichSide(root, n1) != whichSide(root, n2)) {
			return root;
		}
		if(whichSide(root, n1) == "left") {
			return findAncestor(root.left, n1, n2);
		}
		if(whichSide(root, n1) == "right") {
			return findAncestor(root.right, n1, n2);
		}
	}
	
	
	public static void main(String[] args) {
		MyBTNode root = new MyBTNode(20);
		root.buildBSTInOrder(new MyBTNode(2));
		root.buildBSTInOrder(new MyBTNode(12));
		root.buildBSTInOrder(new MyBTNode(5));
		root.buildBSTInOrder(new MyBTNode(29));
		root.buildBSTInOrder(new MyBTNode(45));
		root.buildBSTInOrder(new MyBTNode(6));
		root.buildBSTInOrder(new MyBTNode(102));
	}
}

